package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
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

func GetOneTask(id int) (Task, error) {
	var searchedTask Task
	query, err := DB.Query("SELECT * FROM tasks WHERE id = ?", id)
	if err != nil {
		return searchedTask, err
	}
	query.Next()
	err = query.Scan(&searchedTask.ID, &searchedTask.Name, &searchedTask.Status)
	if err != nil {
		return searchedTask, err
	}
	return searchedTask, nil
}

func CreateTasks(n string) (int64, error) {
	r, err := DB.Exec("INSERT INTO tasks(name, status) values(?, 0)", n)
	if err != nil {
		log.Println(err, n)
	}
	id, err := r.LastInsertId()
	if err != nil {
		log.Println(err, n, id)
	}
	return id, nil
}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return err
	}
	DB = db
	return nil
}
