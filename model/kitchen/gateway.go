package kitchen

import (
	"fmt"
	"database/sql"
	. "github.com/iHelos/VinoHomework/model"

	"github.com/kataras/go-errors"
)

type kitchenGateway struct {
	ID int
	inDB bool
	Name string
	Description string
}

func NewKitchen_Local(name, description string) *kitchenGateway{
	return &kitchenGateway{
		ID:0,
		Name: name,
		Description:description,
		inDB:false,
	}
}


func NewKitchen_DB() *kitchenGateway{
	return &kitchenGateway{
		inDB:true,
	}
}

func (gate *kitchenGateway) Insert() error{
	if gate.inDB == true{
		return errors.New("already inserted")
	}
	q := fmt.Sprintf("insert into %s (%s,%s) values ($1,$2) RETURNING %s", KitchenTable, K_name, K_description, K_ID)
	res, err := db.Query(q, gate.Name, gate.Description)
	if err!=nil{
		return err
	}
	defer res.Close()
	res.Next()
	err = res.Scan(&gate.ID)
	return err
}

func (gate *kitchenGateway) Update() error{
	if gate.inDB == false{
		return errors.New("object not created")
	}
	q := fmt.Sprintf("UPDATE %s SET %s = $1, %s = $2 WHERE %s = $3;", KitchenTable, K_name, K_description, K_ID)
	_, err := db.Exec(q, gate.Name, gate.Description, gate.ID)
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

func (gate *kitchenGateway) Remove() error{
	if gate.inDB == false{
		return errors.New("object not created")
	}
	q1 := fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", DKTable, DK_dish_ID)
	q2 := fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", KitchenTable, K_ID)
	tx, err := db.Begin()
	err = makeQuery(tx, q1, gate.ID)
	if err!= nil{
		return err
	}
	err = makeQuery(tx, q2, gate.ID)
	if err!= nil{
		return err
	}
	tx.Commit()
	return err
}
