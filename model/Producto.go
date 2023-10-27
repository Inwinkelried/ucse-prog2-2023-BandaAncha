package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Tipo              string             `bson:"tipo"`
	Nombre            string             `bson:"nombre"`
	PesoUnitario      int                `bson:"pesoUnitario"`
	PrecioUnitario    int                `bson:"precioUnitario"`
	StockMinimo       int                `bson:"stockMinimo"`
	StockActual       int                `bson:"stockActual"`
	FechaCreacion     time.Time          `bson:"fechaCreacion,omitempty"`
	FechaModificacion time.Time          `bson:"fechaModificacion,omitempty"`
}
