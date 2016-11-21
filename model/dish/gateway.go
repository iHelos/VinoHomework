package dish

import (
	"fmt"
	"log"
)

type DishGateway struct {
	id int
	inDB bool
	Name string
	Description string
	Price int
	Category int
}

func (gate *DishGateway) Insert(){
	q := fmt.Sprintf("insert into %s (%s,%s,%s,%s) values ($1,$2,$3,$4)", DishTable, D_Name, D_Description, D_Price, D_Category)
	rows, err := db.Exec(q, "asdsadasd", "вкусняшка", 15, 7)
	log.Print(err)
	log.Print(rows)
}

func (gate *DishGateway) Update(){
	q := fmt.Sprintf("UPDATE %s SET %s = $1, %s = $2, %s = $3, %s = $4 WHERE %s = $5;", DishTable, D_Name, D_Description, D_Price, D_Category, D_ID)
	rows, err := db.Exec(q, "asdsadasd", "вкусняшка", 15, 7, 1)
	log.Print(err)
	log.Print(rows)
}

func (gate *DishGateway) Remove(){

}