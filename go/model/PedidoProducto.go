package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PedidoProducto struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CodigoProducto string             `bson:"codigoProducto"`
	PesoUnitario   int                `bson:"pesoUnitario"`
	PrecioUnitario int                `bson:"precioUnitario"`
	Cantidad       int                `bson:"cantidad"`
}

func (productoPedido PedidoProducto) ObtenerPesoProductoPedido() int {
	return productoPedido.PesoUnitario * productoPedido.Cantidad
}
