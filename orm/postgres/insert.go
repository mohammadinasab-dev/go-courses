package postgres

import (
	"fmt"

	"gorm.io/gorm"
)

func Insert(db *gorm.DB) {

	// Create a new teacher
	teacher := Teacher{
		Name: "John Doe",
		Students: []Student{
			{Name: "Alice"},
			{Name: "Bob"},
		},
		Department: []Department{
			{Name: "mathematic"},
			{Name: "physic"},
		},
	}

	// Insert the teacher and its associated students into the database
	db.Create(&teacher)

	fmt.Println("Data inserted successfully!")
}
