package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Set up database connection parameters
	db, err := sql.Open("mysql", "sergey:sergey@tcp(localhost:3306)/photo_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err.Error())
	} else {
		fmt.Println("Database connection successful!")
	}

	createTable := `CREATE TABLE users (
        id INT NOT NULL AUTO_INCREMENT,
        name VARCHAR(50) NOT NULL,
        email VARCHAR(100) NOT NULL,
        PRIMARY KEY (id)
    );`

	_, err = db.Exec(createTable)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Table created successfully!")
}
