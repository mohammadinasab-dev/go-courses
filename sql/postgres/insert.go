package postgres

import (
	"database/sql"
	"time"
)

/*
	- SQL query placeholders, such as $1, $2, etc., are used to represent values that will be supplied later.
*/

func InsertData(db *sql.DB) error {
	teacherInsertQuery := "INSERT INTO teacher (name) VALUES ($1)"
	studentInsertQuery := "INSERT INTO student (name, teacher_id, created_at) VALUES ($1, $2, $3)"

	// Insert teachers
	teachers := []string{"John", "Alice", "Bob"}
	for _, teacher := range teachers {
		if _, err := db.Exec(teacherInsertQuery, teacher); err != nil {
			return err
		}
	}

	// Insert students
	students := map[string]string{
		"Student1": "John",
		"Student2": "Alice",
		"Student3": "Bob",
	}
	for student, teacher := range students {
		teacherIDQuery := "SELECT id FROM teacher WHERE name = $1"
		var teacherID int
		//todo context
		if err := db.QueryRow(teacherIDQuery, teacher).Scan(&teacherID); err != nil {
			return err
		}

		//todo context
		if _, err := db.Exec(studentInsertQuery, student, teacherID, time.Now().UTC()); err != nil {
			return err
		}
	}

	return nil
}
