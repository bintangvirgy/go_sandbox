package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// import for its side effect to init sql connection use mysql driver. So when we use sql open and use mysql, sql package is already have mysql initiated
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "go:golang@(127.0.0.1:50001)/golang_dev?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// create table
	// wow, so golang can specify code block??
	{
		query := `
			CREATE OR REPLACE TABLE users(
				id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
				username VARCHAR(255) NOT NULL,
				password VARCHAR(255) NOT NULL,
				created_at DATETIME	
			);
		`

		// execute create table
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}

	}

	// insert new user
	{
		username := "admin"
		password := "admin123"
		created_at := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) values (?,?,?);`, username, password, created_at)

		if err != nil {
			log.Fatal(err)
		}

		userID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(userID)
	}

	// query a row from table
	{
		// declare var to store data
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		// query use id that we get from last insert id.
		// user query row to extract only 1 row
		// to store selected data use scan, and it must be ordered correctly
		query := `SELECT id, username, password, created_at FROM users WHERE id = ?`

		err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)

		if err != nil {
			log.Fatal(err)
		}

		// example of wrong order
		fmt.Println(id, " ", username, " ", password, " ", createdAt)

		err2 := db.QueryRow(query, 1).Scan(&id, &password, &username, &createdAt)

		if err2 != nil {
			log.Fatal(err)
		}

		fmt.Println(id, " ", username, " ", password, " ", createdAt)
	}

	// query multiple row from table
	{
		// to query multiple row from table, we can declare struct and append it to an array.
		type User struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		// to execute command, use db exec
		_, errIns := db.Exec(`INSERT INTO users (username, password, created_at) values (?,?,?);`, "Alice", "Alice123", time.Now())

		if errIns != nil {
			log.Fatal(err)
		}

		// to retrieve data, use query
		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close() // make sure to close row after all operation in current block finished

		var users []User

		// to loop row, use rows.Next
		for rows.Next() {
			var u User
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)

			if err != nil {
				log.Fatal(err)
			}

			users = append(users, u)
		}

		// check error if failed query in iteration
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(users)
	}

	// delete user
	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
