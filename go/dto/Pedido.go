package dto

import (
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type Pedido struct {
	ID                string
	Productos         []PedidoProducto
	Destino           string
	Estado            string
	FechaCreacion     time.Time
	FechaModificacion time.Time
}

func NewPedido(pedido model.Pedido) *Pedido {
	return &Pedido{
		ID:                utils.GetStringIDFromObjectID(pedido.ID),
		Productos:         []PedidoProducto{},
		Destino:           pedido.Destino,
		Estado:            "Pendiente",
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
	}
}
func (pedido Pedido) GetModel() model.Pedido {
	return model.Pedido{
		ID:                utils.GetObjectIDFromStringID(pedido.ID),
		Productos:         pedido.getProductosElegidos(),
		Destino:           pedido.Destino,
		Estado:            pedido.Estado,
		FechaCreacion:     pedido.FechaCreacion,
		FechaModificacion: pedido.FechaModificacion,
	}
}

func (pedido Pedido) getProductosElegidos() []model.PedidoProducto {
	var productosElegidos []model.PedidoProducto
	for _, producto := range pedido.Productos {
		productosElegidos = append(productosElegidos, producto.GetModel())
	}
	return productosElegidos
}
func NewProductosPedido(productosElegidos []model.PedidoProducto) []PedidoProducto {
	var productosElegidosDto []PedidoProducto
	for _, producto := range productosElegidos {
		productosElegidosDto = append(productosElegidosDto, *NewPedidoProducto(&producto))
	}
	return productosElegidosDto
}
