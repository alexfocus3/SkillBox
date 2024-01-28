//Напишите простейшую программу для добавления данных в таблицу БД PostgreSQL

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=postgres password=123 dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO students(s_id, name, start_year) VALUES (1477, 'Bob', 2017),(1478, 'Alice', 2018);")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}
