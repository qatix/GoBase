package main

import (
	"database/sql"
	"fmt"
	"log"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		fmt.Println("errr")
	}
	defer db.Close()

	var (
		id   int
		name string
	)
	rows, err := db.Query("select user_id as id, name from user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
