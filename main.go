package main

import (
	"log"
	"github.com/kataras/iris"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=admin dbname=KitchenDB")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Create table if not exists ", age)

	log.Print(rows)
	iris.Listen("127.0.0.1:8080")
}
