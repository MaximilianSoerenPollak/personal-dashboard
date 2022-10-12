package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type ParamError struct {
	function string
	params   []string
	Err      error
}

func (pa *ParamError) Error() string {
	return fmt.Sprintf("Error %v, occured in function %v, params %v involved", pa.Err, pa.function, pa.params)
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

// Use QueryRow instead of query maybe?
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

func CreateTasks(n string, x int) (Task, error) {
	r, err := DB.Exec("INSERT INTO tasks(name, status) values(?, ?)", n, x)
	if err != nil {
		log.Println(err, n)
	}

	id, err := r.LastInsertId()
	if err != nil {
		log.Println(err, n, id)
	}
	task, err := GetOneTask(int(id))
	if err != nil {
		log.Println(err, n, id, "Function: GetOneTask")
	}
	return task, nil
}

// UpdateTask function    status(st) and name(n) are optional' inputs
//TO-DO FIX: Does not error out if route does not work (sends back HTTP 200)
// TO-DO: PArameters are still not read correctly at all (always come up as null PARAM)
// Resource here: https://blog.petehouston.com/parse-query-string-in-gin-web-application/
func UpdateTask(id int, st, n string) (Task, error) {
	switch {
	case (n == "" && st == "0"):
		{
			log.Println(id, st, n, "Function: UpdateTask, n and st not '' ")
			return Task{}, &ParamError{
				function: "UpdateTask",
				params:   []string{"Name", "Status"},
				Err:      fmt.Errorf("params were empty"),
			}
		}
	case st == "0":
		{
			_, err := DB.Exec("UPDATE tasks SET name=? WHERE id=?", n, id)
			// This is not a properly handeled error I think. But for MVP is aight.
			if err != nil {
				log.Println(err, id, st, n, "Function: UpdateTask, st=''")
			}
		}
	case n == "":
		{
			st_int, err := strconv.Atoi(st)
			if err != nil {
				log.Println("conversion of st failed")
			}
			// This is not a properly handeled error I think. But for MVP is aight.
			_, err = DB.Exec("UPDATE tasks SET status=? WHERE id=?", st_int, id)
			if err != nil {
				log.Println(err, id, st, n, "Function: UpdateTask, n=''")
			}
		}
	default:
		{
			_, err := DB.Exec("UPDATE tasks SET name=?, status=? WHERE id=?", n, st, id)
			if err != nil {
				log.Println(err, id, st, n, "Function: UpdateTask failed at inserting data into DB")
			}
		}
	}
	// Then get the ID of the changed task. => Need to see if this is a good way or not.
	task, err := GetOneTask(id)
	if err != nil {
		log.Println(err, id, st, n, "Function: UpdateTask")
	}
	return task, nil
}

func DeleteTask(id int) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		log.Println(err, "Function: DeleteTask")
		return err
	}
	return nil
}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return err
	}
	DB = db
	return nil
}
