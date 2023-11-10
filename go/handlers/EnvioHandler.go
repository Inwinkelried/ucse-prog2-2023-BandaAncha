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
	//Creamos el filtro
	filtro := dto.FiltroEnvio{
		PatenteCamion: patente,
		Estado:        estado,
		UltimaParada:  ultimaParada,
		FechaMenor:    fechaMenor,
		FechaMayor:    fechaMayor,
	}

	//Llama al service
	envios, err := handler.envioService.ObtenerEnviosFiltrados(filtro)

	//Si hay un error, lo devolvemos
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerEnvios][envio:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
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
	resultado, _ := handler.envioService.DespachadoEnvio(&envio)
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
