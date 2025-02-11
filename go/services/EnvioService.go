package services

import (
	"fmt"
	"log"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/model"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/repositories"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type EnvioServiceInterface interface {
	ObtenerEnvioPorID(envio *dto.Envio) (*dto.Envio, error)
	InsertarEnvio(envio *dto.Envio) (bool, error)
	EnRutaEnvio(envio *dto.Envio) (bool, error)
	DespachadoEnvio(envio *dto.Envio) (bool, error)
	AgregarParada(envio *dto.Envio) (bool, error)
	ObtenerEnvios(filtro dto.FiltroEnvio) ([]dto.Envio, error)
	ObtenerCantidadEnviosPorEstado() ([]utils.CantidadEstado, error)
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
	// Validar que el envío tenga los campos requeridos
	if envio.PatenteCamion == "" {
		return false, fmt.Errorf("la patente del camión es requerida")
	}
	if len(envio.Pedidos) == 0 {
		return false, fmt.Errorf("debe incluir al menos un pedido")
	}
	if envio.Estado == "" {
		envio.Estado = string(model.ADespachar)
	} else if envio.Estado != string(model.ADespachar) {
		return false, fmt.Errorf("el estado inicial del envío debe ser A Despachar")
	}

	// Verificar que el camión exista y tenga capacidad
	camionConID := model.Camion{Patente: envio.PatenteCamion}
	camionEncontrado, err := service.camionRepository.ObtenerCamionPorPatente(camionConID)
	if err != nil {
		return false, fmt.Errorf("error al buscar el camión: %v", err)
	}
	if camionEncontrado.Patente == "" {
		return false, fmt.Errorf("no se encontró el camión con patente %s", envio.PatenteCamion)
	}

	// Calcular peso total de los pedidos
	var pesoTotal int = 0
	for _, pedidoID := range envio.Pedidos {
		pedidoModel := model.Pedido{ID: utils.GetObjectIDFromStringID(pedidoID)}
		pedido, err := service.pedidoRepository.ObtenerPedidoPorID(pedidoModel)
		if err != nil {
			return false, fmt.Errorf("error al obtener el pedido %s: %v", pedidoID, err)
		}

		// Verificar que el pedido esté en estado ACEPTADO
		if pedido.Estado != string(model.Aceptado) {
			return false, fmt.Errorf("el pedido %s debe estar en estado Aceptado", pedidoID)
		}

		// Calcular peso del pedido
		pesoPedido, err := service.pedidoRepository.ObtenerPesoPedido(pedido)
		if err != nil {
			return false, fmt.Errorf("error al calcular peso del pedido %s: %v", pedidoID, err)
		}
		pesoTotal += pesoPedido
	}

	// Verificar que no exceda el peso máximo del camión
	if pesoTotal > camionEncontrado.PesoMaximo {
		return false, fmt.Errorf("el peso total de los pedidos (%d) excede la capacidad del camión (%d)", pesoTotal, camionEncontrado.PesoMaximo)
	}

	// Crear el envío
	_, err = service.envioRepository.InsertarEnvio(envio.GetModel())
	if err != nil {
		return false, fmt.Errorf("error al insertar el envío en la base de datos: %v", err)
	}

	// Actualizar estado de los pedidos a PARA ENVIAR
	for _, pedidoID := range envio.Pedidos {
		pedidoModel := model.Pedido{ID: utils.GetObjectIDFromStringID(pedidoID)}
		pedido, _ := service.pedidoRepository.ObtenerPedidoPorID(pedidoModel)
		pedido.Estado = string(model.ParaEnviar)
		_, err := service.pedidoRepository.ActualizarPedido(pedido)
		if err != nil {
			// Aquí deberíamos hacer rollback del envío creado, pero por ahora solo logueamos el error
			log.Printf("Error al actualizar estado del pedido %s: %v", pedidoID, err)
		}
	}

	return true, nil
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
func (service *EnvioService) ObtenerEnvios(filtro dto.FiltroEnvio) ([]dto.Envio, error) {
	enviosDB, err := service.envioRepository.ObtenerEnvios(filtro)
	if err != nil {
		return nil, err
	}
	enviosDTO := convertirEnviosADTO(enviosDB)
	return enviosDTO, nil
}
func convertirEnviosADTO(envios []model.Envio) []dto.Envio {
	var enviosDTO []dto.Envio
	for _, envioDB := range envios {
		envio := dto.NewEnvio(envioDB)
		enviosDTO = append(enviosDTO, *envio)
	}
	return enviosDTO
}

// reportes
func (service *EnvioService) ObtenerCantidadEnviosPorEstado() ([]utils.CantidadEstado, error) {
	//Por cada estado posible de envio, obtengo la cantidad de envios en ese estado
	cantidadEnviosADespachar, err := service.envioRepository.ObtenerCantidadEnviosPorEstado(model.ADespachar)

	if err != nil {
		return nil, err
	}

	cantidadEnviosEnRuta, err := service.envioRepository.ObtenerCantidadEnviosPorEstado(model.EnRuta)

	if err != nil {
		return nil, err
	}

	cantidadEnviosDespachados, err := service.envioRepository.ObtenerCantidadEnviosPorEstado(model.Despachado)

	if err != nil {
		return nil, err
	}

	//Agrego los resultados a un array de CantidadEstado
	cantidadEnviosPorEstados := []utils.CantidadEstado{
		{Estado: string(model.ADespachar), Cantidad: cantidadEnviosADespachar},
		{Estado: string(model.EnRuta), Cantidad: cantidadEnviosEnRuta},
		{Estado: string(model.Despachado), Cantidad: cantidadEnviosDespachados},
	}

	return cantidadEnviosPorEstados, nil
}
