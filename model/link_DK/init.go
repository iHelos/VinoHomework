package link_DK


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
		DKTable,
		//атрибуты
		DK_ID,
		DK_dish_ID,
		DK_kitchen_ID,
		//ссылка раз
		DK_dish_ID,
		DishTable,
		D_ID,
		//ссылка два
		DK_kitchen_ID,
		KitchenTable,
		K_ID,
	)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}