package handlers

import (
	"log"
	"net/http"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/services"
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
