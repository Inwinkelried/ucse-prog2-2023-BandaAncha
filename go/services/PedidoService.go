package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/repositories"
)

type PedidoServiceInterface interface {
	ObtenerPedidos() ([]*dto.Pedido, error)
	InsertarPedido(pedido *dto.Pedido) (bool, error)
	AceptarPedido(pedido *dto.Pedido) (bool, error)
	CancelarPedido(pedido *dto.Pedido) (bool, error)
	ParaEnviarPedido(pedido *dto.Pedido) (bool, error)
	EnviadoPedido(pedido *dto.Pedido) (bool, error)
	ObtenerPedidosAprobados() ([]*dto.Pedido, error)
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

func (service PedidoService) ObtenerPedidosAprobados() ([]*dto.Pedido, error) {
	pedidosDB, err := service.pedidoRepository.ObtenerPedidosAprobados()
	if err != nil {
		return nil, err
	}
	var pedidos []*dto.Pedido
	for _, pedidodDB := range pedidosDB {
		pedido := dto.NewPedido(pedidodDB)
		pedidos = append(pedidos, pedido)
	}
	return pedidos, err
}
func (service PedidoService) ObtenerPedidos() ([]*dto.Pedido, error) {
	pedidosDB, err := service.pedidoRepository.ObtenerPedidos()
	if err != nil {
		return nil, err
	}
	var pedidos []*dto.Pedido
	for _, pedidodDB := range pedidosDB {
		pedido := dto.NewPedido(pedidodDB)
		pedidos = append(pedidos, pedido)
	}
	return pedidos, err
}

func (service PedidoService) InsertarPedido(pedidoACrear *dto.Pedido) (bool, error) {

	resultado, err := service.pedidoRepository.InsertarPedido(pedidoACrear.GetModel())
	if resultado == nil {
		return false, err
	}
	if err != nil {

		return false, err
	}
	return true, err
}

func (service PedidoService) AceptarPedido(pedido *dto.Pedido) (bool, error) {
	resultado, err := service.pedidoRepository.AceptarPedido(pedido.GetModel())
	if resultado == nil {
		return false, err
	}
	if err != nil {
		return false, err
	}

	return true, err
}

func (service PedidoService) CancelarPedido(pedido *dto.Pedido) (bool, error) {
	if pedido.Estado == "" { //Dejo el "" porque mi BDD tiene "" en vez de Pendiente
		resultado, err := service.pedidoRepository.CancelarPedido(pedido.GetModel())
		if resultado == nil {
			return false, err
		}
		if err != nil {
			return false, err
		}
		return true, err
	} else {
		return false, nil
	}
}
func (service PedidoService) EnviadoPedido(pedido *dto.Pedido) (bool, error) {
	resultado, err := service.pedidoRepository.EnviadoPedido(pedido.GetModel())
	if resultado == nil {
		return false, err
	}
	if err != nil {
		return false, err
	}

	return true, err
}
func (service PedidoService) ParaEnviarPedido(pedido *dto.Pedido) (bool, error) {
	resultado, err := service.pedidoRepository.ParaEnviarPedido(pedido.GetModel())
	if resultado == nil {
		return false, err
	}
	if err != nil {
		return false, err
	}

	return true, err
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
