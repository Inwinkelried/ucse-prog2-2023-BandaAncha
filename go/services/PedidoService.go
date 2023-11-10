package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/repositories"
)

type PedidoServiceInterface interface {
	ObtenerPedidos() []*dto.Pedido
	InsertarPedido(pedido *dto.Pedido) bool
	AceptarPedido(pedido *dto.Pedido) bool
	CancelarPedido(pedido *dto.Pedido) bool
	ParaEnviarPedido(pedido *dto.Pedido) bool
	EnviadoPedido(pedido *dto.Pedido) bool
	ObtenerPedidosAprobados() []*dto.Pedido
	ObtenerPedidosFiltrados(filtro dto.FiltroPedido) ([]dto.Pedido, error)
	ObtenerPedidoPorID(pedido *dto.Pedido) (*dto.Pedido, error)
}

type PedidoService struct {
	envioRepository    repositories.EnvioRepositoryInterface
	camionRepository   repositories.CamionRepositoryInterface
	pedidoRepository   repositories.PedidoRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
}

func NewPedidoService(envioRepository repositories.EnvioRepositoryInterface, camionRepository repositories.CamionRepositoryInterface, pedidoRepository repositories.PedidoRepositoryInterface, productoRepository repositories.ProductoRepositoryInterface) *PedidoService {
	return &PedidoService{
		envioRepository:    envioRepository,
		camionRepository:   camionRepository,
		pedidoRepository:   pedidoRepository,
		productoRepository: productoRepository,
	}
}
func (service PedidoService) ObtenerPedidoPorID(pedidoConID *dto.Pedido) (*dto.Pedido, error) {
	pedidoDB, err := service.pedidoRepository.ObtenerPedidoPorID(pedidoConID.GetModel())

	var pedido *dto.Pedido

	if err != nil {
		return nil, err
	} else {
		pedido = dto.NewPedido(pedidoDB)
	}

	return pedido, nil
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

func (service PedidoService) InsertarPedido(pedidoACrear *dto.Pedido) bool {

	service.pedidoRepository.InsertarPedido(pedidoACrear.GetModel())
	return true
}

func (service PedidoService) AceptarPedido(pedido *dto.Pedido) bool {
	service.pedidoRepository.AceptarPedido(pedido.GetModel())
	return true
}
func (service PedidoService) CancelarPedido(pedido *dto.Pedido) bool {
	if pedido.Estado == "" { //Dejo el "" porque mi BDD tiene "" en vez de Pendiente
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
func (service *PedidoService) ObtenerPedidosFiltrados(filtro dto.FiltroPedido) ([]dto.Pedido, error) {
	pedidos, err := service.pedidoRepository.ObtenerPedidosFiltrados(filtro)
	var pedido *dto.Pedido
	var pedidosDTO []dto.Pedido
	if err != nil {
		return nil, err
	} else {
		for _, pedidoDB := range pedidos {
			pedido = dto.NewPedido(pedidoDB)
			pedidosDTO = append(pedidosDTO, *pedido)
		}
	}
	return pedidosDTO, nil
}
