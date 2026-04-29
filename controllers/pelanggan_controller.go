package controllers

import (
	"context"
	"toko-buku-api/config"
	"toko-buku-api/models"
	"toko-buku-api/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePelanggan(c *gin.Context) {
	var data models.Pelanggan
	c.ShouldBindJSON(&data)

	if err := utils.Validate.Struct(data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	coll := config.DB.Collection("pelanggan")
	res, _ := coll.InsertOne(context.TODO(), data)

	c.JSON(200, res)
}

func GetPelanggan(c *gin.Context) {
	collection := config.DB.Collection("pelanggan")

	cur, _ := collection.Find(context.TODO(), bson.M{})

	var data []models.Pelanggan
	cur.All(context.TODO(), &data)

	c.JSON(200, data)
}

func UpdatePelanggan(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var data models.Pelanggan
	c.BindJSON(&data)

	col := config.DB.Collection("pelanggan")

	update := bson.M{"$set": data}

	col.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	c.JSON(200, gin.H{"Message": "updated pelanggan"})
}

func DeletePelanggan(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	col := config.DB.Collection("pelanggan")

	col.DeleteOne(context.TODO(), bson.M{"_id": id})
	c.JSON(200, gin.H{"message": "deleted pelanggan"})
}
