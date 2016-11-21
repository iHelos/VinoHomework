package kitchen

import (
	"fmt"
	"log"
)

type KitchenGateway struct {
	Name string
	Description string
}

func (gate *KitchenGateway) Insert(){
	q := fmt.Sprintf("insert into %s (%s,%s) values ($1,$2)", KitchenTable, K_name, K_description)
	rows, err := db.Exec(q, "asdsadasd", "вкусняшка")
	log.Print(err)
	log.Print(rows)
}

func (gate *KitchenGateway) Update(){

}

func (gate *KitchenGateway) Remove(){

}

