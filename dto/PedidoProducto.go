package dto

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
)

type PedidoProducto struct {
	ID             string
	CodigoProducto string
	Tipo           string
	Nombre         string
	PesoUnitario   float64
	PrecioUnitario float64
	Cantidad       float64
}

func NewPedidoProducto(pedidoProducto model.PedidoProducto) *PedidoProducto {
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
