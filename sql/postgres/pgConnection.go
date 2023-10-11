package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"text/template"
	"time"

	_ "github.com/lib/pq"
)

/*
	- constructor is a kind of factory pattern
	- `_` is just to call init function
	- ssl is about secure connection
	- text/template
	- time zone
	- DSN stands for "Data Source Name."
*/

func NewPostgres() {
	// Define the connection string template
	const connStrTemplate = "postgres://{{.User}}:{{.Password}}@{{.Host}}/{{.Database}}?sslmode={{.SSLMode}}&timezone={{.TimeZone}}"

	// Define the connection parameters
	params := struct {
		User     string
		Password string
		Host     string
		Database string
		SSLMode  string
		TimeZone string
	}{
		User:     "admin",
		Password: "password",
		Host:     "localhost:5432",
		Database: "quera",
		SSLMode:  "disable",
		TimeZone: "Asia/Tehran",
	}

	// Parse the connection string template
	connStrTmpl, err := template.New("connStr").Parse(connStrTemplate)
	if err != nil {
		panic(err)
	}

	// Render the connection string
	var connStrBuilder strings.Builder
	if err := connStrTmpl.Execute(&connStrBuilder, params); err != nil {
		panic(err)
	}
	dsn := connStrBuilder.String()

	fmt.Println(dsn)

	// Open a connection to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Second)
	db.SetConnMaxIdleTime(30 * time.Second)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(2)

	fmt.Println("Connected to the database!")
}
