package dish
import (
	"database/sql"
	"log"
	"fmt"
	"github.com/iHelos/VinoHomework/model/kitchen"
	"github.com/iHelos/VinoHomework/model/ingredient"
)

var db *sql.DB

func Start(pool *sql.DB){
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
	query = fmt.Sprintf("Create table if not exists %s (%s serial Primary key, %s integer, %s integer, FOREIGN KEY (%s) REFERENCES %s (%s), FOREIGN KEY (%s) REFERENCES %s (%s))",
		DKTable,
		DK_ID,
		DK_dish_ID,
		DK_kitchen_ID,
		DK_dish_ID,
		DishTable,
		D_ID,
		DK_kitchen_ID,
		kitchen.KitchenTable,
		kitchen.K_ID,
	)
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	query = fmt.Sprintf("Create table if not exists %s (%s serial Primary key, %s integer, %s integer, FOREIGN KEY (%s) REFERENCES %s (%s), FOREIGN KEY (%s) REFERENCES %s (%s))",
		DITable,
		DI_ID,
		DI_dish_ID,
		DI_ingredient_ID,
		DI_dish_ID,
		DishTable,
		D_ID,
		DI_ingredient_ID,
		ingredient.IngredientTable,
		ingredient.I_ID,
	)
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}