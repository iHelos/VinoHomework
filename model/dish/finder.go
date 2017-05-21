package dish

import (
	"errors"
	"fmt"
	. "github.com/iHelos/VinoHomework/model"
)

func Find(id int) (*dishGateway, error) {
	q := fmt.Sprintf("select %s,%s,%s,%s,%s from %s where %s = $1",
		D_ID, D_Name, D_Description, D_Price, D_Category, //select
		DishTable,
		//DishTable, CategoryTable,			//join tables
		//D_Category, C_ID,				//join condition
		D_ID, //select condition
	)
	row, err := db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	ok := row.Next()
	if !ok {
		return nil, errors.New("Not Found")
	}
	dish_obj := NewDish_DB()
	err = row.Scan(&dish_obj.ID, &dish_obj.Name, &dish_obj.Description, &dish_obj.Price, &dish_obj.Category)
	return dish_obj, err
}
