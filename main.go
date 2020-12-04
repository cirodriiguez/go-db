package main

import (
	"github.com/cirodriiguez/go-db/service"
	"github.com/cirodriiguez/go-db/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := service.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	storageInvoice := storage.NewPsqlInvoice(storage.Pool())
	serviceInvoice := service.NewService(storageInvoice)
	if err := serviceInvoice.Migrate(); err != nil {
		log.Fatalf("invoice.Migrate: %v", err)
	}

	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := service.NewService(storageInvoiceItem)
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.Migrate: %v", err)
	}
}
