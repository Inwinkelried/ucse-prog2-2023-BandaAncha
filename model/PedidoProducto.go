package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PedidoProducto struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CodigoProducto string             `bson:"codigoProducto"`
	Tipo           string             `bson:"tipo"`
	Nombre         string             `bson:"nombre"`
	PesoUnitario   int                `bson:"pesoUnitario"`
	PrecioUnitario int                `bson:"precioUnitario"`
	Cantidad       int                `bson:"cantidad"`
}
