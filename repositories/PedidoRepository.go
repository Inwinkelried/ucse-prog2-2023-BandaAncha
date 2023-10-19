package repositories

import (
	"context"
	"fmt"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PedidoRepositoryInterface interface {
	ObtenerPedidos() ([]model.Pedido, error)
	EliminarPedido(id primitive.ObjectID) (*mongo.DeleteResult, error)
	InsertarPedido(pedido model.Pedido) (*mongo.InsertOneResult, error)
	ModificarPedido(pedido model.Pedido) (*mongo.UpdateResult, error)
}
type PedidoRepository struct {
	db DB
}

func NewPedidoRepository(db DB) *PedidoRepository {
	return &PedidoRepository{
		db: db,
	}
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
func (repo PedidoRepository) ModificarPedido(pedido model.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"estado": pedido.Estado}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) EliminarPedido(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": id}
	resultado, err := lista.DeleteOne(context.TODO(), filtro)
	return resultado, err
}
