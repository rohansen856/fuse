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
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// func HashPassword(password string) string {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		log.Println("Error hashing password:", err)
// 		return ""
// 	}
// 	return string(bytes)
// }

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")

		if err := helper.MathUserTypeToUid(ctx, userId);
		 err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		var user models.User
		err := userCollection.FindOne(c, bson.M{"_id": userId}).Decode(&user)
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

// func Signup() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
// 		defer cancel()

// 		var user models.User

// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 		}
// 		if err := validate.Struct(user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 		}
// 		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 		}

// 		if count > 0 {
// 			c.JSON(http.StatusConflict, gin.H{
// 				"error": "Email already exists",
// 			})
// 		}
// 		password := HashPassword(*user.Password)
// 		user.Password = &password
// 		count1, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 		}
// 		if count1 > 0 {
// 			c.JSON(http.StatusConflict, gin.H{
// 				"error": "Phone already exists",
// 			})
// 		}
// 		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		user.ID = primitive.NewObjectID()
// 		// user.User_id = &user.ID.Hex()
// 		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_Type, *&user.User_id)
// 		user.Refresh_Token = &refreshToken
// 		user.Token = &token

// 		// user.Password = helper.HashPassword(*user.Password)
// 		result, err := userCollection.InsertOne(ctx, user)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 		}
// 		defer cancel()

// 		c.JSON(http.StatusCreated, result)

// 	}
// }

// func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
// 	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
// 	if err != nil {
// 		return false, "Wrong password"
// 	}
// 	return true, ""
// }

// func Login() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 		defer cancel()

// 		var user models.User
// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Invalid request",
// 			})
// 			return
// 		}

// 		var foundUser models.User
// 		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Invalid email or password",
// 			})
// 			return
// 		}

// 		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
// 		if !passwordIsValid {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": msg,
// 			})
// 			return
// 		}

// 		token, refreshToken, err := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_Type, *&foundUser.User_id)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		helper.UpdateAllTokens(token, refreshToken, *&foundUser.User_id)

// 		err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, foundUser)
// 	}
// }