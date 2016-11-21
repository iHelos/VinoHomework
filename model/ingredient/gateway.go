package ingredient

import (
	"fmt"
	"log"
)

type ingredientGateway struct {
	id int
	inDB bool
	Name string
	Description string
}

func NewIngredient(name, description string) *ingredientGateway{
	return ingredientGateway{
		name: name,
		description:description,
	}
}

func (gate *ingredientGateway) Insert(){
	q := fmt.Sprintf("insert into %s (%s,%s) values ($1,$2)", IngredientTable, I_name, I_description)
	rows, err := db.Exec(q, "asdsadasd", "вкусняшка")
	log.Print(err)
	log.Print(rows)
}

func (gate *ingredientGateway) Update(){
	q := fmt.Sprintf("UPDATE %s SET %s = $1, %s = $2 WHERE %s = $5;", IngredientTable, I_name, I_description, I_ID)
	rows, err := db.Exec(q, "asdsadasd", "asdasasdasd", 1)
	log.Print(err)
	log.Print(rows)
}

func (gate *ingredientGateway) Remove(){

}
