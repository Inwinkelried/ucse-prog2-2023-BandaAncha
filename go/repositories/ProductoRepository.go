package repositories

import (
	"context"
	"fmt"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
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
	ObtenerProductoFiltrados(filtro dto.FiltroProducto) ([]model.Producto, error)
	DescontarStockProducto(producto model.Producto) (*mongo.UpdateResult, error)
}
type ProductoRepository struct {
	db DB
}

func NewProductoRepository(db DB) *ProductoRepository {
	return &ProductoRepository{db: db}
}
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
func (repo ProductoRepository) InsertarProducto(Producto model.Producto) (*mongo.InsertOneResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Productos")
	resultado, err := lista.InsertOne(context.TODO(), Producto)
	return resultado, err
}
func (repo ProductoRepository) ModificarProducto(Producto model.Producto) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Productos")
	filtro := bson.M{"_id": Producto.ID}
	entity := bson.M{"$set": Producto}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo ProductoRepository) EliminarProducto(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Productos")
	filtro := bson.M{"_id": id}
	resultado, err := lista.DeleteOne(context.TODO(), filtro)
	return resultado, err
}
func (repo ProductoRepository) DescontarStockProducto(producto model.Producto) (*mongo.UpdateResult, error) { //TODO: probar
	var productoFiltrado, err = repo.ObtenerProductoPorID(producto)
	if err != nil {
		return nil, err
	}
	var total int = productoFiltrado.StockActual - producto.StockActual
	productoFiltrado.StockActual = total
	return repo.ModificarProducto(productoFiltrado)
}
func (repo *ProductoRepository) obtenerProductos(filtro bson.M) ([]model.Producto, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Productos")
	cursor, err := lista.Find(context.Background(), filtro)
	if err != nil {
		return nil, err
	}
	var productos []model.Producto
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var producto model.Producto
		err := cursor.Decode(&producto)
		if err != nil {
			return nil, err
		}
		productos = append(productos, producto)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return productos, nil
}
func (repo ProductoRepository) ObtenerProductoFiltrados(filtro dto.FiltroProducto) ([]model.Producto, error) {
	filtroGenerado := bson.M{}
	if filtro.TipoProducto != "" {
		filtroGenerado["tipo"] = filtro.TipoProducto
	}
	if filtro.FiltroStockMinimo {
		filtroGenerado["$expr"] = bson.M{"$lt": []interface{}{"$stock_actual", "$stock_minimo"}}
	}
	return repo.obtenerProductos(filtroGenerado)
}
