package processor

import (
	"context"
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
	for tender := range p.inputChan {
		err = p.updatePostgres(tender)
		if err != nil {
			fmt.Printf("dbWorker#%d failed to update tender %s: %v", p.workerID, tender.Contract.ID, err)
			continue
		}
	}

	p.wg.Done()
}

func (p *postgresProcessor) updatePostgres(message *Tender) error {
	var err error
	err = p.updateSuppliers()
	err = p.updateTenders()
	err = p.updateCustomers()

	return err
}

func (p *postgresProcessor) updateSuppliers() error {
	return nil
}

func (p *postgresProcessor) updateTenders() error {
	p.postgresDB.Exec(context.Background(), "INSERT ")

	return nil
}

func (p *postgresProcessor) updateCustomers() error {
	return nil
}
