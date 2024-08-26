package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ritankarsaha/backend/controllers"
	"github.com/ritankarsaha/backend/database"
	"github.com/ritankarsaha/backend/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	
	clientOptions, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = clientOptions.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongodb")
	return clientOptions
}

var Client *mongo.Client = DBinstance()

var commentCollection *mongo.Collection = database.OpenCollection(database.Client,"comments")


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


func tearDown(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	exitVal := m.Run()
	tearDown(Client)
	os.Exit(exitVal)
}