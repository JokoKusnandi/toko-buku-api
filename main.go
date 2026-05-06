package main

import (
    "os"
    "toko-buku-api/config"
    "toko-buku-api/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Connect ke MongoDB
    config.ConnectDB()

    // Setup routes dengan engine Gin
    routes.SetUpRoutes(r)

    // Jalankan server
    port := os.Getenv("PORT")
    r.Run(":" + port) // tambahkan ":" agar port dikenali
}
