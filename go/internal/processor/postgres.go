package processor

import (
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresProcessor struct {
	workerID   int
	inputChan  <-chan *Tender
	wg         *sync.WaitGroup
	postgresDB *pgxpool.Conn
}

func (p *postgresProcessor) startProcessing() {
	var err error
	fmt.Println("waiting for tenders")
	for tender := range p.inputChan {
		fmt.Println("received tender for update", tender.ID)
		err = p.updatePostgres(tender)
		if err != nil {
			fmt.Printf("dbWorker#%d failed to update tender %s: %v", p.workerID, tender.ID, err)
			continue
		}
	}

	p.wg.Done()
}

func (p *postgresProcessor) updatePostgres(message *Tender) error {
	var err error
	err = p.updateSuppliers(message.Suppliers)
	err = p.updateTenders(message)
	err = p.updateCustomers(&message.Customer)

	return err
}

func (p *postgresProcessor) updateSuppliers(sup Suppliers) error {
	fmt.Println("updating suppliers")
	return nil
}

func (p *postgresProcessor) updateTenders(contract *Tender) error {
	fmt.Println("updating tenders")

	return nil
}

func (p *postgresProcessor) updateCustomers(cust *Customer) error {
	fmt.Println("updating customers")
	return nil
}
