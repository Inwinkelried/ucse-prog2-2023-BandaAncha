package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/services"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
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
func (handler *EnvioHandler) ObtenerEnviosFiltrados(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	patente := c.DefaultQuery("patente", "")
	ultimaParada := c.DefaultQuery("ultimaParada", "")
	estado := c.DefaultQuery("estado", "")
	fechaMenorStr := c.DefaultQuery("fechaMenor", "0001-01-01T00:00:00Z")
	fechaMenor, err := time.Parse(time.RFC3339, fechaMenorStr)
	if err != nil {
		fechaMenor = time.Time{}
	}
	fechaMayorStr := c.DefaultQuery("fechaMayor", "0001-01-01T00:00:00Z")
	fechaMayor, err := time.Parse(time.RFC3339, fechaMayorStr)
	if err != nil {
		fechaMayor = time.Time{}
	}
	filtro := dto.FiltroEnvio{
		PatenteCamion: patente,
		Estado:        estado,
		UltimaParada:  ultimaParada,
		FechaMenor:    fechaMenor,
		FechaMayor:    fechaMayor,
	}
	envios, err := handler.envioService.ObtenerEnviosFiltrados(filtro)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerEnvios][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:AulaHandler][method:ObtenerEnvios][cantidad:%d][user:%s]", len(envios), user.Codigo)
	c.JSON(http.StatusOK, envios)
}
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
	log.Printf("[handler:EnvioHandler][method:ObtenerEnvioPorId][envio:%+v][user:%s]", envio, user.Codigo)
	c.JSON(http.StatusOK, envio)
}

func (handler *EnvioHandler) ObtenerEnvios(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	envios, err := handler.envioService.ObtenerEnvios()
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerEnvios][error:%s][user:%s]", err.Error(), user.Codigo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[handler:EnvioHandler][method:ObtenerEnvios][cantidad:%d][user:%s]", len(envios), user.Codigo)
	c.JSON(http.StatusOK, envios)
}
func (handler *EnvioHandler) InsertarEnvio(c *gin.Context) {
	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado, err := handler.envioService.InsertarEnvio(&envio)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:InsertarEnvio][envio:%+v][error:%s]", envio, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !resultado {
		log.Printf("[handler:EnvioHandler][method:InsertarEnvio][envio:%+v][error:%s]", envio, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"No se ha podido insertar el envio. Revise los datos de entrada": err.Error()})
		return
	}
	log.Printf("[handler:EnvioHandler][method:InsertarEnvio][envio:%+v]", envio)
	c.JSON(http.StatusCreated, gin.H{"status": "Creado correctamente"})

}
func (handler *EnvioHandler) DespachadoEnvio(c *gin.Context) {
	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	envio.ID = c.Param("id")
	resultado, err := handler.envioService.DespachadoEnvio(&envio)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:EnRutaEnvio][envio:%+v][error:%s]", envio, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !resultado {
		log.Printf("[handler:EnvioHandler][method:EnRutaEnvio][envio:%+v][error:%s]", envio, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"ERROR: No se ha podido actualizar el envio": err.Error()})
		return
	}
	log.Printf("[handler:EnvioHandler][method:EnRutaEnvio][envio:%+v]", envio)
	c.JSON(http.StatusCreated, gin.H{"status": "Actualizado correctamente"})

}
func (handler *EnvioHandler) EnRutaEnvio(c *gin.Context) {
	var envio dto.Envio
	if err := c.ShouldBindJSON(&envio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	envio.ID = c.Param("id")
	resultado, err := handler.envioService.EnRutaEnvio(&envio)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:EnRutaEnvio][envio:%+v][error:%s]", envio, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !resultado {
		log.Printf("[handler:EnvioHandler][method:EnRutaEnvio][envio:%+v][error:%s]", envio, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"ERROR: No se ha podido actualizar el envio": err.Error()})
		return
	}
	log.Printf("[handler:EnvioHandler][method:EnRutaEnvio][envio:%+v]", envio)
	c.JSON(http.StatusCreated, gin.H{"status": "Actualizado correctamente"})
}
