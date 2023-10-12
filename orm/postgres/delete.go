package postgres

import (
	"errors"

	"gorm.io/gorm"
)

/*
	- soft delete and cascade
	- constraint over soft deleted
*/

func DeleteTeacherByID(db *gorm.DB, teacherID uint) error {
	result := db.Where("id = ?", teacherID).Delete(&Teacher{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Teacher not found")
	}
	return nil
}

func DeleteStudentByID(db *gorm.DB, studentID uint) error {
	result := db.Where("id = ?", studentID).Delete(&Student{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Student not found")
	}
	return nil
}
