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
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock data and variables
var commentCollection = setupMockCommentCollection()

func setupMockCommentCollection() *MockCommentCollection {
	return &MockCommentCollection{}
}

// MockCommentCollection simulates MongoDB collection methods
type MockCommentCollection struct{}

func (m *MockCommentCollection) InsertOne(ctx context.Context, comment interface{}) (interface{}, error) {
	return nil, nil // Simulate successful insert
}

func (m *MockCommentCollection) Find(ctx context.Context, filter interface{}) (*MockCursor, error) {
	return &MockCursor{}, nil // Simulate successful find
}

func (m *MockCommentCollection) DeleteOne(ctx context.Context, filter interface{}) (*DeleteResult, error) {
	return &DeleteResult{DeletedCount: 1}, nil // Simulate successful delete
}

type MockCursor struct{}

func (mc *MockCursor) All(ctx context.Context, results interface{}) error {
	return nil // Simulate successful cursor operation
}

func (mc *MockCursor) Close(ctx context.Context) error {
	return nil // Simulate cursor close
}

type DeleteResult struct {
	DeletedCount int64
}

func TestCreateComment(t *testing.T) {
	router := gin.Default()
	router.POST("/comments", CreateComment)

	comment := map[string]interface{}{
		"post_id":   primitive.NewObjectID(),
		"content":   "This is a test comment",
		"user_id":   primitive.NewObjectID(),
		"user_name": "Test User",
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
	router.GET("/news/:newsID/comments", GetCommentsByNewsID)

	newsID := primitive.NewObjectID().Hex()
	req, _ := http.NewRequest("GET", "/news/"+newsID+"/comments", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteComment(t *testing.T) {
	router := gin.Default()
	router.DELETE("/comments/:commentID", DeleteComment)

	commentID := primitive.NewObjectID().Hex()
	req, _ := http.NewRequest("DELETE", "/comments/"+commentID, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

