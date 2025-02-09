package dto

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
)

type PedidoProducto struct {
	CodigoProducto string `json:"codigo_producto"`
	PrecioUnitario int    `json:"precio_unitario"`
	PesoUnitario   int    `json:"peso_unitario"`
	Cantidad       int    `json:"cantidad"`
}

func NewPedidoProducto(pedidoProducto *model.PedidoProducto) *PedidoProducto {
	return &PedidoProducto{
		CodigoProducto: pedidoProducto.CodigoProducto,
		PesoUnitario:   pedidoProducto.PesoUnitario,
		PrecioUnitario: pedidoProducto.PrecioUnitario,
		Cantidad:       pedidoProducto.Cantidad,
	}
}

func (pedidoProducto PedidoProducto) GetModel() model.PedidoProducto {
	return model.PedidoProducto{
		CodigoProducto: pedidoProducto.CodigoProducto,
		PesoUnitario:   pedidoProducto.PesoUnitario,
		PrecioUnitario: pedidoProducto.PrecioUnitario,
		Cantidad:       pedidoProducto.Cantidad,
	}
}
