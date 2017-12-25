package main

import (
	"database/sql"
	"fmt"
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

	//新增資料

	stmt, err := db.Prepare("insert into tb (name,age) values (?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Tony3", 20)
	if err != nil {
		log.Fatal(err)
	}
	addKey, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("新增了 %d\n", addKey)

	//刪除
	stmt, err = db.Prepare("DELETE FROM tb WHERE name=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err = stmt.Exec("Tony3")
	if err != nil {
		log.Fatal(err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("刪除了 %d 筆\n", num)

	//查詢
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

	//查詢單筆資料
	// var name string
	// db.QueryRow("select name from tb where id = 1").Scan(&name)
	// fmt.Println(name)
}
