package main

import (
	"os"
	"encoding/csv"
	"fmt"
	"strconv"
)

type CsvLine struct {
	StoreId   int32
	Name      string
	Qty       float32
	Type      string
	DateAdded string
}

func main() {
	filename := "order_service_data_after_20170101.csv"

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		data := CsvLine{
			StoreId:   string2Int32(line[0]) ,
			Name:      line[1],
			Qty:       string2Float32(line[2]),
			Type:      line[3],
			DateAdded: line[4],
		}
		fmt.Println(data)
	}
}

func string2Int32(s string) int32  {
	i,err := strconv.ParseInt(s,10,64)
	if err != nil{
		panic(err)
	}
	return int32(i)
}

func string2Float32(s string) float32  {
	f,err := strconv.ParseFloat(s,64)
	if err != nil{
		panic(err)
	}
	return float32(f)
}