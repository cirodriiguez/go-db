package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoice = `CREATE TABLE IF NOT EXISTS invoices(
		id SERIAL NOT NULL,
		client VARCHAR(25) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_id_pk PRIMARY KEY (id)
	)`
)

type PsqlInvoice struct {
	db *sql.DB
}

func NewPsqlInvoice(db *sql.DB) *PsqlInvoice {
	return &PsqlInvoice{db}
}

func (p *PsqlInvoice) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoice)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de factura ejecutada con exito")
	return nil
}
