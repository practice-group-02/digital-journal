package dal

import (
	"database/sql"
	"digital-journal/internal/config"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewStore(DBconfig config.DBConfig) error {
	var err error
	conn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", DBconfig.DBconnection, DBconfig.DBuser, DBconfig.DBpassword, DBconfig.DBhost, DBconfig.DBport, DBconfig.DBname)
	DB, err = sql.Open(DBconfig.DBconnection, conn)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
