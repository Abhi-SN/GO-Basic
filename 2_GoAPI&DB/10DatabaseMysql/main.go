package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // need to add this line manuallay
)

type User struct { // User struct represents a row in the `users` table
	ID    int
	Name  string
	email string
}

func main() {
	db := DB_connection()                  // starts Database connection
	err := initializeDatabaseAndTables(db) // Creates Database And Tables
	if err != nil {
		log.Fatal("Error initialize Database And Tables", err)
	}
	defer db.Close()

	// Inserts Values into Users tables
	err = InsertValues(db)
	if err != nil {
		log.Fatal("Not able to insert values ", err)
	}

	// Updates Values into Users tables
	err = UpdateValues(db)
	if err != nil {
		log.Fatal("Not able to update values ", err)
	}

	// Prints Values From Users tables
	UserInfo(db)

	// Delete values from User table
	DeleteUser(db)

}

func DB_connection() *sql.DB {
	dsn := "root:1256@tcp(127.0.0.1:3306)/GOTESTDB"
	db, err := sql.Open("mysql", dsn) // Initializes the DB driver(drivers can mysql, postgres ..etc ) with config (does not open a connection immediately).

	if err != nil {
		log.Fatal("Failed connection to Databases", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("MySql is Not reachbale:", err)
	}
	fmt.Println("Connected to MYSQL Server")
	return db
}

func initializeDatabaseAndTables(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS GOTESTDB")
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}
	// SelectDB := `USE GOTESTDB`
	// db.Exec(SelectDB)

	createTableQuery := `CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("unable create table: %v", err)
	}
	return nil

}

func InsertValues(db *sql.DB) error {
	insertUsers := "INSERT IGNORE INTO users (name, email) VALUES (?, ?)"
	res, err := db.Exec(insertUsers, "admin", "admin@gmail.com")
	if err != nil {
		return fmt.Errorf("not able insert values :- %v", err)
	}
	id, _ := res.LastInsertId() //LastInsertId() returns: - The last auto-generated ID (usually the id column) from the most recent INSERT statement on that DB connection.
	if id == 0 {
		fmt.Println("The User already exist")
	} else {
		fmt.Println("The Inserted User Id is :- ", id)
	}
	return nil
}

func UpdateValues(db *sql.DB) error {
	updateUser := "UPDATE USERS SET email = ? WHERE id = ?"
	_, err := db.Exec(updateUser, "abd@gamil.com.in", 7)
	if err != nil {
		return fmt.Errorf("not able update values :- %v", err)
	}
	fmt.Println("Updated the user's email")
	return nil
}

func UserInfo(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("Failed to fetch table data:", err)
	}
	defer rows.Close() //Close closes the Rows, preventing further enumeration. If Next is called and returns false and there are no further result sets
	// fmt.Println("ID		| Name		|Email		")

	for rows.Next() { //Next prepares the next result row for reading with the Scan method. It returns true on success, or false
		var userData User
		err := rows.Scan(&userData.ID, &userData.Name, &userData.email) //Scan copies the columns in the current row into the values pointed at by dest.
		if err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		fmt.Printf("ID  %d | Name  %v | Email %v \n", userData.ID, userData.Name, userData.email)
	}
}

func DeleteUser(db *sql.DB) {
	id := 2
	deleteUser := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(deleteUser, id)
	if err != nil {
		log.Fatal(" Failed to delete:", err)
	}

	fmt.Printf("Deleted user with ID: %d\n", id)
}
