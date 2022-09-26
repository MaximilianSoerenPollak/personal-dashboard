package main

import (
	"dashboard/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createTask(c *gin.Context) {
	var task models.Task
	task.Status = 0
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	if err = c.BindJSON(&task); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		log.Println("This is task", task.Name)
		models.CreateTasks(task.Name)
		c.IndentedJSON(http.StatusCreated, task)
	}
}

func readTask(c *gin.Context) {
	tasks, err := models.GetTasks()
	checkErr(err)

	if tasks == nil {
		c.JSON(404, gin.H{"error": "No records found"})
		return
	} else {
		c.IndentedJSON(200, gin.H{"data": tasks})
	}
}

func updateTask(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Updated!"})
}

func deleteTask(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Deleted!"})
}

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()
	router := r.Group("/tasks")
	{
		router.POST("/create", createTask)
		router.GET("/", readTask)
		router.POST("update/:id", updateTask)
		router.DELETE("/delete/:id", deleteTask)
	}
	r.Run()
}
