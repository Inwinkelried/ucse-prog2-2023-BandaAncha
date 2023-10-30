package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EnvioRepositoryInterface interface {
	ObtenerEnvios() ([]model.Envio, error)
	InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error)
	ActualizarEnvio(envio model.Envio) (*mongo.UpdateResult, error)
	ObtenerEnvioPorID(envio model.Envio) (model.Envio, error)
}
type EnvioRepository struct {
	db DB
}

func NewEnvioRepository(db DB) *EnvioRepository {
	return &EnvioRepository{
		db: db,
	}
}

// Falta probar
func (repository EnvioRepository) ObtenerEnvioPorID(envioABuscar model.Envio) (model.Envio, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Envios")
	filtro := bson.M{"_id": envioABuscar.ID}
	cursor, err := collection.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())
	// Itera a través de los resultados
	var envio model.Envio
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&envio)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return envio, err
}

// Obtencion de todos los Envios
func (repo EnvioRepository) ObtenerEnvios() ([]model.Envio, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Envios")
	filtro := bson.M{}

	cursor, err := lista.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var envios []model.Envio
	for cursor.Next(context.Background()) {
		var envio model.Envio
		err := cursor.Decode(&envio)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		envios = append(envios, envio)
	}
	return envios, err
}

// Metodo para instertar un envio nueo
func (repo EnvioRepository) InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Envios")
	resultado, err := lista.InsertOne(context.TODO(), envio)
	return resultado, err
}

// Metodo para modificar un envio. Este metodo es el que me permite actualizar el estado del envio
func (repo EnvioRepository) ActualizarEnvio(envio model.Envio) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Envios")
	envio.FechaModificacion = time.Now()
	filtro := bson.M{"_id": envio.ID}
	entity := bson.M{"$set": envio}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
