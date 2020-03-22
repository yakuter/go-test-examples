package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err error
)

type Task struct {
	ID          int
	Title       string
	StartDate   time.Time
	DueDate     time.Time
	Status      bool
	Priority    bool
	Description string
	CreatedAt   time.Time
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "12345"
	dbName := "gotest"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Repository interface {
	Find(id int) (Task, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: dbConn()}
}

func (r repository) Find(id int) (Task, error) {
	var task Task

	rows, _ := r.db.Query("SELECT * FROM tasks WHERE id=?", id)
	defer rows.Close()

	// Eğer kayıt yoksa
	if rows.Next() == false {
		return task, errors.New("Kayıt bulunamadı!")
	}

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Title, &task.StartDate, &task.DueDate, &task.Status, &task.Priority, &task.Description, &task.CreatedAt)
		if err != nil {
			return task, errors.New("Kayıtlar eşleştirilemedi!")
		}
	}

	if err = rows.Err(); err != nil {
		return task, err
	}

	return task, nil
}

func main() {
	db := dbConn()
	myDB := NewRepository(db)
	var task Task

	// 1 nolu kayıdı bul
	if task, err = myDB.Find(1); err != nil {
		log.Println(err)
		return
	}

	// Sonucu ekrana güzel bir şekilde yazalım
	jsonResult, _ := json.MarshalIndent(task, "", "    ")
	fmt.Println(string(jsonResult))
}
