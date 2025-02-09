package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Camion struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Patente           string             `bson:"patente"`
	PesoMaximo        int                `bson:"peso_maximo"`
	CostoKm           int                `bson:"costo_km"`
	FechaCreacion     time.Time          `bson:"fecha_creacion,omitempty"`
	FechaModificacion time.Time          `bson:"fecha_modificacion,omitempty"`
}
