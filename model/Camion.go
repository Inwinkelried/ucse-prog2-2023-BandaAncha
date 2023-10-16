package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Camion struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Patente           string             `bson:"Patente"`
	PesoMaximo        int                `bson:"PesoMaximo"`
	CostoKm           int                `bson:"CostoPorKilometro"`
	FechaCreacion     time.Time          `bson:"FechaCreacion,omitempty"`
	FechaModificacion time.Time          `bson:"FechaModificacion,omitempty"`
}
