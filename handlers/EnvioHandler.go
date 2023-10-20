package handlers

import (
	"log"
	"net/http"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/services"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
	"github.com/gin-gonic/gin"
)

type EnvioHandler struct {
	envioService services.EnvioServiceInterface
}

func NewEnvioHandler(envioService services.EnvioServiceInterface) *EnvioHandler {
	return &EnvioHandler{
		envioService: envioService,
	}
}
func (handler *EnvioHandler) ObtenerEnvios(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	envios := handler.envioService.ObtenerEnvios()
	log.Printf("[handler:EnvioHandler][method:ObtenerEnvios][cantidad:%d][user:%s]", len(envios), user.Codigo)
	c.JSON(http.StatusOK, envios)
}
func (handler *EnvioHandler) InsertarEnvio(c *gin.Context) {
	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.envioService.InsertarEnvio(&envio)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *EnvioHandler) ModificarEnvio(c *gin.Context) {
	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	envio.ID = c.Param("id")
	resultado := handler.envioService.ModificarEnvio(&envio)
	c.JSON(http.StatusCreated, resultado)
}
