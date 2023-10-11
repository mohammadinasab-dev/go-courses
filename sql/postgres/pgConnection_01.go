package postgres

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

/*
	- singleton


	TCP connection string:
	URL: tcp://localhost:8080
	Scheme/Protocol: tcp
	Host/Domain: localhost
	Port: 8080

	HTTP connection string:
	URL: http://www.example.com:80/path/file.html
	Scheme/Protocol: http
	Host/Domain: www.example.com
	Port: 80
	Path: /path/file.html

	DB connection string
	URL: postgres://username:password@localhost:5432/database_name
	Scheme/Protocol: postgres
	Username: username
	Password: password
	Host/Domain: localhost:5432 (assuming the PostgreSQL server is running on the local machine using the default port 5432)
	Database name: database_name
	RawQuery: sslmode=require (optional, used to specify the SSL mode)
*/

// var once sync.Once

func NewPostgres01() (db *sql.DB) {

	// once.Do(func() {
	// Database credentials
	tehranTimezone, _ := time.LoadLocation("Asia/Tehran")
	dbURL := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword("admin", "password"),
		Host:     "localhost",
		Path:     "quera",
		RawQuery: "sslmode=disable&timezone=" + tehranTimezone.String(),
	}

	// Convert URL to connection string
	connStr := dbURL.String()

	fmt.Println(connStr)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Execute a test query
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to execute test query:", err)
		return
	}

	fmt.Println("Successfully connected to the database!")
	// })

	return
}
