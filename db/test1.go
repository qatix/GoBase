package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
)

func main()  {

	db,err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/testdb")

	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("connection successfully")
	}

	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("ping ok")
	}

	var(
		id int
		name string
		username string
	)
	rows,err := db.Query("select user_id as id,name,username from user where user_id=?",1)
	if err != nil{
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&id,&name,&username)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(id,username,name)
	}
	err = rows.Err()
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close();
}
