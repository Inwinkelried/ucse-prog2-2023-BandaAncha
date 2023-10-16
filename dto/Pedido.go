package dto

import "time"

type Pedido struct {
	ID                string
	Productos         []Producto
	Usuario           Usuario
	FechaCreacion     time.Time
	FechaModificacion time.Time
}
