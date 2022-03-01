package controller

import (
	"accountability_back/config"
	"accountability_back/dto/requests"
	"accountability_back/model"
	"accountability_back/responses"
	"accountability_back/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var userCollection = config.GetCollection(config.DB, "users")
var validate = validator.New()

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		var request requests.LoginRequest
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "Error",
				Data:    gin.H{"data": err.Error()},
			})
			return
		}
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&request); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    gin.H{"data": validationErr.Error()},
			})
			return
		}
		// find user with provided id and unmarshal doc into user obj or return err if any
		var user model.User
		err := userCollection.FindOne(ctx, bson.M{"email": request.Email}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				responses.APIResponse{
					Status:  http.StatusBadRequest,
					Message: "error", Data: gin.H{"data": err.Error()}})
			return
		}

		var userPassword = user.Password
		err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(request.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest,
				responses.APIResponse{
					Status:  http.StatusBadRequest,
					Message: "error", Data: gin.H{"data": err.Error()}})
			return
		}
		token := service.JWTAuthService().GenerateToken(user.Name, user.Email)
		c.JSON(http.StatusOK,
			responses.APIResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data: gin.H{
					"user":  user,
					"token": token,
				},
			},
		)
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		var user model.User
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "Error",
				Data:    gin.H{"data": err.Error()},
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    gin.H{"data": validationErr.Error()},
			})
			return
		}
		var password = user.Password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		newUser := model.User{
			Id:       primitive.NewObjectID(),
			Name:     user.Name,
			Email:    user.Email,
			Password: string(hashedPassword),
		}
		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    gin.H{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.APIResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    gin.H{"data": result},
		})
	}
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user model.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		// find user with provided id and unmarshal doc into user obj or return err if any
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error", Data: gin.H{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data:    gin.H{"data": user}})
	}
}
