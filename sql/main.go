package main

import (
	"fmt"
	"os"
	"sql/postgres"
)

func main() {

	db := postgres.NewPostgres01()
	defer db.Close()

	postgres.Migrate(db)

	fmt.Println(postgres.InsertData(db))

	students, err := postgres.GetStudentsByTeacher(db, 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for _, student := range students {
		fmt.Println(student)
	}

	if err := postgres.SoftDeleteStudent(db, 1); err != nil {
		fmt.Println(err)
	}

}
