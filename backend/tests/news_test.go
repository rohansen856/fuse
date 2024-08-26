package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ritankarsaha/backend/controllers"
	"github.com/ritankarsaha/backend/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var newsCollection = setupMockNewsCollection()

func setupMockNewsCollection() *MockNewsCollection {
	return &MockNewsCollection{}
}

// MockNewsCollection simulates MongoDB collection methods
type MockNewsCollection struct{}

func (m *MockNewsCollection) InsertOne(ctx context.Context, news interface{}) (interface{}, error) {
	return &mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}, nil
}

func (m *MockNewsCollection) Find(ctx context.Context, filter interface{}) (*MockCursor, error) {
	return &MockCursor{}, nil
}

func (m *MockNewsCollection) FindOne(ctx context.Context, filter interface{}) *MockSingleResult {
	return &MockSingleResult{}
}

func (m *MockNewsCollection) UpdateOne(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	return &mongo.UpdateResult{MatchedCount: 1}, nil
}

func (m *MockNewsCollection) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return &mongo.DeleteResult{DeletedCount: 1}, nil
}

type MockCursor struct{}

func (mc *MockCursor) All(ctx context.Context, results interface{}) error {
	*results.(*[]models.News) = append(*results.(*[]models.News), models.News{
		ID:       primitive.NewObjectID(),
		Category: "Tech",
		Datetime: time.Now(),
		Headline: "Test Headline",
		Image:    "https://example.com/image.jpg",
		Related:  "Related News",
		Source:   "Test Source",
		Summary:  "Test Summary",
		URL:      "https://example.com/news",
	})
	return nil
}

func (mc *MockCursor) Close(ctx context.Context) error {
	return nil
}

func (mc *MockCursor) Next(ctx context.Context) bool {
	return false
}

func (mc *MockCursor) Decode(val interface{}) error {
	return nil
}

func (mc *MockCursor) Err() error {
	return nil
}

type MockSingleResult struct{}

func (msr *MockSingleResult) Decode(val interface{}) error {
	*val.(*models.News) = models.News{
		ID:       primitive.NewObjectID(),
		Category: "Tech",
		Datetime: time.Now(),
		Headline: "Test Headline",
		Image:    "https://example.com/image.jpg",
		Related:  "Related News",
		Source:   "Test Source",
		Summary:  "Test Summary",
		URL:      "https://example.com/news",
	}
	return nil
}

func TestCreateNews(t *testing.T) {
	router := gin.Default()
	router.POST("/news", controllers.CreateNews)

	news := models.News{
		Category: "Tech",
		Datetime: time.Now(),
		Headline: "Test News Headline",
		Image:    "https://example.com/image.jpg",
		Related:  "Related News",
		Source:   "Test Source",
		Summary:  "This is a test news summary",
		URL:      "https://example.com/news",
	}

	body, _ := json.Marshal(news)
	req, _ := http.NewRequest("POST", "/news", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetNews(t *testing.T) {
	router := gin.Default()
	router.GET("/news", controllers.GetNews)

	req, _ := http.NewRequest("GET", "/news", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetNewsByID(t *testing.T) {
	router := gin.Default()
	router.GET("/news/:id", controllers.GetNewsByID)

	newsID := primitive.NewObjectID().Hex()
	req, _ := http.NewRequest("GET", "/news/"+newsID, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateNews(t *testing.T) {
	router := gin.Default()
	router.PUT("/news/:id", controllers.UpdateNews)

	newsID := primitive.NewObjectID().Hex()
	updatedNews := models.News{
		Category: "Updated Tech",
		Datetime: time.Now(),
		Headline: "Updated Headline",
		Image:    "https://example.com/updated_image.jpg",
		Related:  "Updated Related News",
		Source:   "Updated Source",
		Summary:  "Updated Summary",
		URL:      "https://example.com/updated_news",
	}

	body, _ := json.Marshal(updatedNews)
	req, _ := http.NewRequest("PUT", "/news/"+newsID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteNews(t *testing.T) {
	router := gin.Default()
	router.DELETE("/news/:id", controllers.DeleteNews)

	newsID := primitive.NewObjectID().Hex()
	req, _ := http.NewRequest("DELETE", "/news/"+newsID, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetNewsByCategory(t *testing.T) {
	router := gin.Default()
	router.GET("/news/category/:category", controllers.GetNewsByCategory)

	req, _ := http.NewRequest("GET", "/news/category/Tech", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
