package main

import (
	"go-study/adapter/input/routes"
	"go-study/configuration/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger.Info("Iniciando aplicação")
	router := gin.Default()

	routes.InitRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Erro ao iniciar aplicação na porta 8080", err)
		return
	}
}
