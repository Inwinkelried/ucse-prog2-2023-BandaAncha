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
	InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error)
	ActualizarEnvio(envio model.Envio) (*mongo.UpdateResult, error)
	ObtenerEnvioPorID(envio model.Envio) (model.Envio, error)
	ObtenerEnvios(filtro dto.FiltroEnvio) ([]model.Envio, error)
	ObtenerCantidadEnviosPorEstado(model.EstadoEnvio) (int, error)
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

func (repo EnvioRepository) InsertarEnvio(envio model.Envio) (*mongo.InsertOneResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Envios")
	envio.FechaCreacion = time.Now()
	envio.FechaModificacion = time.Now()
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

// Obtener envíos aplicando un filtro (si no hay filtros, trae todos los envíos)
func (repo EnvioRepository) ObtenerEnvios(filtro dto.FiltroEnvio) ([]model.Envio, error) {
	filtroGenerado := construirFiltroEnvio(filtro)
	return repo.obtenerEnvios(filtroGenerado)
}

// Método generalizado para obtener envíos con un filtro (vacío si no se pasan filtros)
func (repository *EnvioRepository) obtenerEnvios(filtro bson.M) ([]model.Envio, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Envios")
	cursor, err := collection.Find(context.Background(), filtro)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var envios []model.Envio
	for cursor.Next(context.Background()) {
		var envio model.Envio
		if err := cursor.Decode(&envio); err != nil {
			return nil, err
		}
		envios = append(envios, envio)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return envios, nil
}

func construirFiltroEnvio(filtro dto.FiltroEnvio) bson.M {
	filtroGenerado := bson.M{}

	// Filtrar por estado
	if filtro.Estado != "" {
		filtroGenerado["estado"] = filtro.Estado
	}

	// Filtrar por patente del camión (asegurando que coincida con el formato en MongoDB)
	if filtro.PatenteCamion != "" {
		filtroGenerado["patente_camion"] = filtro.PatenteCamion
	}

	// Filtrar por última parada (dentro del array `paradas`)
	if filtro.UltimaParada != "" {
		filtroGenerado["paradas.ciudad"] = filtro.UltimaParada
	}

	// Filtrar por rango de fechas en `fecha_creacion`
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

	return filtroGenerado
}

// REPORTES
func (repository EnvioRepository) ObtenerCantidadEnviosPorEstado(estado model.EstadoEnvio) (int, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Envios")

	filtro := bson.M{"estado": estado}

	cantidad, err := collection.CountDocuments(context.Background(), filtro)

	if err != nil {
		return 0, err
	}

	return int(cantidad), nil
}
