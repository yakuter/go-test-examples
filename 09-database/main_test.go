package main

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFind(t *testing.T) {
	// the db satisfy the sql.DB struct
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// our sql.DB must exec the follow query
	mock.ExpectQuery("SELECT * FROM tasks").
		// the number 3 must be on the query like "where id = 3"
		WithArgs(1)
		// TODO: define the values that need to be returned from the database

	myDB := NewRepository(db) // passes the mock to our code

	// run the code with the database mock
	if _, err := myDB.Find(1); err != nil {
		t.Errorf("something went wrong: %s", err.Error())
	}
}

func TestAdd(t *testing.T) {
	// the db satisfy the sql.DB struct
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var id int64
	task := Task{
		Title:       "React Native Öğren",
		StartDate:   time.Now(),
		DueDate:     time.Now(),
		Status:      true,
		Priority:    true,
		Description: "Mobil uygulama geliştirme artık günümüzün olmazsa olmazı.",
		CreatedAt:   time.Now(),
	}

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO tasks (title,start_date,due_date,status,priority,description,created_at) VALUES($1,$2,$3,$4,$5,$6,$7)`)).
		WithArgs(task.Title, task.StartDate, task.DueDate, task.Status, task.Priority, task.Description, task.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

	myDB := NewRepository(db) // passes the mock to our code

	lastID, err := myDB.Add(task)
	if err != nil {
		t.Errorf("something went wrong: %s", err.Error())
	}
	fmt.Println(lastID)
}
