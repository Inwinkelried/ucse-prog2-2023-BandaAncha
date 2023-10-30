package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/repositories"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type ProductoInterface interface {
	ObtenerProductos() []*dto.Producto
	InsertarProducto(Producto *dto.Producto) bool
	EliminarProducto(id string) bool
	ModificarProducto(Producto *dto.Producto) bool
	ObtenerProductoPorID(productoConID *dto.Producto) (*dto.Producto, error)
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
} //cambio realizado aca, revisar
func (service *ProductoService) ObtenerProductoPorID(productoConID *dto.Producto) (*dto.Producto, error) {
	productoDB, err := service.ProductoRepository.ObtenerProductoPorID(productoConID.GetModel())

	var producto *dto.Producto

	if err != nil {
		return nil, err
	} else {
		producto = dto.NewProducto(productoDB)
	}
	return producto, nil
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
