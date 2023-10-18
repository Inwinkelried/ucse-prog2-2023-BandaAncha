package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pedido struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Productos         []PedidoProducto   `bson:"productos"`
	Cantidad          int                `bson:"cantidad"`
	FechaCreacion     time.Time          `bson:"fechaCreacion,omitempty"`
	FechaModificacion time.Time          `bson:"fechaModificacion,omitempty"`
	Estado            string             `bson:"estado"`
	Destino           string             `bson:"destino"`
}
