package link_DI

import (
	"database/sql"
	"fmt"
	"log"
	. "github.com/iHelos/VinoHomework/model"
)

var db *sql.DB

func Start(db *sql.DB){
	db = db
	query := fmt.Sprintf("Create table if not exists %s (%s serial Primary key, %s integer, %s integer, FOREIGN KEY (%s) REFERENCES %s (%s), FOREIGN KEY (%s) REFERENCES %s (%s))",
		DITable,
		DI_ID,
		DI_dish_ID,
		DI_ingredient_ID,
		DI_dish_ID,
		DishTable,
		D_ID,
		DI_ingredient_ID,
		IngredientTable,
		I_ID,
	)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}