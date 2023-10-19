package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
)

type PedidoServiceInterface interface {
	ObtenerPedidos() []*dto.Pedido
	InsertarPedido(pedido *dto.Pedido) bool
	EliminarPedido(id string) bool
	ModificarPedido(pedido *dto.Pedido) bool
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
func (service PedidoService) ModificarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.ModificarPedido(pedido.GetModel())
	return true
}
func (service PedidoService) EliminarPedido(id string) bool {
	service.pedidoRepository.EliminarPedido(utils.GetObjectIDFromStringID(id))
	return true
}
