package models

import (
	"database/sql"
	"todogo/helper"
	"todogo/structs"

	"github.com/labstack/echo"
)

func GetTasks(db *sql.DB, c echo.Context) structs.TaskCollection {
	sql := "SELECT * FROM tasks WHERE uid=?"

	id, errOne := helper.GetUID(c)

	if errOne != nil {
		return structs.TaskCollection{}
	}

	rows, err := db.Query(sql, id)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := structs.TaskCollection{}
	for rows.Next() {
		task := structs.Task{}
		err2 := rows.Scan(&task.ID, &task.UID, &task.Name)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

func PostTask(db *sql.DB, name string, uid int64) (int64, error) {
	sql := "INSERT INTO tasks(name,uid) VALUES(?,?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name, uid)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()

}

func PutTask(db *sql.DB, id string, name string) (int64, error) {
	sql := `UPDATE tasks SET name=? WHERE id=?`

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name, id)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()

}
