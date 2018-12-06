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


	rows,err := db.Query("select user_id as id,name,username,mobile,gender from user")
	if err != nil{
		log.Fatal(err)
	}

	defer rows.Close()
	cols,err := rows.Columns()
	if err != nil{
		fmt.Println("handle error")
	}else{
		dest := []interface{}{
			new(uint64),
			new(string),
			new(string),
			new(string),
			new(string),
		}
		fmt.Println("len cols:",len(cols))

		for rows.Next() {
			err = rows.Scan(dest...)
			fmt.Println(dest)
			//how to access field ???
		}
	}


	defer db.Close();
}

