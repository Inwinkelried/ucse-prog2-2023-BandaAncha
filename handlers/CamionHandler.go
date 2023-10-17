package handlers

import (
	"log"
	"net/http"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/services"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
	"github.com/gin-gonic/gin"
)

type CamionHandler struct {
	camionService services.CamionInterface
}

func NewCamionHandler(camionService services.CamionInterface) *CamionHandler {
	return &CamionHandler{
		camionService: camionService,
	}
}
func (handler *CamionHandler) ObtenerCamiones(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	camiones := handler.camionService.ObtenerCamiones()
	log.Printf("[handler:AulaHandler][method:ObtenerAulas][cantidad:%d][user:%s]", len(camiones), user.Codigo)
	c.JSON(http.StatusOK, camiones)
}
func (handler *CamionHandler) InsertarCamion(c *gin.Context) {
	var camion dto.Camion
	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.camionService.InsertarCamion(&camion)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *CamionHandler) ModificarCamion(c *gin.Context) {
	var camion dto.Camion
	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	camion.ID = c.Param("id")
	resultado := handler.camionService.ModificarCamion(&camion)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *CamionHandler) EliminarCamion(c *gin.Context) {
	id := c.Param("id")
	camiones := handler.camionService.EliminarCamion(id)
	c.JSON(http.StatusOK, camiones)
}
