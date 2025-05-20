package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Users struct {
	Id    int
	Name  string
	Email string
}

type Message struct {
	Messages string `json:"Message"`
}

var db *sql.DB

func main() {
	db = DB_connection()
	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/users", GetAllUser).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/newusers", CreateNewUsers).Methods("POST")
	r.HandleFunc("/updateUsers/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/deleteUser/{id}", deleteUser).Methods("DELETE")

	// Start server
	log.Println("Server running on http://localhost:3080")
	log.Fatal(http.ListenAndServe(":3080", r))

}

func DB_connection() *sql.DB {
	dsn := "root:1256@tcp(127.0.0.1:3306)/GOTESTDB"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Unable to connect to DB")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to connect to DB server", err)
	}
	fmt.Println("Connected to MYSQL Server")
	return db

}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("unable to fetch data from DB", err)
	}
	defer rows.Close()
	var users []Users

	for rows.Next() {
		var u Users
		err := rows.Scan(&u.Id, &u.Name, &u.Email)
		if err != nil {
			log.Fatal("Failed to read the data", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func CreateNewUsers(w http.ResponseWriter, r *http.Request) {
	var createUsers Users
	err := json.NewDecoder(r.Body).Decode(&createUsers)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	result, err := db.Exec("INSERT INTO users(name, email) VALUES(?,?)", createUsers.Name, createUsers.Email)
	if err != nil {
		http.Error(w, "Database Insert Faied", http.StatusInternalServerError)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to fetch the ID", http.StatusInternalServerError)
		return
	}
	createUsers.Id = int(id)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(createUsers)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser Users
	vars := mux.Vars(r)
	userId := vars["id"]
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	result, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", updateUser.Name, updateUser.Email, userId)
	if err != nil {
		http.Error(w, "Failed to update DB", http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected() //RowsAffected returns the number of rows affected by an update, insert, or delete.
	if err != nil || rowsAffected == 0 {
		http.Error(w, "No user found to update", http.StatusNotFound)
		return
	}
	// updateUser.Id = stringToInt(id)
	idInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid ID ", http.StatusBadRequest)
		return
	}
	updateUser.Id = idInt
	response := Message{
		Messages: "User updated successfully",
	}
	json.NewEncoder(w).Encode(response)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usrId := vars["id"]
	var getUser Users
	err := json.NewDecoder(r.Body).Decode(&getUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}
	rows, err := db.Query("SELECT * FROM users WHERE id = ?", usrId)
	if err != nil {
		log.Fatal("unable to fetch data from DB", err)
	}
	defer rows.Close()
	var users []Users

	for rows.Next() {
		var u Users
		err := rows.Scan(&u.Id, &u.Name, &u.Email)
		if err != nil {
			log.Fatal("Failed to read the data", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var deleteuser Users
	err := json.NewDecoder(r.Body).Decode(&deleteuser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Not able to delete", http.StatusInternalServerError)
		return
	}
	rowsAfft, err := result.RowsAffected()
	if err != nil || rowsAfft == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	resp := "User Updated successfully for id " + id

	response := Message{
		Messages: resp,
	}
	json.NewEncoder(w).Encode(response)
}
