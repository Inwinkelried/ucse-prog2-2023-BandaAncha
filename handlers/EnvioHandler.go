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

// FALTA PROBAR
func (handler *EnvioHandler) AgregarParada(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	var parada dto.Parada
	if err := c.ShouldBindJSON(&parada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	envio := dto.Envio{
		ID: id,
		Paradas: []dto.Parada{
			parada,
		},
	}
	operacion, err := handler.envioService.AgregarParada(&envio)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:AgregarParada][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !operacion {
		log.Printf("[handler:EnvioHandler][method:AgregarParada][envio:%+v][user:%s]", err.Error(), user.Codigo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //es correcto devolver bad request aca?
		return
	}

	log.Printf("[handler:EnvioHandler][method:AgregarParada][envio:%+v][user:%s]", envio, user.Codigo)

	c.JSON(http.StatusOK, envio)
}

// FALTA PROBAR
func (handler *EnvioHandler) ObtenerEnvioPorID(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	//invocamos al metodo
	envio, err := handler.envioService.ObtenerEnvioPorID(&dto.Envio{ID: id})
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerEnvioPorId][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Agregamos un log para indicar informacion
	c.JSON(http.StatusOK, envio)
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
func (handler *EnvioHandler) DespachadoEnvio(c *gin.Context) {
	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	envio.ID = c.Param("id")
	resultado := handler.envioService.DespachadoEnvio(&envio)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *EnvioHandler) EnRutaEnvio(c *gin.Context) {
	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	envio.ID = c.Param("id")
	resultado := handler.envioService.EnRutaEnvio(&envio)
	c.JSON(http.StatusCreated, resultado)
}
