package main

import (
	"os"
	"toko-buku-api/config"
	"toko-buku-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	routes.SetUpRoutes()

	port := os.Getenv("PORT")
	r.Run(port)
}
