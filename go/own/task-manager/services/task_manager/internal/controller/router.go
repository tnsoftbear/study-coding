package controller

import (
	"fmt"
	"net/http"

	"task_manager/internal/redis_storage"
	"task_manager/internal/types"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	rm := redis_storage.NewRedisManager()

	r.GET("/task", func(c *gin.Context) {
		if tasks, err := rm.LoadTasks(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			response := gin.H{"tasks": tasks}
			c.JSON(http.StatusOK, response)
		}
	})

	r.GET("/task/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		task, err := rm.LoadTask(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
		} else if task == nil {
			c.JSON(http.StatusNotFound, gin.H{"id": id, "message": fmt.Sprintf("Task not found by id: %s", id)})
		} else {
			c.JSON(http.StatusOK, gin.H{"task": task})
		}
	})

	r.POST("/task", func(c *gin.Context) {
		var task types.Task = types.Task{}
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"task": task, "created": false, "message": err.Error()})
			return
		}

		if err := rm.SaveTask(task); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"task": task, "created": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"task": task, "created": true, "message": "Task created successfully"})
	})

	r.DELETE("/task/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		if task, _ := rm.LoadTask(id); task == nil {
			c.JSON(http.StatusNotFound, gin.H{"id": id, "deleted": false, "message": fmt.Sprintf("Task not found by id: %s", id)})
			return
		}

		if err := rm.DeleteTask(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": id, "deleted": true, "message": "Task successfully deleted"})
	})

	return r
}
