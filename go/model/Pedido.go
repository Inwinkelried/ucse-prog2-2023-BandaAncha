package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pedido struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Productos         []PedidoProducto   `bson:"productos"`
	FechaCreacion     time.Time          `bson:"fecha_creacion,omitempty"`
	FechaModificacion time.Time          `bson:"fecha_modificacion,omitempty"`
	Estado            string             `bson:"estado"`
	Destino           string             `bson:"destino"`
}
