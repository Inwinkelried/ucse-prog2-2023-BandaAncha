package main

import (
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/handlers"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/repositories"
	"github.com/Inwinkelried/ucse-prog2-2023-BandaAncha/services"
	"github.com/gin-gonic/gin"
)

var (
	camionHandler *handlers.CamionHandler
	router        *gin.Engine
)

func dependencies() {
	var database repositories.DB
	var camionRepo repositories.CamionRepositoryInterface
	var camionService services.CamionInterface

	database = repositories.NewMongoDB()
	camionRepo = repositories.NewCamionRepository(database)
	camionService = services.NewCamionService(camionRepo)
	camionHandler = handlers.NewCamionHandler(camionService)
}
