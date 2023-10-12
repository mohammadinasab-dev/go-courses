package main

import (
	"fmt"
	"orm/postgres"
)

/*
Chain methods: Chain methods are used to build queries in a fluent and chainable manner.
These methods allow you to progressively add conditions, orderings, limits, and
other query parameters to your query.
Examples of chain methods in GORM include Where(), Order(), Limit(), and Offset().
Chain methods return a new modified query object, allowing you to chain multiple methods together.

Finisher methods: Finisher methods are used to actually execute the query and retrieve the results.
These methods bring the query to its conclusion and return the final result.
Examples of finisher methods in GORM include Find(), First(), Last(), Count(), and Delete().
Finisher methods typically return the queried records or relevant information based on the query.

New Session methods: New Session methods are used to create a new database session within GORM.
These methods establish a new connection to the database and return a session object that you can use to perform database operations.
Examples of New Session methods in GORM include New() and DB().
The session object allows you to perform database transactions and interact with database tables.



*/

func main() {

	db := postgres.NewGormPostgres()
	defer func() {
		dbc, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}
		dbc.Close()
	}()
	// postgres.Migrate(db)
	// postgres.Insert(db)
	res, err := postgres.GetTeacherByID(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	// if err := postgres.UpdateTeacherByID(db, 1, "mike"); err != nil {
	// 	fmt.Println(err)
	// }

	// if err := postgres.DeleteTeacherByID(db, 1); err != nil {
	// 	fmt.Println(err)
	// }

}
