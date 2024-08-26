package controllers

import (
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/ritankarsaha/backend/database"
	"github.com/ritankarsaha/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var sessionCollection *mongo.Collection = database.OpenCollection(database.Client, "session")


//creating the session
func CreateSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var session models.Session
		if err := ctx.ShouldBindJSON(&session); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input data",
			})
			return
		}
		session.ID = primitive.NewObjectID()
		session.Expires = time.Now().Add(24 * time.Hour)

		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_, err := sessionCollection.InsertOne(c, session)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create session",
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Session created successfully",
			"session": session,
		})
	}
}




//getting the session
func GetSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId := ctx.Param("session_id")

		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		var session models.Session
		objID, _ := primitive.ObjectIDFromHex(sessionId)
		err := sessionCollection.FindOne(c, bson.M{"_id": objID}).Decode(&session)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "Session not found",
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"session": session,
		})
	}
}

//deleting the session
func DeleteSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId := ctx.Param("session_id")

		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		objID, _ := primitive.ObjectIDFromHex(sessionId)
		filter := bson.M{"_id": objID}

		result, err := sessionCollection.DeleteOne(c, filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if result.DeletedCount == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Session not found or doesn't exist.",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Session deleted successfully!",
		})
	}
}

