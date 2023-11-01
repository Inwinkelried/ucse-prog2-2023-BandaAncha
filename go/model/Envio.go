package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Envio struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	PatenteCamion     string             `bson:"patente_camion"`
	Pedidos           []string           `bson:"pedidos"`
	Paradas           []Parada           `bson:"paradas"`
	Estado            string             `bson:"estado"`
	FechaCreacion     time.Time          `bson:"fecha_creacion"`
	FechaModificacion time.Time          `bson:"fecha_modificacion"`
}
