package ingredient

import (
	"database/sql"
	"errors"
	"fmt"
	. "github.com/iHelos/VinoHomework/model"
	"github.com/labstack/gommon/log"
)

type ingredientGateway struct {
	ID          int
	inDB        bool
	Name        string
	Description string
}

func NewIngredient_Local(name, description string) *ingredientGateway {
	return &ingredientGateway{
		ID:          0,
		Name:        name,
		Description: description,
		inDB:        false,
	}
}

func NewIngredient_DB() *ingredientGateway {
	return &ingredientGateway{
		inDB: true,
	}
}

func (gate *ingredientGateway) Insert() error {
	if gate.inDB == true {
		return errors.New("already inserted")
	}
	q := fmt.Sprintf("insert into %s (%s,%s) values ($1,$2) RETURNING %s", IngredientTable, I_name, I_description, I_ID)
	res, err := db.Query(q, gate.Name, gate.Description)
	if err != nil {
		return err
	}
	defer res.Close()
	res.Next()
	err = res.Scan(&gate.ID)
	return err
}

func (gate *ingredientGateway) Update() error {
	if gate.inDB == false {
		return errors.New("object not created")
	}
	q := fmt.Sprintf("UPDATE %s SET %s = $1, %s = $2 WHERE %s = $3;", IngredientTable, I_name, I_description, I_ID)
	_, err := db.Exec(q, gate.Name, gate.Description, gate.ID)
	return err
}

func makeQuery(tx *sql.Tx, query string, paramerer interface{}) error {
	//tx, err := db.Begin()
	_, err := tx.Exec(query, paramerer)
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (gate *ingredientGateway) Remove() error {
	if gate.inDB == false {
		return errors.New("object not created")
	}

	q1 := fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", DITable, DI_dish_ID)
	q2 := fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", IngredientTable, I_ID)
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	err = makeQuery(tx, q1, gate.ID)
	if err != nil {
		return err
	}
	log.Print(gate.ID)
	err = makeQuery(tx, q2, gate.ID)
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}
