package repository

import (
	"database/sql"
	"log"
	. "todo-app/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

func GetDataBaseConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitTable() error {
	db := GetDataBaseConnection()
	defer db.Close()
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        done BOOLEAN NOT NULL CHECK (done IN (0,1))
    );
    `
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil

}

func InsertTask(t Task) error {
	db := GetDataBaseConnection()
	defer db.Close()
	_, err := db.Exec("INSERT INTO tasks(title, done) VALUES (?, ?)", t.Title, t.Done)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func GetTasks() ([]Task, error) {
	db := GetDataBaseConnection()
	defer db.Close()
	var tasks = []Task{}
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Title, &task.Done)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTask(id int) (Task, error) {
	db := GetDataBaseConnection()
	defer db.Close()
	row, err := db.Query("SELECT * FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		return Task{}, err
	}
	var task Task
	for row.Next() {
		err := row.Scan(&task.Id, &task.Title, &task.Done)
		if err != nil {
			log.Fatal(err)
			return Task{}, err
		}
	}
	return task, nil
}

func UpdateTask(t Task, id int) error {
	db := GetDataBaseConnection()
	defer db.Close()
	_, err := db.Exec("UPDATE tasks SET title = ?, done = ? WHERE id = ?", t.Title, t.Done, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DeleteTask(id int) error {
	db := GetDataBaseConnection()
	defer db.Close()
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
