package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	db *sql.DB
	once sync.Once
)

type Storage interface {
	Migrate() error
	/*Create(*Model)	error
	Update(*Model)	error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error*/
}

func NewPostgresDB()  {
	once.Do(func() {
		var err error
		connStr := "postgres://godev:godev123@192.168.10.10:5432/godb?sslmode=disable"
		db, err = sql.Open("postgres", connStr)

		if err != nil {
			log.Fatalf("no se logro conectar con la db: %v", err)
		}

		if err = db.Ping(); err != nil{
			log.Fatalf("no se logro realizar el ping a la db: %v", err)
		}

		fmt.Println("Conectado")
	})
}

func Pool() *sql.DB {
	return db
}