package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
	"os"
	"encoding/csv"
)

type DbConfig struct {
	host   string
	user   string
	passwd string
	name   string
}

type Store struct {
	company_name   string
	manager_name   string
	manager_mobile string
	address        sql.NullString
	date_added     string
}

var storeList []Store

func exportDataToCsv(data []Store, filename string) {
	file, err := os.Create(filename)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	headLine := []string{
		"店名",
		"负责人姓名",
		"负责人电话",
		"地址",
		"注册时间",
	}
	err2 := writer.Write(headLine)
	checkError("Cannot write head to csv", err2)

	for _, item := range data {
		line := []string{
			item.company_name,
			item.manager_name,
			item.manager_mobile,
			item.address.String,
			item.date_added,
		}
		err := writer.Write(line)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func main() {

	dev_mode := 0

	dbconfig := DbConfig{
		"127.0.0.1",
		"root",
		"123456",
		"cckd",
	}

	if dev_mode != 1 {
		dbconfig = DbConfig{
			"rdsvjfrfvvjfrfv.mysql.rds.aliyuncs.com",
			"cckd",
			"z3c28D69vO1Fa0l",
			"cckd2",
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbconfig.user, dbconfig.passwd, dbconfig.host, dbconfig.name)
	fmt.Println(dsn)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connection successfully")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ping ok")
	}

	rows, err := db.Query("select company_name,manager_name,manager_mobile,address,date_added from store")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		item := Store{}
		err := rows.Scan(&item.company_name, &item.manager_name, &item.manager_mobile, &item.address, &item.date_added)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(item)

		storeList = append(storeList, item)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("rows:", len(storeList))
	exportDataToCsv(storeList, "cckd_store_list.csv")
	fmt.Println("export data done")

	defer db.Close();
}
