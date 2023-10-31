package dto

import (
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type Producto struct {
	ID                string    `json:"id,omitempty"`
	Tipo              string    `json:"tipo"`
	Nombre            string    `json:"nombre"`
	PesoUnitario      int       `json:"peso_unitario"`
	PrecioUnitario    int       `json:"precio_unitario"`
	StockMinimo       int       `json:"stock_minimo"`
	StockActual       int       `json:"stock_actual"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}

func NewProducto(producto model.Producto) *Producto {
	return &Producto{
		ID: utils.GetStringIDFromObjectID(producto.ID),

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
		ID:   utils.GetObjectIDFromStringID(producto.ID),
		Tipo: producto.Tipo,

		Nombre:            producto.Nombre,
		PesoUnitario:      producto.PesoUnitario,
		PrecioUnitario:    producto.PrecioUnitario,
		StockMinimo:       producto.StockMinimo,
		StockActual:       producto.StockActual,
		FechaCreacion:     producto.FechaCreacion,
		FechaModificacion: producto.FechaModificacion,
	}
}
