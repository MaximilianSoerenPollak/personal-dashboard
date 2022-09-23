package main

import (
	"dashboard/models"
	"github.com/gin-gonic/gin"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createTask(c *gin.Context) {
	c.JSON(200, gin.H{"message": "A new Record Created!"})
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
