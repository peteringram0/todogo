package models

import (
	"database/sql"
	"todogo/structs"
)

func CreateUser(db *sql.DB, user structs.User) (int64, error) {

	sql := `INSERT INTO users(email, name, picture) VALUES(?,?,?)`

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(user.Email, user.Name, user.Picture)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()

}

func GetUser(db *sql.DB, email string) (string, error) {

	sql := `SELECT id FROM users WHERE email=?`

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	var id string
	err = stmt.QueryRow(email).Scan(&id)

	// Exit if we get an error
	if err != nil {
		// log.Fatal(err)
		// log.Println(err)
		return "", err
	}

	return id, nil

}
