package main

import (
	"authentication/internal/entrypoints/handler/rest"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Booting application...")

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	log.Info("configuring application port " + port)

	if err := run(port); err != nil {
		log.Errorf("error running server", err)
	}
}

func run(port string) error {
	log.Info("loading routes")
	router := gin.Default()

	pingHandler := rest.PingHandlerFactory()
	pingHandler.RegisterRoutes(router)

	authHandler := rest.AuthHandlerFactory()
	authHandler.RegisterRouters(router)

	log.Info("run application...")
	return router.Run(":" + port)
}
