package dal

import (
	"database/sql"
	"e-shop-management-system/internal/config"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(DBconfig config.DBConfig) (*Store, error) {
	conn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", DBconfig.DBconnection, DBconfig.DBuser, DBconfig.DBpassword, DBconfig.DBhost, DBconfig.DBport, DBconfig.DBname)
	db, err := sql.Open(DBconfig.DBconnection, conn)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}