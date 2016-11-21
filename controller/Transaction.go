package controller

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/iHelos/VinoHomework/model/ingredient"
	"github.com/iHelos/VinoHomework/model/dish"
	"github.com/iHelos/VinoHomework/model/kitchen"

	"log"
)

type BusinessTransaction struct {
	connection_pool *sql.DB
}

func (*BusinessTransaction) Start() (ok bool){
	db, err := sql.Open("postgres", "postgresql://admin:keks@139.59.133.90:5432/KitchenDB")
	if err != nil {
		log.Fatal(err)
	}

	ingredient.Start(db)
	kitchen.Start(db)
	dish.Start(db)
	return true
}

func (*BusinessTransaction) CreateDish() (ok bool){
	asd := dish.DishGateway{}
	asd.Insert()

	lol := ingredient.IngredientGateway{}
	lol.Insert()

	return true
}