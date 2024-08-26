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
	"github.com/ritankarsaha/backend/models"
	"github.com/ritankarsaha/backend/routes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// Setup a mock MongoDB server using mtest
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Initialize Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Mock the user collection
	userCollection = mt.Coll

	// Setup routes
	routes.UserRoutes(router)

	t.Run("User Found", func(t *testing.T) {
		// Setup mock response
		userId := primitive.NewObjectID()
		user := models.User{
			ID:        userId,
			Fullname:  StringPointer("John Doe"),
			Email:     StringPointer("johndoe@example.com"),
			Phone:     StringPointer("1234567890"),
			Role:      StringPointer("CONSUMER"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "userdb.user", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: user.ID},
			{Key: "fullname", Value: *user.Fullname},
			{Key: "email", Value: *user.Email},
			{Key: "phone", Value: *user.Phone},
			{Key: "role", Value: *user.Role},
			{Key: "created_at", Value: user.CreatedAt},
			{Key: "updated_at", Value: user.UpdatedAt},
		}))

		// Perform the request
		req, _ := http.NewRequest(http.MethodGet, "/user/"+userId.Hex(), nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		// Assert the response
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]models.User
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, user.ID, response["user"].ID)
		assert.Equal(t, *user.Fullname, *response["user"].Fullname)
	})

	t.Run("User Not Found", func(t *testing.T) {
		// Setup mock response for no documents found
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "User not found",
		}))

		// Perform the request
		req, _ := http.NewRequest(http.MethodGet, "/user/nonexistentid", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		// Assert the response
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestUpdateUser(t *testing.T) {
	// Setup a mock MongoDB server using mtest
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Initialize Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Mock the user collection
	userCollection = mt.Coll

	// Setup routes
	routes.UserRoutes(router)

	t.Run("Successful Update", func(t *testing.T) {
		// Setup mock response
		userId := primitive.NewObjectID()

		// Define update data
		updateData := bson.M{
			"fullname": "Jane Doe",
			"phone":    "0987654321",
		}

		updateDataBytes, _ := json.Marshal(updateData)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		// Perform the request
		req, _ := http.NewRequest(http.MethodPatch, "/user/"+userId.Hex(), bytes.NewBuffer(updateDataBytes))
		req.Header.Set("Content-Type", "application/json")
		req = req.WithContext(context.WithValue(req.Context(), "user_role", "MANAGER"))
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		// Assert the response
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Unauthorized Address Update", func(t *testing.T) {
		// Setup mock response
		userId := primitive.NewObjectID()

		// Define update data with address
		updateData := bson.M{
			"address": "New Address",
		}

		updateDataBytes, _ := json.Marshal(updateData)

		// Perform the request
		req, _ := http.NewRequest(http.MethodPatch, "/user/"+userId.Hex(), bytes.NewBuffer(updateDataBytes))
		req.Header.Set("Content-Type", "application/json")
		req = req.WithContext(context.WithValue(req.Context(), "user_role", "CONSUMER"))
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		// Assert the response
		assert.Equal(t, http.StatusForbidden, rec.Code)
	})

	t.Run("Invalid Input", func(t *testing.T) {
		// Setup mock response
		userId := primitive.NewObjectID()

		// Define invalid update data
		invalidData := []byte(`{invalid-json}`)

		// Perform the request
		req, _ := http.NewRequest(http.MethodPatch, "/user/"+userId.Hex(), bytes.NewBuffer(invalidData))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		// Assert the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

// Helper function to create string pointers
func StringPointer(s string) *string {
	return &s
}
