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
	filtro := utils.FiltroPedido{
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
	pedidos := handler.pedidoService.ObtenerPedidosAprobados()
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
	//Agregamos un log para indicar informacion
	c.JSON(http.StatusOK, pedido)
}
func (handler *PedidoHandler) ObtenerPedidos(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	pedidos := handler.pedidoService.ObtenerPedidos()
	log.Printf("[handler:PedidoHandler][method:ObtenerPedidos][cantidad:%d][user:%s]", len(pedidos), user.Codigo)
	c.JSON(http.StatusOK, pedidos)
}
func (handler *PedidoHandler) InsertarPedido(c *gin.Context) {
	var pedido dto.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.pedidoService.InsertarPedido(&pedido)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *PedidoHandler) AceptarPedido(c *gin.Context) {
	var pedido dto.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pedido.ID = c.Param("id")
	resultado := handler.pedidoService.AceptarPedido(&pedido)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *PedidoHandler) CancelarPedido(c *gin.Context) {
	var pedido dto.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pedido.ID = c.Param("id")
	resultado := handler.pedidoService.CancelarPedido(&pedido)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *PedidoHandler) EnviadoPedido(c *gin.Context) {
	var pedido dto.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pedido.ID = c.Param("id")
	resultado := handler.pedidoService.EnviadoPedido(&pedido)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *PedidoHandler) ParaEnviarPedido(c *gin.Context) {
	var pedido dto.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pedido.ID = c.Param("id")
	resultado := handler.pedidoService.ParaEnviarPedido(&pedido)
	c.JSON(http.StatusCreated, resultado)
}
