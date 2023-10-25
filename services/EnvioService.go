package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
)

type EnvioServiceInterface interface {
	ObtenerEnvios() []*dto.Envio
	ObtenerEnvioPorID(envio *dto.Envio) (*dto.Envio, error)
	InsertarEnvio(envio *dto.Envio) bool
	EnRutaEnvio(envio *dto.Envio) bool
	DespachadoEnvio(envio *dto.Envio) bool
	AgregarParada(envio *dto.Envio) (bool, error)
}
type EnvioService struct {
	envioRepository repositories.EnvioRepositoryInterface
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface) *EnvioService {
	return &EnvioService{
		envioRepository: envioRepository,
	}
}
func (service *EnvioService) AgregarParada(envio *dto.Envio) (bool, error) {
	envioDB, err := service.envioRepository.ObtenerEnvioPorID(envio.GetModel())
	if err != nil {
		return false, err
	}
	envioDB.Paradas = append(envioDB.Paradas, envio.Paradas[0].GetModel())
	service.envioRepository.ActualizarEnvio(envioDB)
	//Actualizamos el envio en la base de datos, que ahora tiene la nueva parada
	return true, err
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
func (service *EnvioService) InsertarEnvio(envio *dto.Envio) bool {
	service.envioRepository.InsertarEnvio(envio.GetModel())
	return true
}
func (service *EnvioService) EnRutaEnvio(envio *dto.Envio) bool {
	envio.Estado = "En Ruta"
	service.envioRepository.ActualizarEnvio(envio.GetModel())
	return true
}
func (service *EnvioService) DespachadoEnvio(envio *dto.Envio) bool {
	envio.Estado = "Despachado"
	service.envioRepository.ActualizarEnvio(envio.GetModel())
	return true
}
