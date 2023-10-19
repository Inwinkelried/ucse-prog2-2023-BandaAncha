package main

import (
	"log"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/handlers"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/services"
	"github.com/gin-gonic/gin"
)

var (
	camionHandler   *handlers.CamionHandler
	productoHandler *handlers.ProductoHandler
	pedidoHandler   *handlers.PedidoHandler
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
	//Uso del middleware para todas las rutas del grupo
	//group.Use(authMiddleware.ValidateToken)
	//group.Use(middlewares.CORSMiddleware())
	//PEDIDOS
	groupPedido.GET("/", pedidoHandler.ObtenerPedidos)
	//group.GET("/:id", aulaHandler.ObtenerAulaPorID)
	groupPedido.POST("/", pedidoHandler.InsertarPedido)
	groupPedido.PUT("/:id", pedidoHandler.ModificarPedido)
	groupPedido.DELETE("/:id", pedidoHandler.EliminarPedido)
	//PRODUCTOS
	groupProducto.GET("/", productoHandler.ObtenerProductos)
	//group.GET("/:id", aulaHandler.ObtenerAulaPorID)
	groupProducto.POST("/", productoHandler.InsertarProducto)
	groupProducto.PUT("/:id", productoHandler.ModificarProducto)
	groupProducto.DELETE("/:id", productoHandler.EliminarProducto)

	//CAMIONES
	groupCamion.GET("/", camionHandler.ObtenerCamiones)
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
	pedidoService = services.NewPedidoService(pedidoRepo)
	pedidoHandler = handlers.NewPedidoHandler(pedidoService)

}
