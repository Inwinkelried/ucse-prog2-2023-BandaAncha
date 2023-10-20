package handlers

import (
	"log"
	"net/http"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/services"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
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
