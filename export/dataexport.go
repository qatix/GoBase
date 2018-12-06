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

type ServiceItem struct {
	store_id      int32
	name          string
	quantity      float64
	business_type int32
	date_added    string
}

var serviceItems []ServiceItem
var bmap = map[int32]string{
	1:   "洗车",
	2:   "美容",
	3:   "保养",
	4:   "机电",
	5:   "钣金",
	6:   "理赔",
	7:   "改装",
	8:   "轮胎",
	9:   "喷漆",
	10:  "代办",
	20:  "用品",
	50:  "批发",
	100: "其它",
}

func exportData(dbconfig DbConfig, storeId int32) {
	fmt.Printf("start to export data :%d\n",storeId)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbconfig.user, dbconfig.passwd, dbconfig.host, dbconfig.name)
	fmt.Println(dsn)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connection successfully")
	}

	//var(
	//	date_added string
	//	name string
	//	business_type int
	//	quantity float64
	//)
	qsql := "select s.date_added,sd.name,sd.business_type,sd.quantity from sales_detail sd left join sales s ON(s.sales_id=sd.sales_id) where sd.ref_type=1 and s.status=1 and s.date_added >= '2017-01-01 00:00:00'"

	rows, err := db.Query(qsql)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		//var item ServiceItem
		item := ServiceItem{}
		err := rows.Scan(&item.date_added, &item.name, &item.business_type, &item.quantity)
		if err != nil {
			log.Fatal(err)
		}
		item.store_id = storeId
		//fmt.Println(item)
		serviceItems = append(serviceItems, item)
	}

	db.Close()
}

func exportDataToCsv(data []ServiceItem, filename string) {
	file, err := os.Create(filename)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range data {
		business_type_str := "unknow"
		if val,ok := bmap[item.business_type];ok{
			business_type_str = val
		}

		line := []string{
			fmt.Sprintf("%d", item.store_id),
			item.name,
			fmt.Sprintf("%.2f", item.quantity),
			business_type_str,
			//fmt.Sprintf("%d", item.business_type),
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

	dev_mode := 1

	dbconfig := DbConfig{
		"127.0.0.1",
		"root",
		"123456",
		"ccadmin",
	}

	if dev_mode != 1 {
		dbconfig = DbConfig{
			"rds2ynnny2ynnny.mysql.rds.aliyuncs.com",
			"ccwkadmin",
			"Buluting1802",
			"ccwk_admin_v1",
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

	var (
		store_id int32
		dbname   string
		user     string
		passwd   string
		host     string
	)
	rows, err := db.Query("SELECT s.store_id,d.db_name as dbname,du.db_username as user,du.db_password as passwd,h.address as `host` FROM db d LEFT JOIN db_user du ON(d.db_user_id=du.db_user_id) LEFT JOIN host h ON(h.host_id=du.db_host_id) left join store s ON(d.store_id=s.store_id) where s.system_level=3 order by s.store_id asc")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&store_id, &dbname, &user, &passwd, &host)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(store_id, dbname, user, passwd, host)

		if dev_mode == 1 && store_id != 10008 {
			continue
		}

		dbconfig2 := DbConfig{host, user, passwd, dbname}
		exportData(dbconfig2, store_id)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(serviceItems)
	fmt.Println("rows:" , len(serviceItems))
	exportDataToCsv(serviceItems, "order_service_data_after_20170101.csv")
	fmt.Println("export data done")

	defer db.Close();
}
