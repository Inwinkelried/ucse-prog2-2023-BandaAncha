package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PedidoRepositoryInterface interface {
	ObtenerPedidos() ([]model.Pedido, error)
	InsertarPedido(pedido model.Pedido) (*mongo.InsertOneResult, error)
	ObtenerPedidosAprobados() ([]model.Pedido, error)
	AceptarPedido(pedido model.Pedido) (*mongo.UpdateResult, error)
	CancelarPedido(pedido model.Pedido) (*mongo.UpdateResult, error)
	ParaEnviarPedido(pedido model.Pedido) (*mongo.UpdateResult, error)
	EnviadoPedido(pedido model.Pedido) (*mongo.UpdateResult, error)
	ObtenerPedidoPorID(pedidoConId model.Pedido) (model.Pedido, error)
	ObtenerPesoPedido(pedido model.Pedido) (int, error)
	ActualizarPedido(pedido model.Pedido) (*mongo.UpdateResult, error)
}
type PedidoRepository struct {
	db DB
}

func NewPedidoRepository(db DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func (repo PedidoRepository) ObtenerPesoPedido(pedido model.Pedido) (int, error) {
	//Obtener el el pedido por id, luego tomar su lista de productos. A cada producto multiplicarle su peso por la cantidad. A eso sumarlo y retornar ese total
	pedidoObtenido, err := repo.ObtenerPedidoPorID(pedido)
	if err != nil {
		return 0, err
	}
	var pesoTotal int
	for _, producto := range pedidoObtenido.Productos {
		pesoTotal += producto.PesoUnitario * producto.Cantidad
	}
	return pesoTotal, nil

}
func (repo PedidoRepository) ActualizarPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	pedido.FechaModificacion = time.Now()
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": pedido}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) ObtenerPedidosAprobados() ([]model.Pedido, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"Estado": "Aceptado"}
	cursor, err := lista.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())
	var Pedidos []model.Pedido
	for cursor.Next(context.Background()) {
		var Pedido model.Pedido
		err := cursor.Decode(&Pedido)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		Pedidos = append(Pedidos, Pedido)
	}
	return Pedidos, err
}
func (repository PedidoRepository) ObtenerPedidoPorID(pedidoABuscar model.Pedido) (model.Pedido, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedidoABuscar.ID}
	cursor, err := collection.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())
	// Itera a través de los resultados
	var pedido model.Pedido
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&pedido)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return pedido, err
}
func (repo PedidoRepository) ObtenerPedidos() ([]model.Pedido, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{}
	cursor, err := lista.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())

	var Pedidos []model.Pedido
	for cursor.Next(context.Background()) {
		var Pedido model.Pedido
		err := cursor.Decode(&Pedido)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		Pedidos = append(Pedidos, Pedido)
	}
	return Pedidos, err
}
func (repo PedidoRepository) InsertarPedido(pedido model.Pedido) (*mongo.InsertOneResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	resultado, err := lista.InsertOne(context.TODO(), pedido)
	return resultado, err
}
func (repo PedidoRepository) AceptarPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"Estado": "Aceptado"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) CancelarPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"Estado": "Cancelado"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) ParaEnviarPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"Estado": "Para enviar"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) EnviadoPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"Estado": "Enviado"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}

// ------------------------------------------------------------------------------