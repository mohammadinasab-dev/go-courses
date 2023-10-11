package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)


/*
	- Performance improvement: When you prepare a statement using Prepare, the database server parses, compiles, and optimizes the SQL query. 
	  The prepared statement is then stored on the database server, allowing it to be reused multiple times. 
	  This eliminates the need for the database server to repeat the query parsing and 
	  optimization steps for each execution, resulting in improved performance, 
	  especially when executing the same query multiple times with different parameter values.
*/

func InsertData01(db *sql.DB) error {
	teacherInsertQuery := "INSERT INTO teacher (name) VALUES ($1)"
	studentInsertQuery := "INSERT INTO student (name, teacher_id) VALUES ($1, $2)"

	// Prepare teacher statement
	teacherStmt, err := db.Prepare(teacherInsertQuery)
	if err != nil {
		return err
	}
	defer teacherStmt.Close()

	// Prepare student statement
	studentStmt, err := db.Prepare(studentInsertQuery)
	if err != nil {
		return err
	}
	defer studentStmt.Close()

	// Begin the transaction
	txn, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		return err
	}

	// Insert teachers
	teachers := []Teacher{
		{Name: "John"},
		{Name: "Alice"},
		{Name: "Bob"},
	}
	for _, teacher := range teachers {
		if _, err := txn.Exec(teacherInsertQuery, teacher.Name); err != nil {
			txn.Rollback()
			return err
		}
	}

	// Insert students
	students := map[string]Student{
		"Student1": {Name: "John", Teacher: "John"},
		"Student2": {Name: "Alice", Teacher: "Alice"},
		"Student3": {Name: "Bob", Teacher: "Bob"},
	}
	for student, info := range students {
		teacherIDQuery := "SELECT id FROM teacher WHERE name = $1"
		var teacherID int
		if err := db.QueryRow(teacherIDQuery, info.Teacher).Scan(&teacherID); err != nil {
			txn.Rollback()
			return err
		}

		if _, err := txn.Exec(studentInsertQuery, student, teacherID); err != nil {
			txn.Rollback()
			return err
		}
	}

	// Commit the transaction
	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}