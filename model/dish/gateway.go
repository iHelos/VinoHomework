package dish

import (
	"fmt"
	"errors"
	"database/sql"
	. "github.com/iHelos/VinoHomework/model"
	"github.com/labstack/gommon/log"
)

type dishGateway struct {
	ID int
	inDB bool
	Name string
	Description string
	Price int
	Category int
	Ingredient []string
	Kitchen []string
}

func NewDish() *dishGateway{
	return &dishGateway{
		ID:0,
		Name: "",
		Description:"",
		inDB:false,
	}
}

func NewDish_local(name, description string) *dishGateway{
	return &dishGateway{
		ID:0,
		Name: name,
		Description:description,
		inDB:false,
	}
}


func NewDish_DB() *dishGateway{
	return &dishGateway{
		inDB:true,
	}
}

func (gate *dishGateway) Insert() error{
	if gate.inDB == true{
		return errors.New("already inserted")
	}
	q := fmt.Sprintf("insert into %s (%s,%s,%s,%s) values ($1,$2,$3,$4) RETURNING %s", DishTable, D_Name, D_Description, D_Price, D_Category, D_ID)
	res, err := db.Query(q, gate.Name, gate.Description, gate.Price, gate.Category)
	if err!=nil{
		return err
	}
	defer res.Close()
	res.Next()
	err = res.Scan(&gate.ID)

	q_kitchen := fmt.Sprintf("insert into %s (%s,%s) values ($1,$2)", DKTable, DK_dish_ID, DK_kitchen_ID)
	for _, element := range gate.Kitchen {
		_, err = db.Exec(q_kitchen, gate.ID, element)
		if err!=nil{
			log.Print(err)
		}
	}

	q_ingredient := fmt.Sprintf("insert into %s (%s,%s) values ($1,$2)", DITable, DI_dish_ID, DI_ingredient_ID)
	for _, element := range gate.Ingredient {
		_, err = db.Exec(q_ingredient, gate.ID, element)
		if err!=nil{
			log.Print(err)
		}
	}

	return err
}

func (gate *dishGateway) Update() error{
	if gate.inDB == false{
		return errors.New("object not created")
	}
	q := fmt.Sprintf("UPDATE %s SET %s = $1, %s = $2, %s = $3, %s = $4 WHERE %s = $5;", DishTable, D_Name, D_Description, D_Price, D_Category, D_ID)
	_, err := db.Exec(q, gate.Name, gate.Description, gate.Price, gate.Category, gate.ID)
	return err
}

func makeQuery(tx *sql.Tx, query string, paramerer interface{}) error{
	_, err := tx.Exec(query, paramerer)
	if err!= nil{
		tx.Rollback()
		return err
	}
	return nil
}

func (gate *dishGateway) Remove() error{
	if gate.inDB == false{
		return errors.New("object not created")
	}
	q1 := fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", DITable, DI_dish_ID)
	q2 := fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", DKTable, DK_dish_ID)
	q3 := fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", DishTable, D_ID)

	tx, err := db.Begin()
	err = makeQuery(tx, q1, gate.ID)
	if err!= nil{
		return err
	}
	err = makeQuery(tx, q2, gate.ID)
	if err!= nil{
		return err
	}
	err = makeQuery(tx, q3, gate.ID)
	if err!= nil{
		return err
	}
	tx.Commit()
	return err
}

func (gate *dishGateway) AddIngredient() error{
	if gate.inDB == false{
		return errors.New("object not created")
	}
	q := fmt.Sprintf("insert into %s (%s,%s,%s,%s) values ($1,$2,$3,$4)", DishTable, D_Name, D_Description, D_Price, D_Category)
	_, err := db.Exec(q, gate.Name, gate.Description, gate.Price, gate.Category)
	return err
}