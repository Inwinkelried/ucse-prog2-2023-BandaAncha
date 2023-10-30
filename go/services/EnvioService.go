package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/repositories"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type EnvioServiceInterface interface {
	ObtenerEnvios() []*dto.Envio
	ObtenerEnvioPorID(envio *dto.Envio) (*dto.Envio, error)
	InsertarEnvio(envio *dto.Envio) bool
	EnRutaEnvio(envio *dto.Envio) bool
	DespachadoEnvio(envio *dto.Envio) (bool, error)
	AgregarParada(envio *dto.Envio) (bool, error)
}
type EnvioService struct {
	envioRepository    repositories.EnvioRepositoryInterface
	camionRepository   repositories.CamionRepositoryInterface
	pedidoRepository   repositories.PedidoRepositoryInterface
	productoRepository repositories.ProductoRepositoryInterface
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface) *EnvioService {
	return &EnvioService{
		envioRepository: envioRepository,
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

	var envio *dto.Envio

	if err != nil {
		return nil, err
	} else {
		envio = dto.NewEnvio(envioDB)
	}

	return envio, nil
}

func (service *EnvioService) ObtenerEnvios() []*dto.Envio {
	enviosDB, _ := service.envioRepository.ObtenerEnvios()
	var envios []*dto.Envio
	for _, enviosDB := range enviosDB {
		envio := dto.NewEnvio(enviosDB)
		envios = append(envios, envio)
	}
	return envios
}

func (service *EnvioService) VerificarPesoEnvio(envio *dto.Envio) (bool, error) {
	camionConID := model.Camion{Patente: envio.IDcamion}
	//aca ocurre el error
	camionEncontrado, err := service.camionRepository.ObtenerCamionPorPatente(camionConID)
	if err != nil {
		return false, err
	}
	if camionEncontrado.Patente == "" {
		return false, nil // No se encontró ningún camión con esa patente
	}
	var pesototal int = 0
	for _, pedido := range envio.Pedidos {
		pedidoAFiltrar := model.Pedido{ID: utils.GetObjectIDFromStringID(pedido)}
		pedidoFiltrado, err := service.pedidoRepository.ObtenerPedidoPorId(pedidoAFiltrar)
		if err != nil {
			return false, err
		}
		pesoPedido, err := service.pedidoRepository.ObtenerPesoPedido(*pedidoFiltrado)
		pesototal = pesototal + pesoPedido
	}
	if pesototal <= camionEncontrado.PesoMaximo {
		return true, nil
	} else {
		return false, nil
	}
}

func (service *EnvioService) InsertarEnvio(envio *dto.Envio) bool { //falta probar

	camionEsValido, err := service.VerificarPesoEnvio(envio)
	if err != nil {
		return false
	}
	if camionEsValido {
		service.envioRepository.InsertarEnvio(envio.GetModel())
		return true
	} else {
		return false
	}
}
func (service *EnvioService) EnRutaEnvio(envio *dto.Envio) bool {
	envio.Estado = "En Ruta"
	service.envioRepository.ActualizarEnvio(envio.GetModel())
	return true
}

func (service *EnvioService) DespachadoEnvio(envio *dto.Envio) (bool, error) { //hay que probar
	envio.Estado = "Despachado"
	// Pedidos := envio.Pedidos
	// for _, pedido := range Pedidos {
	// 	Productos := pedido.Productos
	// 	for _, producto := range Productos {
	// 		productoAFiltrar := model.Producto{ID: utils.GetObjectIDFromStringID(producto.CodigoProducto)}
	// 		productoDB, _ := service.productoRepository.ObtenerProductoPorID(productoAFiltrar)
	// 		productoDB.StockActual = productoDB.StockActual - producto.Cantidad
	// 		service.productoRepository.ModificarProducto(productoDB)
	// 	}
	// }
	service.envioRepository.ActualizarEnvio(envio.GetModel())
	return true, nil
}
