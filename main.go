package main

import (
	"dashboard/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// checkErr function    Small custom Error logger function.
// TO-DO Need to re-do this function so that it has the creation of the error log file etc.
// TO-DO Make sure to use this fucntion everywhere so it's more cohesive in the code.
var	file, err = os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createTask(c *gin.Context) {
	var task models.Task
    log.SetOutput(file)

	if err != nil {
		log.Fatal(err)
	}
	if err = c.BindJSON(&task); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		log.Println("This is task", task.Name)
		task, err := models.CreateTasks(task.Name, task.Status)
		checkErr(err)
		c.IndentedJSON(http.StatusCreated, task)

	}
}

// readTask function    This is to read all tasks in the Database
func readTask(c *gin.Context) {
	tasks, err := models.GetTasks()
	checkErr(err)

	if tasks == nil {
		c.JSON(404, gin.H{"error": "no records found"})
		return
	} else {
		c.IndentedJSON(200, gin.H{"data": tasks})
	}
}

// readOneTask function    This is to read just ONE task in the  DB.
func readOneTask(c *gin.Context) {
	id := c.Param("id")
	// Converting the string param 'id' to an int via 'strconv.Atoi'
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not convert ID to int"})
	}
	task, err := models.GetOneTask(id_int)
	checkErr(err)
	if task.Name == "" && task.Status == 0 {
		c.JSON(404, gin.H{"error": "no records found"})
		return
	} else {
		c.IndentedJSON(200, gin.H{"data": task})
	}

}

// updateTask function    TO-DO: Still need to implement this function
func updateTask(c *gin.Context) {
    log.SetOutput(file)
	test, ok := c.Params.Get("name")
	if ok {
		log.Println("This is the param name in a test", test)
	} else {
		log.Println("The param could not be gotten it's empty. ok came back false")
	}
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Please make sure you attach the ID parameter"})
	}
	name := c.Param("name")
	log.Println("This is the param name in updateTask after it was read", name)
	status := c.Param("status")
	log.Println("This is the param status in updateTask after it was read", status)
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Could not convert ID to int"})
	}
	task, err := models.UpdateTask(id_int, status, name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Something went horribly wrong"})
	}
	c.JSON(200, gin.H{"message": "Record Updated!", "Updated Record": task})
}

// deleteTask function    TO-DO: Still need to implement this function
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
        router.POST("/update/:id", updateTask)
		router.DELETE("/delete/:id", deleteTask)
	}
	r.Run()
}
