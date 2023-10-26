package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CamionRepositoryInterface interface {
	ObtenerCamiones() ([]model.Camion, error)
	EliminarCamion(id primitive.ObjectID) (*mongo.DeleteResult, error)
	InsertarCamion(camion model.Camion) (*mongo.InsertOneResult, error)
	ModificarCamion(camion model.Camion) (*mongo.UpdateResult, error)
	ObtenercamionPorID(camionABuscar model.Camion) (model.Camion, error)
}
type CamionRepository struct {
	db DB
}

func NewCamionRepository(db DB) *CamionRepository {
	return &CamionRepository{
		db: db,
	}
}

// Obtencion de todos los camiones
func (repo CamionRepository) ObtenerCamiones() ([]model.Camion, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Camiones")
	filtro := bson.M{}

	cursor, err := lista.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var Camiones []model.Camion
	for cursor.Next(context.Background()) {
		var Camion model.Camion
		err := cursor.Decode(&Camion)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		Camiones = append(Camiones, Camion)
	}
	return Camiones, err
}

// Metodo para instertar un camion nuevo
func (repo CamionRepository) InsertarCamion(camion model.Camion) (*mongo.InsertOneResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Camiones")
	resultado, err := lista.InsertOne(context.TODO(), camion)
	return resultado, err
}

// Metodo para obtener un camion por ID
func (repository CamionRepository) ObtenercamionPorID(camionABuscar model.Camion) (model.Camion, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Camiones")
	filtro := bson.M{"_id": camionABuscar.ID}
	cursor, err := collection.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())
	// Itera a través de los resultados
	var camion model.Camion
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&camion)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return camion, err
}

// Metodo para modificar un camion
func (repo CamionRepository) ModificarCamion(camion model.Camion) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Camiones")
	filtro := bson.M{"_id": camion.ID}
	entity := bson.M{"$set": bson.M{"CostoPorKilometro": camion.CostoKm, "FechaModificacion": time.Now()}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}

// Metodo para eliminar un camion
func (repo CamionRepository) EliminarCamion(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Camiones")
	filtro := bson.M{"_id": id}
	resultado, err := lista.DeleteOne(context.TODO(), filtro)
	return resultado, err
}
