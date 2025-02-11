package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/services"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils/logging"
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
	patente := c.DefaultQuery("patente", "")
	ultimaParada := c.DefaultQuery("ultimaParada", "")
	estado := c.DefaultQuery("estado", "")
	fechaMenor := parseFechaQuery(c.DefaultQuery("fechaMenor", "0001-01-01T00:00:00Z"))
	fechaMayor := parseFechaQuery(c.DefaultQuery("fechaMayor", "0001-01-01T00:00:00Z"))
	filtro := dto.FiltroEnvio{
		PatenteCamion: patente,
		Estado:        estado,
		UltimaParada:  ultimaParada,
		FechaMenor:    fechaMenor,
		FechaMayor:    fechaMayor,
	}
	envios, err := handler.envioService.ObtenerEnvios(filtro)
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerEnvios][error:%s][user:%s]", err.Error(), user.Codigo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[handler:EnvioHandler][method:ObtenerEnvios][cantidad:%d][user:%s]", len(envios), user.Codigo)
	c.JSON(http.StatusOK, envios)
}
func parseFechaQuery(fechaStr string) time.Time {
	fecha, err := time.Parse(time.RFC3339, fechaStr)
	if err != nil {
		log.Printf("[error:ParseFechaQuery] No se pudo parsear la fecha: %s - Usando valor por defecto", fechaStr)
		return time.Time{}
	}
	return fecha
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
		log.Printf("[handler:EnvioHandler][method:InsertarEnvio][envio:%+v]", envio)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se ha podido insertar el envío. Revise los datos de entrada"})
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

// REPORTES
func (handler *EnvioHandler) ObtenerCantidadEnviosPorEstado(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	//Obtenemos el array de cantidades del service
	cantidades, err := handler.envioService.ObtenerCantidadEnviosPorEstado()
	//Si hay un error, lo devolvemos
	if err != nil {
		logging.LoggearErrorYResponder(c, "EnvioHandler", "ObtenerCantidadEnviosPorEstado", err, &user)
		return
	}
	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "EnvioHandler", "ObtenerCantidadEnviosPorEstado", cantidades, &user)
}
