package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
)

type EnvioServiceInterface interface {
	ObtenerEnvios() []*dto.Envio
	InsertarEnvio(envio *dto.Envio) bool
	EnRutaEnvio(envio *dto.Envio) bool
	DespachadoEnvio(envio *dto.Envio) bool
}
type EnvioService struct {
	envioRepository repositories.EnvioRepositoryInterface
}

func NewEnvioService(envioRepository repositories.EnvioRepositoryInterface) *EnvioService {
	return &EnvioService{
		envioRepository: envioRepository,
	}
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
