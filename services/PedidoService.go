package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
)

type PedidoServiceInterface interface {
	ObtenerPedidos() []*dto.Pedido
	InsertarPedido(pedido *dto.Pedido) bool

	ObtenerPedidosAprobados() []*dto.Pedido
}

type PedidoService struct {
	pedidoRepository repositories.PedidoRepositoryInterface
}

func NewPedidoService(pedidoRepository repositories.PedidoRepositoryInterface) *PedidoService {
	return &PedidoService{
		pedidoRepository: pedidoRepository,
	}
}
func (service PedidoService) ObtenerPedidosAprobados() []*dto.Pedido {
	pedidosDB, _ := service.pedidoRepository.ObtenerPedidosAprobados()
	var pedidos []*dto.Pedido
	for _, pedidodDB := range pedidosDB {
		pedido := dto.NewPedido(pedidodDB)
		pedidos = append(pedidos, pedido)
	}
	return pedidos
}
func (service PedidoService) ObtenerPedidos() []*dto.Pedido {
	pedidosDB, _ := service.pedidoRepository.ObtenerPedidos()
	var pedidos []*dto.Pedido
	for _, pedidodDB := range pedidosDB {
		pedido := dto.NewPedido(pedidodDB)
		pedidos = append(pedidos, pedido)
	}
	return pedidos
}
func (service PedidoService) InsertarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.InsertarPedido(pedido.GetModel())
	return true
}
