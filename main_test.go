package main

import (
	// "testing"
	// "github.com/stretchr/testify"
	// "net/http"
	// "net/http/httptest"

	"dashboard/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// func TestDatabaseConnection(t *testing.T) {
// 	db, _ := sql.Open("sqlite3", "./sqlite.db")
//     fmt.Println(db.Ping())
// }

var DB_TEST *sql.DB

func TestDataBase(t *testing.T) {
	db, _ := sql.Open("sqlite3", "test1.db")
	DB_TEST = db
	fmt.Println("The database was created")
}

func TestCreateTable(t *testing.T) {
	log.SetOutput(file)
	query, err := DB_TEST.Prepare("CREATE TABLE IF NOT EXISTS test (ID INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT, Status INTEGER);")
	if err != nil {
		log.Println("Failed to make test table in test db", err)
		t.FailNow()
	}
	query.Exec()
	_, err = DB_TEST.Query("SELECT * FROM test;")
	if err != nil {
		fmt.Println("could not select from test")
	}
}

func TestInsertIntoDB(t *testing.T) {
	log.SetOutput(file)
	returnedTask := models.Task{}
	inputTask := models.Task{
		Name:   "TestTask",
		Status: 2,
	}
	assert := assert.New(t)
	r, err := DB_TEST.Exec("INSERT INTO test(Name, Status) VALUES (?, ?);", inputTask.Name, inputTask.Status)
	if err != nil {
		log.Println("Failed to insert value into test", err)
		t.FailNow()
	}
	id, err := r.LastInsertId()
	if err != nil {
		log.Println("Failed to get last inserted ID from test db(tasks).", err)
		t.FailNow()
	}
	err = DB_TEST.QueryRow("SELECT * FROM test WHERE ID = ?;", id).Scan(&returnedTask.ID, &returnedTask.Name, &returnedTask.Status)
	if err != nil {
		log.Printf("Failed to query test db(tasks) for task with id = %d and err: %s", id, err)
		t.FailNow()
	}
	if err != nil {
		log.Println("Failed to scan the Query that we returned in test.", err)
		t.FailNow()
	}
	assert.Equal(returnedTask.ID, 1, "The ID should be 1")
	assert.Equal(returnedTask.Name, "TestTask", "The Name should be 'TestTask'")
	assert.Equal(returnedTask.Status, 2, "The status should be 2")
}

func TestUnitTestingGETRoute(t *testing.T, *gin.Context, httptest.ResponseRecorder) {
    log.SetOutput(file)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    req := http.Request{
        URL: "/",
        Method: "GET",
    }
    
         
}

func TestDeleteDB(t *testing.T) {
	defer DB_TEST.Close()
	err := os.Remove("test1.db")
	if err != nil {
		log.Fatalln(err)
	}
}
