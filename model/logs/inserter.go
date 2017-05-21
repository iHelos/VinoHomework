package model

import (
	. "github.com/iHelos/VinoHomework/model"
	"fmt"
)

type logGateway struct {
	text string
}

func NewDBLog(text string) *logGateway {
	return &logGateway{text}
}

func (gate *logGateway) Insert() {
	q := fmt.Sprintf("insert into %s (%s) values ($1)", LOGSTable, LOGS_log)
	_, err := db.Exec(q, gate.text)
	if(err != nil){
		fmt.Println(err)
	}
}
