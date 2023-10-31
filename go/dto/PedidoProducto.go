package dto

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type PedidoProducto struct {
	ID             string `json:"id,omitempty"`
	CodigoProducto string `json:"codigo_producto"`
	Tipo           string `json:"tipo"`
	Nombre         string `json:"nombre"`
	PesoUnitario   int    `json:"peso_unitario"`
	PrecioUnitario int    `json:"precio_unitario"`
	Cantidad       int    `json:"cantidad"`
}

func NewPedidoProducto(pedidoProducto *model.PedidoProducto) *PedidoProducto {
	return &PedidoProducto{
		ID:             utils.GetStringIDFromObjectID(pedidoProducto.ID),
		CodigoProducto: pedidoProducto.CodigoProducto,
		Tipo:           pedidoProducto.Tipo,
		Nombre:         pedidoProducto.Nombre,
		PesoUnitario:   pedidoProducto.PesoUnitario,
		PrecioUnitario: pedidoProducto.PrecioUnitario,
		Cantidad:       pedidoProducto.Cantidad,
	}
}

func (pedidoProducto PedidoProducto) GetModel() model.PedidoProducto {
	return model.PedidoProducto{
		ID:             utils.GetObjectIDFromStringID(pedidoProducto.ID),
		CodigoProducto: pedidoProducto.CodigoProducto,
		Tipo:           pedidoProducto.Tipo,
		Nombre:         pedidoProducto.Nombre,
		PesoUnitario:   pedidoProducto.PesoUnitario,
		PrecioUnitario: pedidoProducto.PrecioUnitario,
		Cantidad:       pedidoProducto.Cantidad,
	}
}
