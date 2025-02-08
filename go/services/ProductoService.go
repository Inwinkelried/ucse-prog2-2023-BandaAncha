package services

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/dto"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/repositories"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/utils"
)

type ProductoInterface interface {
	ObtenerProductos() ([]*dto.Producto, error)
	InsertarProducto(Producto *dto.Producto) error
	EliminarProducto(id string) error
	ModificarProducto(Producto *dto.Producto) error
	ObtenerProductoPorID(productoConID *dto.Producto) (*dto.Producto, error)
	ObtenerProductosFiltrados(filtro dto.FiltroProducto) ([]dto.Producto, error)
}
type ProductoService struct {
	ProductoRepository repositories.ProductoRepositoryInterface
}

func NewProductoService(ProductoRepository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{
		ProductoRepository: ProductoRepository,
	}
}
func (service *ProductoService) ObtenerProductos() ([]*dto.Producto, error) {
	productosDb, err := service.ProductoRepository.ObtenerProductos()
	var Productos []*dto.Producto
	for _, productosDb := range productosDb {
		Producto := dto.NewProducto(productosDb)
		Productos = append(Productos, Producto)
	}
	return Productos, err
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

func (service *ProductoService) InsertarProducto(Producto *dto.Producto) error {
	_, err := service.ProductoRepository.InsertarProducto(Producto.GetModel())
	return err
}
func (service *ProductoService) ModificarProducto(Producto *dto.Producto) error {
	_, err := service.ProductoRepository.ModificarProducto(Producto.GetModel())
	return err
}
func (service *ProductoService) EliminarProducto(id string) error {
	_, err := service.ProductoRepository.EliminarProducto(utils.GetObjectIDFromStringID(id))
	return err
}
func (service *ProductoService) ObtenerProductosFiltrados(filtro dto.FiltroProducto) ([]dto.Producto, error) {
	productos, err := service.ProductoRepository.ObtenerProductoFiltrados(filtro)
	var producto *dto.Producto
	var productosDTO []dto.Producto
	if err != nil {
		return nil, err
	} else {
		for _, productoDB := range productos {
			producto = dto.NewProducto(productoDB)
			productosDTO = append(productosDTO, *producto)
		}
	}
	return productosDTO, nil
}
