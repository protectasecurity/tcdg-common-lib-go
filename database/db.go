package database

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
)

func ConnectDB(dbparams map[string]interface{}) (*sql.DB, error) {
	connectionString := dbparams["engine"].(string) + "://" + dbparams["username"].(string) + ":" + dbparams["password"].(string) + "@" + dbparams["host"].(string) + ":" + dbparams["port"].(string) + "/" + dbparams["dbname"].(string)

	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
