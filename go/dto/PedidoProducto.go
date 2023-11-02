package dto

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type PedidoProducto struct {
	ID             string `json:"id,omitempty"`
	CodigoProducto string `json:"codigo_producto"`
	PrecioUnitario int    `json:"precio_unitario"`
	PesoUnitario   int    `json:"peso_unitario"`
	Cantidad       int    `json:"cantidad"`
}

func NewPedidoProducto(pedidoProducto *model.PedidoProducto) *PedidoProducto {
	return &PedidoProducto{
		ID:             utils.GetStringIDFromObjectID(pedidoProducto.ID),
		CodigoProducto: pedidoProducto.CodigoProducto,
		PesoUnitario:   pedidoProducto.PesoUnitario,
		PrecioUnitario: pedidoProducto.PrecioUnitario,
		Cantidad:       pedidoProducto.Cantidad,
	}
}

func (pedidoProducto PedidoProducto) GetModel() model.PedidoProducto {
	return model.PedidoProducto{
		ID:             utils.GetObjectIDFromStringID(pedidoProducto.ID),
		CodigoProducto: pedidoProducto.CodigoProducto,
		PesoUnitario:   pedidoProducto.PesoUnitario,
		PrecioUnitario: pedidoProducto.PrecioUnitario,
		Cantidad:       pedidoProducto.Cantidad,
	}
}
