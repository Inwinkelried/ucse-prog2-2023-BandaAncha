package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	CodigoProducto    string             `bson:"codigoProducto"`
	Tipo              string             `bson:"tipo"`
	Nombre            string             `bson:"nombre"`
	PesoUnitario      float64            `bson:"pesoUnitario"`
	PrecioUnitario    float64            `bson:"precioUnitario"`
	StockMinimo       float64            `bson:"stockMinimo"`
	StockActual       float64            `bson:"stockActual"`
	FechaCreacion     time.Time          `bson:"fechaCreacion,omitempty"`
	FechaModificacion time.Time          `bson:"fechaModificacion,omitempty"`
}
