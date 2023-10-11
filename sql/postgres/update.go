package postgres

import "database/sql"

func SoftDeleteStudent(db *sql.DB, teacherID int) error {
	query := `
		UPDATE student
		SET deleted_at = NOW()
		WHERE id = $1
	`

	_, err := db.Exec(query, teacherID)
	if err != nil {
		return err
	}

	return nil
}
