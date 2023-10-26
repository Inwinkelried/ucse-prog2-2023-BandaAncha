package dto

import (
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
)

type Producto struct {
	ID                string
	CodigoProducto    string
	Tipo              string
	Nombre            string
	PesoUnitario      int
	PrecioUnitario    int
	StockMinimo       int
	StockActual       int
	FechaCreacion     time.Time
	FechaModificacion time.Time
}

func NewProducto(producto model.Producto) *Producto {
	return &Producto{
		ID:                utils.GetStringIDFromObjectID(producto.ID),
		CodigoProducto:    producto.CodigoProducto,
		Tipo:              producto.Tipo,
		Nombre:            producto.Nombre,
		PesoUnitario:      producto.PesoUnitario,
		PrecioUnitario:    producto.PrecioUnitario,
		StockMinimo:       producto.StockMinimo,
		StockActual:       producto.StockActual,
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
	}
}
func (producto Producto) GetModel() model.Producto {
	return model.Producto{
		ID:                utils.GetObjectIDFromStringID(producto.ID),
		Tipo:              producto.Tipo,
		CodigoProducto:    producto.CodigoProducto,
		Nombre:            producto.Nombre,
		PesoUnitario:      producto.PesoUnitario,
		PrecioUnitario:    producto.PrecioUnitario,
		StockMinimo:       producto.StockMinimo,
		StockActual:       producto.StockActual,
		FechaCreacion:     producto.FechaCreacion,
		FechaModificacion: producto.FechaModificacion,
	}
}
