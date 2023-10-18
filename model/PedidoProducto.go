package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PedidoProducto struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CodigoProducto string             `bson:"codigoProducto"`
	Tipo           string             `bson:"tipo"`
	Nombre         string             `bson:"nombre"`
	PesoUnitario   float64            `bson:"pesoUnitario"`
	PrecioUnitario float64            `bson:"precioUnitario"`
	Cantidad       float64            `bson:"cantidad"`
}
