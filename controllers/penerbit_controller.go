package controllers

import (
	"context"
	"net/http"
	"time"
	"toko-buku-api/config"
	"toko-buku-api/models"
	"toko-buku-api/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePenerbit(c *gin.Context) {
	var data models.Penerbit

	if err := utils.Validate.Struct(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("penerbit")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, _ := collection.InsertOne(ctx, data)

	c.JSON(http.StatusOK, result)

}

func GetPenerbit(c *gin.Context) {
	collection := config.DB.Collection("penerbit")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, _ := collection.Find(ctx, bson.M{})

	var data []models.Penerbit
	cursor.All(ctx, &data)

	c.JSON(http.StatusOK, data)
}

func UpdatePenerbit(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var data models.Penerbit
	c.BindJSON(&data)

	collection := config.DB.Collection("penerbit")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": data}

	collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	c.JSON(http.StatusOK, gin.H{"Message": "updated penerbit"})
}

func DeletePenerbit(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	collection := config.DB.Collection("penerbit")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection.DeleteOne(ctx, bson.M{"_id": id})
	c.JSON(http.StatusOK, gin.H{"message": "deleted penerbit"})
}
