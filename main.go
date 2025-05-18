package main

import (
	"github.com/gin-gonic/gin"
	"task_service/handlers"
)

func main() {
	r := gin.Default()
	taskHandler := handlers.NewTaskHandler()

	// Routes
	r.GET("/tasks", taskHandler.GetTasks)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	r.Run(":8080")
}
