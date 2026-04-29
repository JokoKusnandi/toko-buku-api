package controllers

import (
	"context"
	"toko-buku-api/config"
	"toko-buku-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTransaksi(c *gin.Context) {
	var trx models.Transaksi

	c.ShouldBindJSON(&trx)

	col := config.DB.Collection("transaksi")
	cur, _ := col.Find(context.TODO(), bson.M{})

	var data []models.Transaksi
	cur.All(context.TODO(), &data)

	c.JSON(200, data)
}

func GetTransaksi(c *gin.Context) {
	collection := config.DB.Collection("transaksi")

	cur, _ := collection.Find(context.TODO(), bson.M{})

	var data []models.Transaksi
	cur.All(context.TODO(), &data)

	c.JSON(200, data)
}
