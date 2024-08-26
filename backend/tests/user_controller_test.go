package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ritankarsaha/backend/controllers"
	"github.com/ritankarsaha/backend/database"
	"github.com/ritankarsaha/backend/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func TestGetUser(t *testing.T) {
	router := gin.Default()
	router.GET("/users/:user_id", controllers.GetUser())

	userID := primitive.NewObjectID()
	fullname := "John Doe"
	email := "johndoe@example.com"
	phone := "1234567890"
	role := "JOURNALIST"

	testUser := models.User{
		ID:         userID,
		Fullname:   &fullname,
		Email:      &email,
		Phone:      &phone,
		Role:       &role,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, _ = userCollection.InsertOne(context.TODO(), testUser)

	// Create a request to the endpoint
	req, _ := http.NewRequest("GET", "/users/"+userID.Hex(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check response
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]models.User
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, *testUser.Email, *response["user"].Email)
	assert.Equal(t, *testUser.Fullname, *response["user"].Fullname)
}
