package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
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
	ObtenerPedidosFiltrados(filtro dto.FiltroPedido) ([]model.Pedido, error)
	ObtenerCantidadPedidosPorEstado(model.EstadoPedido) (int, error)
}
type PedidoRepository struct {
	db DB
}

func NewPedidoRepository(db DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func (repository *PedidoRepository) ObtenerCantidadPedidosPorEstado(estado model.EstadoPedido) (int, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Pedidos")

	filtro := bson.M{"estado": estado}

	cantidad, err := collection.CountDocuments(context.Background(), filtro)

	if err != nil {
		return 0, err
	}

	return int(cantidad), nil
}

func (repo PedidoRepository) ObtenerPesoPedido(pedido model.Pedido) (int, error) {
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
	filtro := bson.M{"estado": "Aceptado"}
	cursor, err := lista.Find(context.TODO(), filtro)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var Pedidos []model.Pedido
	for cursor.Next(context.Background()) {
		var Pedido model.Pedido
		err := cursor.Decode(&Pedido)
		if err != nil {
			return nil, err
		}
		Pedidos = append(Pedidos, Pedido)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return Pedidos, nil
}
func (repository PedidoRepository) ObtenerPedidoPorID(pedidoABuscar model.Pedido) (model.Pedido, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedidoABuscar.ID}
	cursor, err := collection.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())
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
	pedido.Estado = "Pendiente"
	pedido.FechaCreacion = time.Now()
	pedido.FechaModificacion = time.Now()
	resultado, err := lista.InsertOne(context.TODO(), pedido)
	return resultado, err
}
func (repo PedidoRepository) AceptarPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"estado": "Aceptado"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) CancelarPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"estado": "Cancelado"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) ParaEnviarPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"estado": "Para enviar"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) EnviadoPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"estado": "Enviado"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repository *PedidoRepository) obtenerPedidos(filtro bson.M) ([]model.Pedido, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	cursor, err := collection.Find(context.Background(), filtro)
	if err != nil {
		return nil, err
	}
	var pedidos []model.Pedido
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var pedido model.Pedido
		err := cursor.Decode(pedido)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return pedidos, nil
}
func (repo PedidoRepository) ObtenerPedidosFiltrados(filtro dto.FiltroPedido) ([]model.Pedido, error) {
	filtroGenerado := bson.M{}
	if filtro.Estado != "" {
		filtroGenerado["Estado"] = filtro.Estado
	}
	if !filtro.FechaMenor.IsZero() || !filtro.FechaMayor.IsZero() {
		filtroFecha := bson.M{}
		if !filtro.FechaMenor.IsZero() {
			filtroFecha["$gte"] = filtro.FechaMenor
		}
		if !filtro.FechaMayor.IsZero() {
			filtroFecha["$lte"] = filtro.FechaMayor
		}
		filtroGenerado["fecha_creacion"] = filtroFecha
	}
	return repo.obtenerPedidos(filtroGenerado)
}
