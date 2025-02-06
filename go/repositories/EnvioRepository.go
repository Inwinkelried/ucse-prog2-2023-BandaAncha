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

type EnvioRepositoryInterface interface {
	ObtenerEnvios() ([]model.Envio, error)
	InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error)
	ActualizarEnvio(envio model.Envio) (*mongo.UpdateResult, error)
	ObtenerEnvioPorID(envio model.Envio) (model.Envio, error)
	ObtenerEnviosFiltrados(filtro dto.FiltroEnvio) ([]model.Envio, error)
}
type EnvioRepository struct {
	db DB
}

func NewEnvioRepository(db DB) *EnvioRepository {
	return &EnvioRepository{
		db: db,
	}
}

// TODO: probar esta función
func (repository EnvioRepository) ObtenerEnvioPorID(envioABuscar model.Envio) (model.Envio, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Envios")
	filtro := bson.M{"_id": envioABuscar.ID}
	cursor, err := collection.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())
	var envio model.Envio
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&envio)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return envio, err
}

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

func (repo EnvioRepository) InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Envios")
	resultado, err := lista.InsertOne(context.TODO(), envio)
	return resultado, err
}

func (repo EnvioRepository) ActualizarEnvio(envio model.Envio) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Envios")
	envio.FechaModificacion = time.Now()
	filtro := bson.M{"_id": envio.ID}
	entity := bson.M{"$set": envio}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repository *EnvioRepository) obtenerEnvios(filtro bson.M) ([]model.Envio, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Envios")
	cursor, err := collection.Find(context.Background(), filtro)
	if err != nil {
		return nil, err
	}
	var envios []model.Envio
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var envio model.Envio
		err := cursor.Decode(&envio)
		if err != nil {
			return nil, err
		}
		envios = append(envios, envio)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return envios, nil
}
func (repo EnvioRepository) ObtenerEnviosFiltrados(filtro dto.FiltroEnvio) ([]model.Envio, error) {
	filtroGenerado := bson.M{}
	if filtro.Estado != "" {
		filtroGenerado["Estado"] = filtro.Estado
	}
	if filtro.PatenteCamion != "" {
		filtroGenerado["patente_camion"] = filtro.PatenteCamion
	}
	// TODO: agregar filtro para las paradas del camión

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

	return repo.obtenerEnvios(filtroGenerado)

}
