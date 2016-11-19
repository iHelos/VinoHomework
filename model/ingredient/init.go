package ingredient

import "database/sql"

var connection *sql.DB

func setConnectionPool(pool *sql.DB){
	connection = pool
}