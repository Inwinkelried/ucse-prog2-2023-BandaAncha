package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Tipo              string             `bson:"tipo"`
	Nombre            string             `bson:"nombre"`
	PesoUnitario      int                `bson:"peso_unitario"`
	PrecioUnitario    int                `bson:"precio_unitario"`
	StockMinimo       int                `bson:"stock_minimo"`
	StockActual       int                `bson:"stock_actual"`
	FechaCreacion     time.Time          `bson:"fecha_creacion,omitempty"`
	FechaModificacion time.Time          `bson:"fecha_modificacion,omitempty"`
}
