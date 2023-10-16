package dto

import "time"

type Usuario struct {
	ID                string
	Nombre            string
	Pedidos           []Pedido
	FechaCreacion     time.Time
	FechaModificacion time.Time
}
