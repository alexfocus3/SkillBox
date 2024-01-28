//Напишите программу, демонстрирующую простейший запрос к БД PostgreSQL и вывод результатов запроса.

package main

import (
	"database/sql"
	"fmt"

	//"unicode/utf8"

	_ "github.com/lib/pq"
)

type student struct {
	s_id       int
	name       string
	start_year int
}

func main() {

	connStr := "user=postgres password=123 dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from students")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	students := []student{}

	for rows.Next() {
		p := student{}
		err := rows.Scan(&p.s_id, &p.name, &p.start_year)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//r, _ := utf8.EncodeRune([]byte(p.name))
		//fmt.Println(r)
		students = append(students, p)
	}
	for _, p := range students {
		fmt.Println(p.s_id, p.name, p.start_year)
	}
}
