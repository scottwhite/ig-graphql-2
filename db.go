package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sqlx.DB
}

func NewDB() (*DB, error) {
	host := viper.GetString("db.host")
	port := viper.GetInt("db.port")
	user := viper.GetString("db.username")
	password := viper.GetString("db.password")
	dbname := viper.GetString("db.database")

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user, password, host, port, dbname)

	log.Println("db conn", psqlInfo)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	return &DB{db}, nil
}
