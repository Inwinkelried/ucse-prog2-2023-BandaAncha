package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
)

type EnvioServiceInterface interface {
	ObtenerEnvios() []*dto.Envio
	InsertarEnvio(envio *dto.Envio) bool
	ModificarEnvio(envio *dto.Envio) bool
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
func (service *EnvioService) ModificarEnvio(envio *dto.Envio) bool {
	service.envioRepository.ModificarEnvio(envio.GetModel())
	return true
}
