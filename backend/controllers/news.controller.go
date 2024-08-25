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

var newsCollection *mongo.Collection = database.OpenCollection(database.Client, "news")

func CreateNews(c *gin.Context) {
	var input models.News

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	result, err := newsCollection.InsertOne(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News created successfully", "news_id": result.InsertedID})
}


func GetNews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var newsList []models.News

		cursor, err := newsCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var news models.News
			if err := cursor.Decode(&news); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			newsList = append(newsList, news)
		}

		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"allNews": newsList})
}







func GetNewsByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	newsID := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(newsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var news models.News

	err = newsCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&news)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}

func UpdateNews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	newsID := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(newsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedNews models.News

	if err := c.ShouldBindJSON(&updatedNews); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updatedNews}

	result, err := newsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News updated successfully"})
}


func DeleteNews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	newsID := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(newsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	filter := bson.M{"_id": objectID}
	result, err := newsCollection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}
func GetNewsByCategory(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	category := c.Param("category")

	filter := bson.M{"category": category}

		
		var newsItems []models.News
		cursor, err := newsCollection.Find(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding news items"})
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var newsItem models.News
			if err = cursor.Decode(&newsItem); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding news item"})
				return
			}
			newsItems = append(newsItems, newsItem)
		}

		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating through news items"})
			return
		}

		c.JSON(http.StatusOK, newsItems)

}
