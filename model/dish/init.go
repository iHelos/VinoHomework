package dish

import (
	"database/sql"
	"fmt"
	. "github.com/iHelos/VinoHomework/model"
	"log"
)

var db *sql.DB

func Start(pool *sql.DB) {
	db = pool
	query := fmt.Sprintf("Create table if not exists %s (%s serial PRIMARY KEY, %s varchar(40) UNIQUE)", CategoryTable, C_ID, C_Name)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	query = fmt.Sprintf("Create table if not exists %s (%s serial PRIMARY KEY, %s varchar(40) UNIQUE, %s text, %s integer, %s integer, FOREIGN KEY (%s) REFERENCES %s (%s))",
		DishTable,
		D_ID,
		D_Name,
		D_Description,
		D_Price,
		D_Category,
		D_Category,
		CategoryTable,
		C_ID,
	)
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
