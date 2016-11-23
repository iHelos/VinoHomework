package ingredient

import (
	"fmt"
	. "github.com/iHelos/VinoHomework/model"

	"errors"
)

func Find(id int) (*ingredientGateway, error){
	q := fmt.Sprintf("select %s,%s,%s from %s where %s = $1",
		I_ID, I_name, I_description, 	//select
		IngredientTable,
		//DishTable, CategoryTable,			//join tables
		//D_Category, C_ID,				//join condition
		I_ID,						//select condition
	)
	row, err := db.Query(q,id)
	if err!=nil{
		return nil,err
	}
	defer row.Close()
	ok := row.Next()
	if !ok{
		return nil,  errors.New("Not Found")
	}
	obj := NewIngredient_DB()
	err = row.Scan(&obj.ID, &obj.Name, &obj.Description)
	return obj, err
}