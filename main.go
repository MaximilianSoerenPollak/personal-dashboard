package main

import (
    "github.com/gin-gonic/gin"
)

type tasklist struct{
    Name string `json:"name"`
    Tasks []task `json:"tasks"`
}  

type task struct {
    Name string `json:"name"`
}
    

var list1 = tasklist{
    Name: "Tasklist1",
    Tasks: []task{},
}

func (tl *tasklist) AddTaskToList(t string) []task {
    tsk := task{t}
    tl.Tasks = append(tl.Tasks, tsk)
    return tl.Tasks
}

func main() {
    r := gin.Default()
    list1.AddTaskToList("TestingTheFunction")
    r.GET("/tab", func(c * gin.Context) { 
        c.JSON(200, gin.H {
            "message": "nine",
        })
    })
    for _,v := range list1.Tasks {
    r.GET("/tasks", func(c *gin.Context) {
        c.JSON(200, gin.H {
            "Tasks": v,
        })
    })
    }
    r.Run()
}
