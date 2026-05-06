package routes

import (
	"toko-buku-api/controllers"
	"toko-buku-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	auth := r.Group("/")
	auth.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddlewares())

	//book
	protected.POST("/buku", controllers.CreateBuku)
	protected.GET("/buku", controllers.GetAllBuku)
	protected.GET("/buku/:id", controllers.GetBukuByID)
	protected.PUT("/buku/:id", controllers.UpdateBuku)
	protected.DELETE("/buku/:id", controllers.DeleteBuku)

	//penerbit
	protected.POST("/penerbit", controllers.CreatePenerbit)
	protected.GET("/penerbit", controllers.GetPenerbit)
	protected.PUT("/penerbit/:id", controllers.UpdatePenerbit)
	protected.DELETE("/penerbit/:id", controllers.DeletePenerbit)

	//pelanggan
	protected.POST("/pelanggan", controllers.CreatePelanggan)
	protected.GET("/pelanggan", controllers.GetPelanggan)
	protected.PUT("/pelanggan/:id", controllers.UpdatePelanggan)
	protected.DELETE("/pelanggan/:id", controllers.DeletePelanggan)

	//transaksi
	protected.POST("/transaksi", controllers.CreateTransaksi)
	protected.GET("/transaksi", controllers.GetTransaksi)

	//laporan
	protected.GET("/laporan/total", controllers.TotalPendapatan)
	protected.GET("/laporan/terlaris", controllers.BukuTerlaris)

	// r.GET("/Getbuku", controllers.GetBuku())
	// r.POST("/Createbuku", controllers.CreateBuku())
}
