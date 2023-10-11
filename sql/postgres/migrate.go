package postgres

import "database/sql"

func Migrate(db *sql.DB) {
	// Create 'teacher' table
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS teacher (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL
)`)
	if err != nil {
		panic(err)
	}

	// Create 'student' table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS student (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL UNIQUE,
	teacher_id INT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (teacher_id) REFERENCES teacher(id) ON DELETE CASCADE
)`)
	if err != nil {
		panic(err)
	}

}
