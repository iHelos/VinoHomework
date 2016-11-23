package kitchen

import (
	"fmt"
	. "github.com/iHelos/VinoHomework/model"
	"errors"
)

func Find(id int) (*kitchenGateway, error){
	q := fmt.Sprintf("select %s,%s,%s from %s where %s = $1",
		K_ID, K_name, K_description, 	//select
		KitchenTable,
		//DishTable, CategoryTable,			//join tables
		//D_Category, C_ID,				//join condition
		K_ID,						//select condition
	)
	row, err := db.Query(q,id)
	if err!=nil{
		return nil,err
	}
	defer row.Close()
	ok := row.Next()
	if !ok{
		return nil, errors.New("Not Found")
	}
	kitchen_obj := NewKitchen_DB()
	err = row.Scan(&kitchen_obj.ID, &kitchen_obj.Name, &kitchen_obj.Description)
	return kitchen_obj, err
}