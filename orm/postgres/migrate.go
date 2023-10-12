package postgres

import (
	"fmt"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name       string    `gorm:"size:50;not null;unique"`
	Students   []Student `gorm:"constraint:OnDelete:CASCADE"`
	Department []Department
}

type Student struct {
	gorm.Model
	Name      string `gorm:"size:50;not null;unique"`
	TeacherID uint
}

type Department struct {
	gorm.Model
	Name      string
	TeacherID uint
}

func (t *Teacher) BeforeCreate(tx *gorm.DB) (err error) {
	// This hook will be called before creating a new teacher
	// Add your custom logic here
	fmt.Println("Before creating a teacher: ", t.Name)
	return nil
}

func (t *Teacher) AfterCreate(tx *gorm.DB) (err error) {
	// This hook will be called after creating a new teacher
	// Add your custom logic here
	fmt.Println("After creating a teacher: ", t.Name)
	return nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Department{})
}
