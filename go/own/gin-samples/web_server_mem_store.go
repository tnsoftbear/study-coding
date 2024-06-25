package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id      	string `json:"id"`
	Name		string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

var (
	taskMap map[string]Task = make(map[string]Task)
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/task", func(c *gin.Context) {
		tasks := []Task{}
		for _, task := range taskMap {
			tasks = append(tasks, task)
		}
		json := gin.H{"tasks":tasks}
		c.JSON(http.StatusOK, json)
	})

	r.GET("/task/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		task, ok := taskMap[id]
		if ok {
			c.JSON(http.StatusOK, gin.H{"task": task})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"id": id, "message": fmt.Sprintf("Task not found by id: %s", id)})
		}
	})

	r.POST("/task", func(c *gin.Context) {
		var task Task = Task{}
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusOK, map[string]any{"task": task, "created": false, "message": err.Error()})
		} else {
			taskMap[task.Id] = task
			c.JSON(http.StatusCreated, gin.H{"task": task, "created": true, "message": "Task created successfully"})
		}
	})

	r.DELETE("/task/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		if _, ok := taskMap[id]; !ok {
			c.JSON(http.StatusNotFound, gin.H{"id": id, "deleted": false, "message": fmt.Sprintf("Task not found by id: %s", id)})
		} else {
			delete(taskMap, id)
			c.JSON(http.StatusOK, gin.H{"id": id, "deleted": true, "message": "Task successfully deleted"})
		}
	})

	return r
}

func main() {
 	r := setupRouter()
 	r.Run(":8080")
}