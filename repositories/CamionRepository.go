package repositories

import (
	"context"
	"fmt"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/model"
)

type CamionRepositoryInterface interface {
	ObtenerCamiones() ([]model.Camion, error)
	EliminarCamion(id primitive.ObjectID) (*mongo.DeleteResult, error)
	InsertarAula(camion model.Camion) (*mongo.InsertOneResult, error)
	ModificarAula(camion model.Camion) (*mongo.UpdateResult, error)
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
