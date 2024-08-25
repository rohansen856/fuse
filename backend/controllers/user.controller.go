package controllers

import (
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/ritankarsaha/backend/database"
	helper "github.com/ritankarsaha/backend/helpers"
	"github.com/ritankarsaha/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")

		if err := helper.MathUserTypeToUid(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		var user models.User
		objID, _ := primitive.ObjectIDFromHex(userId)
		err := userCollection.FindOne(c, bson.M{"_id": objID}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "User not found",
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}


func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")

		if err := helper.MathUserTypeToUid(ctx, userId); 
		err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var updateData bson.M
		if err := ctx.BindJSON(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input",
			})
			return
		}

		userRole, _ := ctx.Get("user_role") 
		if _, exists := updateData["address"]; exists && userRole != "MANAGER" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "User has been unauthorized to change/modify.",
			})
			return
		}

		delete(updateData, "_id")
		delete(updateData, "email") 
		delete(updateData, "role")  

		if len(updateData) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Please enter a valid field to update the user data",
			})
			return
		}

		updateData["updated_at"] = time.Now()

		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		objID, _ := primitive.ObjectIDFromHex(userId)
		filter := bson.M{"_id": objID}

		update := bson.M{
			"$set": updateData,
		}

		result, err := userCollection.UpdateOne(c, filter, update)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if result.MatchedCount == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "User not found or doesn't exist.",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "User data has been updated successfully!",
		})
	}
}
