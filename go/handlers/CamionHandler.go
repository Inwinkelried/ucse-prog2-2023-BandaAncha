package handlers

import (
	"log"
	"net/http"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/services"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
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
func (handler *CamionHandler) ObtenerCamionPorID(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	//invocamos al metodo
	camion, err := handler.camionService.ObtenerCamionPorID(&dto.Camion{ID: id})
	if err != nil {
		log.Printf("[handler:CamionHandler][method:ObtenerCamionPorId][camion:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, camion)
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
