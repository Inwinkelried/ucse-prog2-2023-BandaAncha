package main

import (
	"log"
	"time"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/handlers"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/middlewares"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/repositories"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/services"
	"github.com/gin-gonic/gin"
)

var (
	camionHandler   *handlers.CamionHandler
	productoHandler *handlers.ProductoHandler
	pedidoHandler   *handlers.PedidoHandler
	envioHandler    *handlers.EnvioHandler
	router          *gin.Engine
)

func main() {
	router = gin.Default()

	// Apply CORS middleware first
	router.Use(middlewares.CORSMiddleware())

	// Handle trailing slashes
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	// Add this temporary debug endpoint
	router.GET("/debug", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "API is running",
			"time":   time.Now(),
		})
	})

	//Iniciar objetos de handler
	dependencies()
	log.Println("Cargando dependencias...")
	//Iniciar rutas
	mappingRoutes()
	log.Println("Mapeando rutas...")

	log.Println("Iniciando el servidor...")
	router.Run(":8080")
	log.Print("Servidor iniciado en el puerto 8080")
}
func mappingRoutes() {

	//Listado de rutas
	groupCamion := router.Group("/trucks")
	groupProducto := router.Group("/products")
	groupPedido := router.Group("/orders")
	groupEnvio := router.Group("/shippings")

	// groupCamion.Use(middlewares.CORSMiddleware())
	// groupProducto.Use(middlewares.CORSMiddleware())
	// groupPedido.Use(middlewares.CORSMiddleware())
	// groupEnvio.Use(middlewares.CORSMiddleware())
	//Uso del middleware para todas las rutas del grupo
	// router.Use(authMiddleware.ValidateToken)

	//ENVIO
	groupEnvio.GET("/", envioHandler.ObtenerEnvios)
	groupEnvio.POST("/", envioHandler.InsertarEnvio)
	groupEnvio.PUT("/AddStop/:id", envioHandler.AgregarParada)
	groupEnvio.PUT("/SetDelivered/:id", envioHandler.DespachadoEnvio)
	groupEnvio.GET("/:id", envioHandler.ObtenerEnvioPorID)
	groupEnvio.PUT("/SetSent/:id", envioHandler.EnRutaEnvio)

	//REPORTES ENVIOS
	router.GET("/envios/cantidadPorEstado", envioHandler.ObtenerCantidadEnviosPorEstado)

	//PEDIDOS
	groupPedido.GET("/", pedidoHandler.ObtenerPedidos)
	groupPedido.GET("/Filter", pedidoHandler.ObtenerPedidosFiltrados)
	groupPedido.POST("/", pedidoHandler.InsertarPedido)
	groupPedido.GET("/Approved", pedidoHandler.ObtenerPedidosAprobados)

	groupPedido.PUT("/Confirm/:id", pedidoHandler.AceptarPedido) //AHORA ESTEE

	groupPedido.PUT("/Cancel/:id", pedidoHandler.CancelarPedido)
	groupPedido.PUT("/Send/:id", pedidoHandler.ParaEnviarPedido)
	groupPedido.PUT("/Sent/:id", pedidoHandler.EnviadoPedido)
	groupPedido.GET("/:id", pedidoHandler.ObtenerPedidoPorID)
	//REPOORTES PEDIDOS
	groupPedido.GET("/cantidadPorEstado", pedidoHandler.ObtenerCantidadPedidosPorEstado)
	//PRODUCTOS
	groupProducto.GET("/", productoHandler.ObtenerProductos)
	groupProducto.GET("/Filter", productoHandler.ObtenerProductosFiltrados)
	groupProducto.POST("/", productoHandler.InsertarProducto)
	groupProducto.GET("/:id", productoHandler.ObtenerProductoPorID)
	groupProducto.PUT("/:id", productoHandler.ModificarProducto)
	groupProducto.DELETE("/:id", productoHandler.EliminarProducto)
	//CAMIONES
	groupCamion.GET("/", camionHandler.ObtenerCamiones)
	groupCamion.GET("/:id", camionHandler.ObtenerCamionPorID)
	groupCamion.POST("/", camionHandler.InsertarCamion)
	groupCamion.PUT("/:id", camionHandler.ModificarCamion)
	groupCamion.DELETE("/:id", camionHandler.EliminarCamion)

}
func dependencies() {
	var database repositories.DB
	database = repositories.NewMongoDB()
	//CAMIONES
	var camionRepo repositories.CamionRepositoryInterface
	var camionService services.CamionInterface
	camionRepo = repositories.NewCamionRepository(database)
	camionService = services.NewCamionService(camionRepo)
	camionHandler = handlers.NewCamionHandler(camionService)
	//PRODUCTOS
	var productoRepo repositories.ProductoRepositoryInterface
	var productoService services.ProductoInterface
	productoRepo = repositories.NewProductoRepository(database)
	productoService = services.NewProductoService(productoRepo)
	productoHandler = handlers.NewProductoHandler(productoService)
	//PEDIDOS
	var pedidoRepo repositories.PedidoRepositoryInterface
	var pedidoService services.PedidoServiceInterface
	pedidoRepo = repositories.NewPedidoRepository(database)
	pedidoService = services.NewPedidoService(nil, camionRepo, pedidoRepo, productoRepo)
	pedidoHandler = handlers.NewPedidoHandler(pedidoService)
	//ENVIO
	var envioRepo repositories.EnvioRepositoryInterface
	var envioService services.EnvioServiceInterface
	envioRepo = repositories.NewEnvioRepository(database)
	envioService = services.NewEnvioService(envioRepo, camionRepo, pedidoRepo, productoRepo)
	envioHandler = handlers.NewEnvioHandler(envioService)

}
