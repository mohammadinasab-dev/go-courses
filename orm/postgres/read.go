package postgres

import (
	"gorm.io/gorm"
)

/*
	- The Preload function in GORM is used to load associations (i.e., related models) along with the main model in a single database query.
	  It allows you to eager-load associated data, reducing the number of queries needed to fetch related records.

	- First method retrieves the first record that matches the provided condition without considering any specific ordering.

*/

func GetTeacherByID(db *gorm.DB, teacherID uint) (*Teacher, error) {
	teacher := &Teacher{}
	if err := db.Preload("Students").Preload("Department").Order("created_at desc").First(teacher, teacherID).Error; err != nil {
		return nil, err
	}
	return teacher, nil
}


func GetTeachersByName(db *gorm.DB, name string) ([]Teacher, error) {
	var teachers []Teacher
	query := "SELECT * FROM teachers WHERE name = ?"
	err := db.Raw(query, name).Scan(&teachers).Error
	if err != nil {
		return nil, err
	}
	return teachers, nil
}