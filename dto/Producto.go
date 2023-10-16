package dto

import "time"

type Producto struct {
	ID                string
	Tipo              string
	Nombre            string
	PesoUnitario      int
	PrecioUnitario    int
	StockMinimo       int
	StockActual       int
	FechaCreacion     time.Time
	FechaModificacion time.Time
}
