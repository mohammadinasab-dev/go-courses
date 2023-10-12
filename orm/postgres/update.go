package postgres

import (
	"errors"

	"gorm.io/gorm"
)


/*
	use context
*/


func UpdateTeacherByID(db *gorm.DB, teacherID uint, newName string) error {
	result := db.Model(&Teacher{}).Where("id = ?", teacherID).Update("name", newName)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Teacher not found")
	}
	return nil
}

func UpdateStudentByID(db *gorm.DB, studentID uint, newName string) error {
	result := db.Model(&Student{}).Where("id = ?", studentID).Update("name", newName)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Student not found")
	}
	return nil
}
