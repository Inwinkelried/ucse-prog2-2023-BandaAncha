package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/repositories"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type EnvioServiceInterface interface {
	ObtenerEnvios() ([]*dto.Envio, error)
	ObtenerEnvioPorID(envio *dto.Envio) (*dto.Envio, error)
	InsertarEnvio(envio *dto.Envio) (bool, error)
	EnRutaEnvio(envio *dto.Envio) (bool, error)
	DespachadoEnvio(envio *dto.Envio) (bool, error)
	AgregarParada(envio *dto.Envio) (bool, error)
	ObtenerEnviosFiltrados(filtro dto.FiltroEnvio) ([]dto.Envio, error)
}
type EnvioService struct {
	envioRepository    repositories.EnvioRepositoryInterface
	camionRepository   repositories.CamionRepositoryInterface
	pedidoRepository   repositories.PedidoRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface, camionRepository repositories.CamionRepositoryInterface, pedidoRepository repositories.PedidoRepositoryInterface, productoRepository repositories.ProductoRepositoryInterface) *EnvioService {
	return &EnvioService{
		envioRepository:    envioRepository,
		camionRepository:   camionRepository,
		pedidoRepository:   pedidoRepository,
		productoRepository: productoRepository,
	}
}
func (service *EnvioService) AgregarParada(envio *dto.Envio) (bool, error) {
	if envio.Estado != "Despachado" {
		envioDB, err := service.envioRepository.ObtenerEnvioPorID(envio.GetModel())
		if err != nil {
			return false, err
		}
		envioDB.Paradas = append(envioDB.Paradas, envio.Paradas[0].GetModel())
		service.envioRepository.ActualizarEnvio(envioDB)
		//Actualizamos el envio en la base de datos, que ahora tiene la nueva parada
		return true, err
	} else {
		return false, nil
	}

}
func (service *EnvioService) ObtenerEnvioPorID(envioConID *dto.Envio) (*dto.Envio, error) {
	envioDB, err := service.envioRepository.ObtenerEnvioPorID(envioConID.GetModel())
	if err != nil {
		return nil, err
	}
	var envio = dto.NewEnvio(envioDB)
	return envio, nil
}

func (service *EnvioService) ObtenerEnvios() ([]*dto.Envio, error) {
	enviosDB, err := service.envioRepository.ObtenerEnvios()
	if err != nil {
		return nil, err
	}
	var envios []*dto.Envio
	for _, enviosDB := range enviosDB {
		envio := dto.NewEnvio(enviosDB)
		envios = append(envios, envio)
	}
	return envios, nil
}

func (service *EnvioService) VerificarPesoEnvio(envio *dto.Envio) (bool, error) {
	camionConID := model.Camion{Patente: envio.PatenteCamion} // Busco el camion por patente
	//aca ocurre el error
	camionEncontrado, err := service.camionRepository.ObtenerCamionPorPatente(camionConID)
	if err != nil {
		return false, err
	}
	if camionEncontrado.Patente == "" { // Si no encuentra el camion...
		return false, nil
	}
	var pesototal int = 0
	for _, pedido := range envio.Pedidos {
		pedidoAFiltrar := model.Pedido{ID: utils.GetObjectIDFromStringID(pedido)}
		pedidoFiltrado, err := service.pedidoRepository.ObtenerPedidoPorID(pedidoAFiltrar)
		if err != nil {
			return false, err
		}
		pesoPedido, err := service.pedidoRepository.ObtenerPesoPedido(pedidoFiltrado)
		pesototal = pesototal + pesoPedido
	}
	if pesototal <= camionEncontrado.PesoMaximo {
		return true, nil
	} else {
		return false, nil
	}
}

func (service *EnvioService) InsertarEnvio(envio *dto.Envio) (bool, error) {

	camionEsValido, err := service.VerificarPesoEnvio(envio)
	if err != nil {
		return false, err
	}
	if camionEsValido {
		service.envioRepository.InsertarEnvio(envio.GetModel())
		return true, nil
	} else {
		return false, err
	}
}
func (service *EnvioService) EnRutaEnvio(envio *dto.Envio) (bool, error) {
	envioParaActualizar, err := service.ObtenerEnvioPorID(envio)
	if err != nil {
		return false, err
	}
	envioParaActualizar.Estado = "En Ruta"
	service.envioRepository.ActualizarEnvio(envioParaActualizar.GetModel())
	return true, nil
}

func (service *EnvioService) DespachadoEnvio(envio *dto.Envio) (bool, error) {
	envioParaActualizar, err := service.ObtenerEnvioPorID(envio)
	if err != nil {
		return false, err
	}
	envioParaActualizar.Estado = "Despachado"
	for _, pedido := range envioParaActualizar.Pedidos {
		pedidoAFiltrar := model.Pedido{ID: utils.GetObjectIDFromStringID(pedido)}
		pedidoFiltrado, err := service.pedidoRepository.ObtenerPedidoPorID(pedidoAFiltrar)
		if err != nil {
			return false, err
		}
		for _, productoPedido := range pedidoFiltrado.Productos {
			productoADescontar := model.Producto{ID: utils.GetObjectIDFromStringID(productoPedido.CodigoProducto), StockActual: productoPedido.Cantidad} //Preguntarle a Maxi si esto esta bien
			_, err := service.productoRepository.DescontarStockProducto(productoADescontar)
			if err != nil {
				return false, err
			}
		}
	}
	service.envioRepository.ActualizarEnvio(envioParaActualizar.GetModel())
	return true, nil
}
func (service *EnvioService) ObtenerEnviosFiltrados(filtro dto.FiltroEnvio) ([]dto.Envio, error) {
	envios, err := service.envioRepository.ObtenerEnviosFiltrados(filtro)
	var envio *dto.Envio
	var enviosDTO []dto.Envio
	if err != nil {
		return nil, err
	} else {
		for _, envioDB := range envios {
			envio = dto.NewEnvio(envioDB)
			enviosDTO = append(enviosDTO, *envio)
		}
	}
	return enviosDTO, nil
}
