package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
)

type CamionInterface interface {
	ObtenerCamiones() []*dto.Camion
	InsertarCamion(camion *dto.Camion) bool
	EliminarCamion(id string) bool
	ModificarCamion(camion *dto.Camion) bool
}
type CamionService struct {
	camionRepository repositories.CamionRepositoryInterface
}

func NewCamionService(camionRepository repositories.CamionRepositoryInterface) *CamionService {
	return &CamionService{
		camionRepository: camionRepository,
	}
}
func (service *CamionService) ObtenerCamiones() []*dto.Camion {
	camionesDB, _ := service.camionRepository.ObtenerCamiones()
	var camiones []*dto.Camion
	for _, camionesDB := range camionesDB {
		camion := dto.NewCamion(camionesDB)
		camiones = append(camiones, camion)
	}
	return camiones
}
func (service *CamionService) InsertarCamion(camion *dto.Camion) bool {
	service.camionRepository.InsertarCamion(camion.GetModel())
	return true
}
func (service *CamionService) ModificarCamion(camion *dto.Camion) bool {
	service.camionRepository.ModificarCamion(camion.GetModel())
	return true
}
func (service *CamionService) EliminarCamion(id string) bool {
	service.camionRepository.EliminarCamion(utils.GetObjectIDFromStringID(id))
	return true
}
