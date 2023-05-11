package main

import (
	"context"
	"database/sql"
	"github.com/suportebeloj/desafio-hitss/src/cmd/settings"
	"log"
)

var (
	ctx    context.Context = context.Background()
	dbConn *sql.DB
)

func main() {
	conn, err := settings.NewPostgresConnection()
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()

	dbConn = conn

}
