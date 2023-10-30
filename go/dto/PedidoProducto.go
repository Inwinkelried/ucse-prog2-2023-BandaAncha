package dto

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type PedidoProducto struct {
	ID             string
	CodigoProducto string
	Tipo           string
	Nombre         string
	PesoUnitario   int
	PrecioUnitario int
	Cantidad       int
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
