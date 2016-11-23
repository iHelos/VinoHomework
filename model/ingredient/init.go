package ingredient

import (
	"database/sql"
	"log"
	"fmt"
	. "github.com/iHelos/VinoHomework/model"
)

var db *sql.DB

func Start(pool *sql.DB){
	db = pool
	query := fmt.Sprintf("Create table if not exists %s (%s serial PRIMARY KEY, %s varchar(40) UNIQUE, %s text)", IngredientTable, I_ID, I_name, I_description)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}