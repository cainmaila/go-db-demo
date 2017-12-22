package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:0080@tcp(192.168.99.100:3306)/cain")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// rows, err := db.Query("select id, name, age from tb where name = ?", "Cain")
	rows, err := db.Query("select id, name, age from tb")
	if err != nil {
		log.Fatal(err)
	}
	var id, age int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//單筆資料
	// var name string
	// db.QueryRow("select name from tb where id = 1").Scan(&name)
	// fmt.Println(name)
}
