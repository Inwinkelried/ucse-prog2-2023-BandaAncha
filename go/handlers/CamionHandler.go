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
	camiones, err := handler.camionService.ObtenerCamiones()
	if err != nil {
		log.Printf("[handler:CamionHandler][method:ObtenerCamiones][error:%s][user:%s]", err.Error(), user.Codigo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:CamionHandler][method:ObtenerCamiones][camiones:%v][cantidad:%d][user:%s]", camiones, len(camiones), user.Codigo)
	c.JSON(http.StatusOK, camiones)
}
func (handler *CamionHandler) ObtenerCamionPorID(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	camion, err := handler.camionService.ObtenerCamionPorID(&dto.Camion{ID: id})
	if err != nil {
		log.Printf("[handler:CamionHandler][method:ObtenerCamionPorId][camion:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, camion)
}
func (handler *CamionHandler) InsertarCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	var camion dto.Camion
	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error en el formato": err.Error()})
		return
	}
	if err := handler.camionService.InsertarCamion(&camion); err != nil {
		log.Printf("[handler:CamionHandler][method:CrearCamion][envio:%+v][user:%s]", err.Error(), user.Codigo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:CamionHandler][method:CrearCamion][camion:%+v][user:%s]", camion, user.Codigo)
	c.JSON(http.StatusCreated, gin.H{"status": "Creado correctamente"})
}
func (handler *CamionHandler) ModificarCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	var camion dto.Camion
	if err := c.ShouldBindJSON(&camion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	camion.ID = c.Param("id")
	if err := handler.camionService.ModificarCamion(&camion); err != nil {
		log.Printf("[handler:CamionHandler][method:CrearCamion][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:CamionHandler][method:ModificarCamion][camion:%+v][user:%s]", camion, user.Codigo)
	c.JSON(http.StatusCreated, gin.H{"status": "Modificado correctamente"})
}
func (handler *CamionHandler) EliminarCamion(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	if err := handler.camionService.EliminarCamion(id); err != nil {
		log.Printf("[handler:CamionHandler][method:CrearCamion][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Eliminado correctamente"})
}
