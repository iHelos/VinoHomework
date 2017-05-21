package model

import (
	. "github.com/iHelos/VinoHomework/model"
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
)

var db *sql.DB

func Start(pool *sql.DB) {
	db = pool
	query := fmt.Sprintf("Create table if not exists %s (%s serial PRIMARY KEY, %s text)", LOGSTable, LOGS_ID, LOGS_log)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
