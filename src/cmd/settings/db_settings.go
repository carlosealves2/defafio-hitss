package settings

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

	"fmt"
	"os"
	"strconv"
)

type ICredential interface {
	GetDSN() string
}

type PostgresCredential struct {
	host, user, pass, dbName string
	port                     int
}

var (
	DbConn *sql.DB
)

func init() {
	conn, err := NewPostgresConnection()
	if err != nil {
		log.Fatalln(err)
	}

	DbConn = conn
}

// function to collect database crendentials from environment
// and return a DSN string to connection
func (p *PostgresCredential) GetDSN() string {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl != "" {
		return databaseUrl
	}

	p.host = os.Getenv("DB_HOST")

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	if dbPort != 0 {
		p.port = dbPort
	} else {
		p.port = 5432
	}

	p.user = os.Getenv("DB_USER")
	p.pass = os.Getenv("DB_PASS")
	p.dbName = os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", p.host, p.port, p.dbName, p.user, p.pass)
}

func NewPostgresConnection() (*sql.DB, error) {
	postgres := PostgresCredential{}
	conn, err := sql.Open("postgres", postgres.GetDSN())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}
