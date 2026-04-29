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

func CreateBuku(c *gin.Context) {
	var buku models.Buku
	// c.BindJSON(&buku)

	if err := c.ShouldBindJSON(&buku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Validate.Struct(buku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("buku")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, _ := collection.InsertOne(ctx, buku)

	c.JSON(http.StatusOK, result)
}

func GetAllBuku(c *gin.Context) {
	collection := config.DB.Collection("buku")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// cursor, _ := collection.Find(ctx, map[string]interface{}{})
	cursor, _ := collection.Find(ctx, bson.M{})

	var result []models.Buku
	cursor.All(ctx, &result)

	c.JSON(http.StatusOK, result)
}

func GetBukuByID(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	collection := config.DB.Collection("buku")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var buku models.Buku
	collection.FindOne(ctx, bson.M{"_id": id})

	c.JSON(http.StatusOK, buku)
}

func UpdateBuku(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var buku models.Buku

	collection := config.DB.Collection("buku")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": buku}

	collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteBuku(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	collection := config.DB.Collection("buku")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection.DeleteOne(ctx, bson.M{"_id": id})
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
