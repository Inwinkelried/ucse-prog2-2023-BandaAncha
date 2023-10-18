package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/utils"
)

type ProductoInterface interface {
	ObtenerProductos() []*dto.Producto
	InsertarProducto(Producto *dto.Producto) bool
	EliminarProducto(id string) bool
	ModificarProducto(Producto *dto.Producto) bool
}
type ProductoService struct {
	ProductoRepository repositories.ProductoRepositoryInterface
}

func NewProductoService(ProductoRepository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{
		ProductoRepository: ProductoRepository,
	}
}
func (service *ProductoService) ObtenerProductos() []*dto.Producto {
	ProductosDB, _ := service.ProductoRepository.ObtenerProductos()
	var Productos []*dto.Producto
	for _, ProductosDB := range ProductosDB {
		Producto := dto.NewProducto(ProductosDB)
		Productos = append(Productos, Producto)
	}
	return Productos
}
func (service *ProductoService) InsertarProducto(Producto *dto.Producto) bool {
	service.ProductoRepository.InsertarProducto(Producto.GetModel())
	return true
}
func (service *ProductoService) ModificarProducto(Producto *dto.Producto) bool {
	service.ProductoRepository.ModificarProducto(Producto.GetModel())
	return true
}
func (service *ProductoService) EliminarProducto(id string) bool {
	service.ProductoRepository.EliminarProducto(utils.GetObjectIDFromStringID(id))
	return true
}
