package repositories

import (
	"context"
	"fmt"

	//"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductoRepositoryInterface interface {
	ObtenerProductos() ([]model.Producto, error)
	EliminarProducto(id primitive.ObjectID) (*mongo.DeleteResult, error)
	InsertarProducto(Producto model.Producto) (*mongo.InsertOneResult, error)
	ModificarProducto(Producto model.Producto) (*mongo.UpdateResult, error)
	ObtenerProductoPorID(productoAFiltrar model.Producto) (model.Producto, error)
}
type ProductoRepository struct {
	db DB
}

func NewProductoRepository(db DB) *ProductoRepository {
	return &ProductoRepository{db: db}
}

// Obtencion de todos los Productos
func (repo ProductoRepository) ObtenerProductos() ([]model.Producto, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Productos")
	filtro := bson.M{}

	cursor, err := lista.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var productos []model.Producto
	for cursor.Next(context.Background()) {
		var Producto model.Producto
		err := cursor.Decode(&Producto)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		productos = append(productos, Producto)
	}
	return productos, err
}

// Metodo para obtener un producto filtrado por ID
func (repository *ProductoRepository) ObtenerProductoPorID(productoAFiltrar model.Producto) (model.Producto, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Productos")
	filtro := bson.M{"_id": productoAFiltrar.ID}
	var producto model.Producto
	err := collection.FindOne(context.Background(), filtro).Decode(&producto)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return producto, err
}

// Metodo para instertar un Producto nuevo
func (repo ProductoRepository) InsertarProducto(Producto model.Producto) (*mongo.InsertOneResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Productos")
	resultado, err := lista.InsertOne(context.TODO(), Producto)
	return resultado, err
}

// Metodo para modificar un Producto
func (repo ProductoRepository) ModificarProducto(Producto model.Producto) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Productos")
	filtro := bson.M{"_id": Producto.ID}
	entity := bson.M{"$set": Producto}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}

// Metodo para eliminar un Producto
func (repo ProductoRepository) EliminarProducto(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Productos")
	filtro := bson.M{"_id": id}
	resultado, err := lista.DeleteOne(context.TODO(), filtro)
	return resultado, err
}
