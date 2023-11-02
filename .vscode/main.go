package main

import (
	"log"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/go/handlers"
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
	//Iniciar objetos de handler
	dependencies()
	//Iniciar rutas
	mappingRoutes()

	log.Println("Iniciando el servidor...")
	router.Run(":8080")
}
func mappingRoutes() {
	//middleware para permitir peticiones del mismo server localhost

	//cliente para api externa
	//var authClient clients.AuthClientInterface
	//authClient = clients.NewAuthClient()
	//creacion de middleware de autenticacion
	//authMiddleware := middlewares.NewAuthMiddleware(authClient)

	//Listado de rutas
	groupCamion := router.Group("/trucks")
	groupProducto := router.Group("/products")
	groupPedido := router.Group("/orders")
	groupEnvio := router.Group("/shippings")
	//Uso del middleware para todas las rutas del grupo
	//group.Use(authMiddleware.ValidateToken)
	//group.Use(middlewares.CORSMiddleware())

	//ENVIO
	groupEnvio.GET("/", envioHandler.ObtenerEnvios)
	//hay que probar
	groupEnvio.POST("/", envioHandler.InsertarEnvio)
	// hay q probar
	groupEnvio.PUT("/AddStop/:id", envioHandler.AgregarParada)
	//hay que probar
	groupEnvio.PUT("/SetDelivered/:id", envioHandler.DespachadoEnvio)
	groupEnvio.GET("/:id", envioHandler.ObtenerEnvioPorID)
	groupEnvio.PUT("/SetSent/:id", envioHandler.EnRutaEnvio)

	//PEDIDOS
	groupPedido.GET("/", pedidoHandler.ObtenerPedidos)
	//group.GET("/:id", aulaHandler.ObtenerAulaPorID)
	groupPedido.POST("/", pedidoHandler.InsertarPedido)
	groupPedido.GET("/Approved", pedidoHandler.ObtenerPedidosAprobados)
	groupPedido.PUT("/Confirm/:id", pedidoHandler.AceptarPedido)
	groupPedido.PUT("/Cancel/:id", pedidoHandler.CancelarPedido)
	groupPedido.PUT("/Send/:id", pedidoHandler.ParaEnviarPedido)
	groupPedido.PUT("/Sent/:id", pedidoHandler.EnviadoPedido)
	groupPedido.GET("/:id", pedidoHandler.ObtenerPedidoPorID)

	//PRODUCTOS
	groupProducto.GET("/", productoHandler.ObtenerProductos)
	//group.GET("/:id", aulaHandler.ObtenerAulaPorID)
	groupProducto.POST("/", productoHandler.InsertarProducto)
	groupProducto.GET("/:id", productoHandler.ObtenerProductoPorID)
	groupProducto.PUT("/:id", productoHandler.ModificarProducto)
	groupProducto.DELETE("/:id", productoHandler.EliminarProducto)

	//CAMIONES
	groupCamion.GET("/", camionHandler.ObtenerCamiones)
	groupCamion.GET("/:id", camionHandler.ObtenerCamionPorID)
	//group.GET("/:id", aulaHandler.ObtenerAulaPorID)
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
