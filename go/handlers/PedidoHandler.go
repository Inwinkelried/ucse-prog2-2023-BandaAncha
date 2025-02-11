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

type PedidoHandler struct {
	pedidoService services.PedidoServiceInterface
}

func NewPedidoHandler(pedidoService services.PedidoServiceInterface) *PedidoHandler {
	return &PedidoHandler{
		pedidoService: pedidoService,
	}
}
func (handler *PedidoHandler) ObtenerPedidosFiltrados(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	estado := c.DefaultQuery("estado", "")
	fechaMenorStr := c.DefaultQuery("fecha_menor", "0001-01-01T00:00:00Z")
	fechaMenor, err := time.Parse(time.RFC3339, fechaMenorStr)
	if err != nil {
		fechaMenor = time.Time{}
	}
	fechaMayorStr := c.DefaultQuery("fecha_mayor", "0001-01-01T00:00:00Z")
	fechaMayor, err := time.Parse(time.RFC3339, fechaMayorStr)
	if err != nil {
		fechaMayor = time.Time{}
	}
	filtro := dto.FiltroPedido{
		Estado:     estado,
		FechaMayor: fechaMayor,
		FechaMenor: fechaMenor,
	}
	pedidos, err := handler.pedidoService.ObtenerPedidosFiltrados(filtro)
	if err != nil {
		log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][error:%s][user:%s]", err.Error(), user.Codigo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][cantidad:%d][user:%s]", len(pedidos), user.Codigo)
	c.JSON(http.StatusOK, pedidos)
}
func (handler *PedidoHandler) ObtenerPedidosAprobados(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	pedidos, err := handler.pedidoService.ObtenerPedidosAprobados()
	if err != nil {
		log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][error:%s][user:%s]", err.Error(), user.Codigo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:PedidoHandler][method:ObtenerPedidosAprobados][cantidad:%d][user:%s]", len(pedidos), user.Codigo)
	c.JSON(http.StatusOK, pedidos)
}
func (handler *PedidoHandler) ObtenerPedidoPorID(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	pedido, err := handler.pedidoService.ObtenerPedidoPorID(&dto.Pedido{ID: id})
	if err != nil {
		log.Printf("[handler:EnvioHandler][method:ObtenerEnvioPorId][pedido:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pedido)
}
func (handler *PedidoHandler) ObtenerPedidos(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	pedidos, err := handler.pedidoService.ObtenerPedidos()
	if err != nil {
		log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][error:%s][user:%s]", err.Error(), user.Codigo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][cantidad:%d][user:%s]", len(pedidos), user.Codigo)
	c.JSON(http.StatusOK, pedidos)
}
func (handler *PedidoHandler) InsertarPedido(c *gin.Context) {
	var pedido dto.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado, err := handler.pedidoService.InsertarPedido(&pedido)
	if err != nil {
		log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v][error:%s]", pedido, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !resultado {
		log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v][error:%s]", pedido, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v]", pedido)
	c.JSON(http.StatusCreated, gin.H{"status": "Creado correctamente"})
}
func (handler *PedidoHandler) AceptarPedido(c *gin.Context) {
	var pedido dto.Pedido
	log.Printf("[handler:PedidoHandler][method:AceptarPedido] Recibiendo solicitud para aceptar pedido")
	// Agregar el ID del pedido desde los parámetros de la URL
	pedido.ID = c.Param("id")
	log.Printf("[handler:PedidoHandler][method:AceptarPedido][pedido:%+v] Intentando aceptar pedido", pedido)
	// Llamar al servicio para aceptar el pedido
	resultado, err := handler.pedidoService.AceptarPedido(&pedido)
	if err != nil {
		log.Printf("[handler:PedidoHandler][method:AceptarPedido][pedido:%+v][error:Error en la lógica de negocio][detalles:%s]", pedido, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !resultado {
		log.Printf("[handler:PedidoHandler][method:AceptarPedido][pedido:%+v][error:El servicio rechazó la operación]", pedido)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo aceptar el pedido"})
		return
	}
	log.Printf("[handler:PedidoHandler][method:AceptarPedido][pedido:%+v] Pedido aceptado con éxito", pedido)
	c.JSON(http.StatusCreated, gin.H{"status": "Pedido aceptado"})
}

func (handler *PedidoHandler) CancelarPedido(c *gin.Context) {
	var pedido dto.Pedido
	pedido.ID = c.Param("id")
	resultado, err := handler.pedidoService.CancelarPedido(&pedido)
	if err != nil {
		log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v][error:%s]", pedido, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !resultado {
		log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v][error:%s]", pedido, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v]", pedido)
	c.JSON(http.StatusCreated, gin.H{"status": "Pedido cancelado"})
}
func (handler *PedidoHandler) EnviadoPedido(c *gin.Context) {
	var pedido dto.Pedido
	pedido.ID = c.Param("id")
	resultado, err := handler.pedidoService.EnviadoPedido(&pedido)
	if err != nil {
		log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v][error:%s]", pedido, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !resultado {
		log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v][error:%s]", pedido, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v]", pedido)
	c.JSON(http.StatusCreated, gin.H{"status": "Pedido enviado"})
}
func (handler *PedidoHandler) ParaEnviarPedido(c *gin.Context) {
	var pedido dto.Pedido
	pedido.ID = c.Param("id")
	resultado, err := handler.pedidoService.ParaEnviarPedido(&pedido)
	if err != nil {
		log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v][error:%s]", pedido, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !resultado {
		log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v][error:%s]", pedido, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:PedidoHandler][method:InsertarPedido][pedido:%+v]", pedido)
	c.JSON(http.StatusCreated, gin.H{"status": "Pedido Para Enviar"})
}

func (handler *PedidoHandler) ObtenerCantidadPedidosPorEstado(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	cantidades, err := handler.pedidoService.ObtenerCantidadPedidosPorEstado()
	if err != nil {
		logging.LoggearErrorYResponder(c, "PedidoHandler", "ObtenerCantidadPedidosPorEstado", err, &user)
		return
	}

	logging.LoggearResultadoYResponder(c, "PedidoHandler", "ObtenerCantidadPedidosPorEstado", cantidades, &user)
}
