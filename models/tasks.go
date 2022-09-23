package models

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Task struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Status int `json:"status"`
}

func GetTasks() ([]Task, error) {
    rows, err := DB.Query("SELECT * FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

   tasks := make([]Task, 0)
    for rows.Next() {
        singleTask := Task{}
        err = rows.Scan(&singleTask.ID,
        &singleTask.Name, &singleTask.Status)

        if err != nil {
            return nil, err
        }
        tasks = append(tasks, singleTask)
    }
    err = rows.Err()

    if err != nil {
        return nil, err
    }
    return tasks, nil 
}



func ConnectDatabase() error {
    db, err := sql.Open("sqlite3", "./sqlite.db")
    if err != nil {
        return err
    }
    DB = db 
    return nil 
}

