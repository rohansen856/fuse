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
var commentCollection = setupMockCommentCollection()

func setupMockCommentCollection() *MockCommentCollection {
	return &MockCommentCollection{}
}

// MockCommentCollection simulates MongoDB collection methods
type MockCommentCollection struct{}

func (m *MockCommentCollection) InsertOne(ctx context.Context, comment interface{}) (interface{}, error) {
	return &mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}, nil
}

func (m *MockCommentCollection) Find(ctx context.Context, filter interface{}) (*MockCursor, error) {
	return &MockCursor{}, nil
}

func (m *MockCommentCollection) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return &mongo.DeleteResult{DeletedCount: 1}, nil
}

type MockCursor struct{}

func (mc *MockCursor) All(ctx context.Context, results interface{}) error {
	*results.(*[]models.Comment) = append(*results.(*[]models.Comment), models.Comment{
		ID:        primitive.NewObjectID(),
		NewsID:    primitive.NewObjectID(),
		Content:   "This is a test comment",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

func TestCreateComment(t *testing.T) {
	router := gin.Default()
	router.POST("/comments", controllers.CreateComment)

	comment := models.Comment{
		NewsID:    primitive.NewObjectID(),
		Content:   "This is a test comment",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	body, _ := json.Marshal(comment)
	req, _ := http.NewRequest("POST", "/comments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCommentsByNewsID(t *testing.T) {
	router := gin.Default()
	router.GET("/comments/:newsID", controllers.GetCommentsByNewsID)

	newsID := primitive.NewObjectID().Hex()
	req, _ := http.NewRequest("GET", "/comments/"+newsID, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteComment(t *testing.T) {
	router := gin.Default()
	router.DELETE("/comments/:commentID", controllers.DeleteComment)

	commentID := primitive.NewObjectID().Hex()
	req, _ := http.NewRequest("DELETE", "/comments/"+commentID, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}