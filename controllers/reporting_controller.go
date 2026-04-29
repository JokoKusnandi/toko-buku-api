package controllers

import (
	"context"
	"toko-buku-api/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// total revenue
func TotalPendapatan(c *gin.Context) {
	col := config.DB.Collection("transaksi")

	pipeline := []bson.M{
		{"$unwind": "$details"},
		{"$group": bson.M{
			"_id":   nil,
			"total": bson.M{"$sum": bson.M{"multiply": []interface{}{"$details.jumlah", "details.harga"}}},
		}},
	}

	cur, _ := col.Aggregate(context.TODO(), pipeline)

	var result []bson.M
	cur.All(context.TODO(), &result)

	c.JSON(200, result)
}

func BukuTerlaris(c *gin.Context) {
	col := config.DB.Collection("transaksi")

	pipeline := []bson.M{
		{"$unwind": "$details"},
		{"$group": bson.M{
			"_id":           "$details.buku_id",
			"total_terjual": bson.M{"$sum": "$details.jumlah"},
		}},
		{"$sort": bson.M{"total_terjual": -1}},
	}

	cur, _ := col.Aggregate(context.TODO(), pipeline)

	var result []bson.M
	cur.All(context.TODO(), &result)

	c.JSON(200, result)
}
