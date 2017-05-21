package kitchen

import (
	"database/sql"
	"fmt"
	. "github.com/iHelos/VinoHomework/model"
	"log"
)

var db *sql.DB

func Start(pool *sql.DB) {
	db = pool
	query := fmt.Sprintf("Create table if not exists %s (%s serial PRIMARY KEY, %s varchar(40) UNIQUE, %s text)", KitchenTable, K_ID, K_name, K_description)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
