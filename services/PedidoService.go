package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
)

type PedidoServiceInterface interface {
	ObtenerPedidos() []*dto.Pedido
	InsertarPedido(pedido *dto.Pedido) bool
	AceptarPedido(pedido *dto.Pedido) bool
	CancelarPedido(pedido *dto.Pedido) bool
	ParaEnviarPedido(pedido *dto.Pedido) bool
	EnviadoPedido(pedido *dto.Pedido) bool
}

type PedidoService struct {
	pedidoRepository repositories.PedidoRepositoryInterface
}

func NewPedidoService(pedidoRepository repositories.PedidoRepositoryInterface) *PedidoService {
	return &PedidoService{
		pedidoRepository: pedidoRepository,
	}
}
func (service PedidoService) ObtenerPedidos() []*dto.Pedido {
	pedidosDB, _ := service.pedidoRepository.ObtenerPedidos()
	var pedidos []*dto.Pedido
	for _, pedidodDB := range pedidosDB {
		pedido := dto.NewPedido(pedidodDB)
		pedidos = append(pedidos, pedido)
	}
	return service.ObtenerPedidos()
}
func (service PedidoService) InsertarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.InsertarPedido(pedido.GetModel())
	return true
}
func (service PedidoService) AceptarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.AceptarPedido(pedido.GetModel())
	return true
}
func (service PedidoService) CancelarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.CancelarPedido(pedido.GetModel())
	return true
}
func (service PedidoService) EnviadoPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.EnviadoPedido(pedido.GetModel())
	return true
}
func (service PedidoService) ParaEnviarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.EnviadoPedido(pedido.GetModel())
	return true
}
