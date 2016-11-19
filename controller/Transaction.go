package controller

import (
	"database/sql"
)

type BusinessTransaction struct {
	connection_pool *sql.DB
}

func (*BusinessTransaction) startDataBase(){

}