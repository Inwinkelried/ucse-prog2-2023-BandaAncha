package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	//"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/repositories"
	//"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type PedidoServiceInterface interface {
	ObtenerPedidos() []*dto.Pedido
	InsertarPedido(pedido *dto.Pedido) bool
	AceptarPedido(pedido *dto.Pedido) bool
	CancelarPedido(pedido *dto.Pedido) bool
	ParaEnviarPedido(pedido *dto.Pedido) bool
	EnviadoPedido(pedido *dto.Pedido) bool
	ObtenerPedidosAprobados() []*dto.Pedido
}

type PedidoService struct {
	envioRepository    repositories.EnvioRepositoryInterface
	camionRepository   repositories.CamionRepositoryInterface
	pedidoRepository   repositories.PedidoRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
}

func NewPedidoService(pedidoRepository repositories.PedidoRepositoryInterface) *PedidoService {
	return &PedidoService{
		pedidoRepository: pedidoRepository,
	}
}

// Necesito un metodo similar a agregar parada de envio, pero que agregue productos al pedido
// ACA HAY QUE HACER EL METODO PARA AGREGAR PRODUCTOS AL PEDIDO
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

func (service PedidoService) InsertarPedido(pedidoACrear *dto.Pedido) bool {

	service.pedidoRepository.InsertarPedido(pedidoACrear.GetModel())
	return true
}

func (service PedidoService) AceptarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.AceptarPedido(pedido.GetModel())
	return true
}
func (service PedidoService) CancelarPedido(pedido *dto.Pedido) bool {
	if pedido.Estado == "Pendiente" {
		service.pedidoRepository.CancelarPedido(pedido.GetModel())
		return true
	} else {
		return false
	}
}
func (service PedidoService) EnviadoPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.EnviadoPedido(pedido.GetModel())
	return true
}
func (service PedidoService) ParaEnviarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.ParaEnviarPedido(pedido.GetModel())
	return true
}
