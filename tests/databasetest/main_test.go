package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFindDatabaseRecord(t *testing.T) {
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
