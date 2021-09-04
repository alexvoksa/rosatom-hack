package postgres

import (
	"github.com/alexvoksa/rosatom-hack/go/internal/processor"
)

type Repository interface {
	UpsertProducts(new *processor.Product) error
	UpsertCustomers(new *processor.Customer, orderAmount float64) error
	UpsertTenders(new *processor.Tender) error
	UpsertSuppliers(new *processor.Suppliers) error
}
