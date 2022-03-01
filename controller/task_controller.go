package controller

import (
	"accountability_back/config"
	"accountability_back/model"
	"accountability_back/responses"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

var taskCollection = config.GetCollection(config.DB, "tasks")

var countryTz = map[string]string{
	"Nigeria": "Africa/Lagos",
}

func TimeIn(name string) time.Time {
	loc, err := time.LoadLocation(countryTz[name])
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		var task model.Task
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "Error",
				Data:    gin.H{"data": err.Error()},
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&task); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    gin.H{"data": validationErr.Error()},
			})
			return
		}
		newTask := model.Task{
			Id:          primitive.NewObjectID(),
			Name:        task.Name,
			CreatorId:   task.CreatorId,
			Description: task.Description,
			Completed:   false,
			Date:        TimeIn("Nigeria"),
		}

		result, err := taskCollection.InsertOne(ctx, newTask)
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

func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		taskId := c.Param("taskId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(taskId)

		result, err := taskCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    gin.H{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.APIResponse{
					Status:  http.StatusNotFound,
					Message: "error",
					Data:    gin.H{"data": "No task found with specified Id"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{
				Status:  http.StatusNoContent,
				Message: "success",
				Data:    gin.H{"data": ""}},
		)
	}
}

func GetAllMyTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		log.Println("userId", userId)
		var tasks []model.Task
		defer cancel()

		results, err := taskCollection.Find(ctx, bson.M{"creatorid": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    gin.H{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			log.Println("running in stream")
			var task model.Task
			log.Println("task", task)
			if err = results.Decode(&task); err != nil {
				c.JSON(http.StatusInternalServerError,
					responses.APIResponse{
						Status:  http.StatusInternalServerError,
						Message: "error",
						Data:    gin.H{"data": err.Error()}})
			}

			tasks = append(tasks, task)

		}

		c.JSON(http.StatusOK,
			responses.APIResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data: gin.H{
					"taskCount": len(tasks),
					"data":      tasks,
				},
			},
		)

	}
}

func GetMyTaskForToday() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		log.Println("userId", userId)
		var tasks []model.Task
		defer cancel()

		results, err := taskCollection.Find(ctx, bson.M{"creatorid": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    gin.H{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			log.Println("running in stream")
			var task model.Task
			log.Println("task", task)
			if err = results.Decode(&task); err != nil {
				c.JSON(http.StatusInternalServerError,
					responses.APIResponse{
						Status:  http.StatusInternalServerError,
						Message: "error",
						Data:    gin.H{"data": err.Error()}})
			}
			if task.Date.Truncate(24 * time.Hour).Equal(TimeIn("Nigeria").Truncate(24 * time.Hour)) {
				log.Println("date match")
				tasks = append(tasks, task)
			}

		}

		c.JSON(http.StatusOK,
			responses.APIResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data: gin.H{
					"taskCount": len(tasks),
					"data":      tasks,
				},
			},
		)

	}
}
