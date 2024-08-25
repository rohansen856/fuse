package controllers

import (
	"context"
	"github.com/ritankarsaha/backend/database"
	"github.com/ritankarsaha/backend/models"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var commentCollection *mongo.Collection = database.OpenCollection(database.Client,"comments")

func CreateComment(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.ID = primitive.NewObjectID()
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	_, err := commentCollection.InsertOne(ctx, comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func GetCommentsByNewsID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
	newsID := c.Param("newsID")
	oid, err := primitive.ObjectIDFromHex(newsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID"})
		return
	}

	filter := bson.M{"post_id": oid}

	cursor, err := commentCollection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	if err = cursor.All(ctx, &comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func DeleteComment(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
	commentID := c.Param("commentID")
	oid, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	filter := bson.M{"_id": oid}

	result, err := commentCollection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})

}
