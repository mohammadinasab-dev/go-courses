package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func GetStudentsByTeacher(db *sql.DB, teacherID int) ([]Student, error) {
	query := `
		SELECT name, teacher_id
		FROM student
		WHERE teacher_id = $1
	`

	rows, err := db.QueryContext(context.TODO(), query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []Student{}
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.Name, &student.Teacher)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}