package main

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
)

func main() {

	db, err := sql.Open("oracle", "oracle://STARHUBRA:RSPRD_SH15@10.90.70.155:1525/SHRASPRD")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}
