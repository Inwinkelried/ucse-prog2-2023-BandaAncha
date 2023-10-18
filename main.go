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
	//Uso del middleware para todas las rutas del grupo
	//group.Use(authMiddleware.ValidateToken)
	//group.Use(middlewares.CORSMiddleware())

	//PRODUCTOS
	groupProducto.GET("/", camionHandler.ObtenerCamiones)
	//group.GET("/:id", aulaHandler.ObtenerAulaPorID)
	groupProducto.POST("/", camionHandler.InsertarCamion)
	groupProducto.PUT("/:id", camionHandler.ModificarCamion)
	groupProducto.DELETE("/:id", camionHandler.EliminarCamion)

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
}
