package main

import (
	"log"

	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/handlers"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/services"
	"github.com/gin-gonic/gin"
)

var (
	camionHandler *handlers.CamionHandler
	router        *gin.Engine
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
	//Uso del middleware para todas las rutas del grupo
	//group.Use(authMiddleware.ValidateToken)

	//group.Use(middlewares.CORSMiddleware())

	groupCamion.GET("/", camionHandler.ObtenerCamiones)
	//group.GET("/:id", aulaHandler.ObtenerAulaPorID)
	groupCamion.POST("/", camionHandler.InsertarCamion)
	groupCamion.PUT("/:id", camionHandler.ModificarCamion)
	groupCamion.DELETE("/:id", camionHandler.EliminarCamion)

}
func dependencies() {
	var database repositories.DB
	var camionRepo repositories.CamionRepositoryInterface
	var camionService services.CamionInterface

	database = repositories.NewMongoDB()
	camionRepo = repositories.NewCamionRepository(database)
	camionService = services.NewCamionService(camionRepo)
	camionHandler = handlers.NewCamionHandler(camionService)
}
