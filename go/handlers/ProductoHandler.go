package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/services"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
	ProductoService services.ProductoInterface
}

func NewProductoHandler(productoService services.ProductoInterface) *ProductoHandler {
	return &ProductoHandler{
		ProductoService: productoService,
	}
}
func (handler *ProductoHandler) ObtenerProductos(c *gin.Context) {
	productos := handler.ProductoService.ObtenerProductos()
	log.Printf("[handler:ProductoHandler][method:ObtenerProductos][productos:%v][cantidad:%d]", productos, len(productos))
	c.JSON(http.StatusOK, productos)
}
func (handler *ProductoHandler) ObtenerProductosFiltrados(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	filtrarPorStockMinimoStr := c.DefaultQuery("filtrarPorStockMinimo", "false")
	filtrarPorStockMinimo, err := strconv.ParseBool(filtrarPorStockMinimoStr)
	if err != nil {
		log.Printf("[handler:ProductoHandler][method:ObtenerProductos][error:%s][user:%s]", err.Error(), user.Codigo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tipoProducto := c.DefaultQuery("tipoProducto", "")
	filtroProducto := utils.FiltroProducto{
		FiltroStockMinimo: filtrarPorStockMinimo,
		TipoProducto:      tipoProducto,
	}
	productos, err := handler.ProductoService.ObtenerProductosFiltrados(filtroProducto)
	if err != nil {
		log.Printf("[handler:ProductoHandler][method:ObtenerProductos][error:%s][user:%s]", err.Error(), user.Codigo)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[handler:ProductoHandler][method:ObtenerProductos][cantidad:%d][user:%s]", len(productos), user.Codigo)
	c.JSON(http.StatusOK, productos)
}
func (handler *ProductoHandler) ObtenerProductoPorID(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	//invocamos al metodo
	producto, err := handler.ProductoService.ObtenerProductoPorID(&dto.Producto{ID: id})
	if err != nil {
		log.Printf("[handler:ProductoHandler][method:ObtenerProductoPorId][producto:%+v][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, producto)
}
func (handler *ProductoHandler) InsertarProducto(c *gin.Context) {
	var producto dto.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.ProductoService.InsertarProducto(&producto)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *ProductoHandler) ModificarProducto(c *gin.Context) {
	var producto dto.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	producto.ID = c.Param("id")
	resultado := handler.ProductoService.ModificarProducto(&producto)
	c.JSON(http.StatusCreated, resultado)
}
func (handler *ProductoHandler) EliminarProducto(c *gin.Context) {
	id := c.Param("id")
	productos := handler.ProductoService.EliminarProducto(id)
	c.JSON(http.StatusOK, productos)
}
