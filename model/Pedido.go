package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pedido struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Productos         []PedidoProducto   `bson:"productos"`
	FechaCreacion     time.Time          `bson:"fechaCreacion,omitempty"`
	FechaModificacion time.Time          `bson:"fechaModificacion,omitempty"`
	Estado            string             `bson:"Estado"`
	Destino           string             `bson:"destino"`
}
