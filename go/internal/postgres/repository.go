package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"

	"github.com/alexvoksa/rosatom-hack/go/internal/processor"
)

var _ Repository = &repo{}

type repo struct {
	db     *pgx.Conn
	logger *zap.SugaredLogger
}

const upsertProductsStmt = `INSERT INTO products (guid, okpd_code, okei_id, name, taxes, price, amount)
values ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT DO NOTHING;`

func (r *repo) UpsertProducts(new *processor.Product) error {
	resp, err := r.db.Exec(context.Background(), upsertProductsStmt,
		new.GUID, new.OKPD.Code, new.OKEI.Code, new.Name, new.VATRate, new.PriceRUR, new.Quantity)
	if err != nil {
		return err
	}

	if resp.RowsAffected() != 1 {
		r.logger.Warnf("for product %s num of affected rows is: %d", new.GUID, resp.RowsAffected())
	}

	return nil
}

const upsertCustomersExec = `INSERT INTO customers (reg_num, name, inn, kpp, registered_at,
okpo, ordered_amount, successfull_tenders, unsuccessfull_tenders)
values ($1,$2,$3,$4,$5,$6,$7,1, 0) ON CONFLICT (reg_num)
DO UPDATE SET
successfull_tenders = customers.successfull_tenders + 1,
unsuccessfull_tenders = customers.unsuccessfull_tenders + 0,
ordered_amount = customers.ordered_amount + $8;`

func (r *repo) UpsertCustomers(new *processor.Customer, orderAmount float64) error {
	resp, err := r.db.Exec(context.Background(), upsertCustomersExec,
		new.RegNumber, new.FullName, new.INN, new.KPP, new.RegistrationDate, new.OKPO, orderAmount)
	if err != nil {
		return err
	}

	if resp.RowsAffected() != 1 {
		r.logger.Warnf("for customer %s affected row is: %d", new.RegNumber, resp.RowsAffected())
	}

	return nil
}

const upsertTendersStmt = `
INSERT INTO tenders (id, reg_num, price, published_at, url, customer, name, tender_products, resolution)
values ($1,$2, $3,$4,$5,$6, 'aa', 'aaa', 1) ON CONFLICT DO NOTHING RETURNING id;`

func (r *repo) UpsertTenders(new *processor.Tender) error {
	resp, err := r.db.Exec(context.Background(), upsertTendersStmt,
		new.ID, new.RegNumber, new.PriceInfo.Price, new.PublishDate, new.Href, new.Customer.RegNumber)
	if err != nil {
		return err
	}

	if resp.RowsAffected() != 1 {
		r.logger.Warnf("for tender %s affected row is: %d", new.RegNumber, resp.RowsAffected())
	}

	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString(fmt.Sprintf("INSERT INTO tender_products (tender_id, products_id) VALUES "))
	for i := range new.Products.Product {
		product := &new.Products.Product[i]
		sqlBuilder.WriteString(fmt.Sprintf("('%s', '%s')", new.ID, product.GUID))
		if i != len(new.Products.Product)-1 {
			sqlBuilder.WriteString(", ")
		}
	}

	sqlBuilder.WriteString(";")

	_, err = r.db.Exec(context.Background(), sqlBuilder.String())

	return err
}

const upsertSuppliersStmt = `
INSERT INTO suppliers (ogrn, name, short_name, email,phone, address, inn, kpp, registered_at, reg_num, classification,
description, reputation, sold_amount, successful_tenders, unsuccessful_tenders, is_innovate) 
values ($1,$2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT (ogrn)
DO UPDATE SET
successfull_tenders = suppliers.successfull_tenders + ?,
unsuccessfull_tenders = suppliers.unsuccessfull_tenders + ?;`

func (r *repo) UpsertSuppliers(new *processor.Suppliers) error {
	//	sqlBuilder := strings.Builder{}
	//	sqlBuilder.WriteString(`INSERT INTO suppliers
	//(ogrn, name, short_name, email,phone, address, inn, kpp, registered_at, reg_num, classification,
	//description, reputation, sold_amount, successful_tenders, unsuccessful_tenders, is_innovate)
	//values `)
	//
	//	for i := range new.Supplier {
	//
	//
	//	}
	//
	//
	//
	//
	//	resp, err := r.db.Exec(context.Background(), upsertSuppliersStmt,
	//		new., new.RegNumber, new.PriceInfo.Price, new.PublishedAt, new.Href)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if resp.RowsAffected() != 1 {
	//		r.logger.Warnf("for tender %s affected row is: %d", new.RegNumber, resp.RowsAffected())
	//	}
	//
	//	sqlBuilder = strings.Builder{}
	//	sqlBuilder.WriteString(fmt.Sprintf("INSERT INTO tender_products (tender_id, products_id) VALUES "))
	//	for i := range new.Products.Product {
	//		product := &new.Products.Product[i]
	//		sqlBuilder.WriteString(fmt.Sprintf("(%s, %s)", new.ID, product.GUID))
	//		if i != len(new.Products.Product)-1 {
	//			sqlBuilder.WriteString(", ")
	//		}
	//	}
	//
	//	sqlBuilder.WriteString(";")
	//
	//	_, err = r.db.Exec(context.Background(), sqlBuilder.String(),
	//		new.ID, new.RegNumber, new.PriceInfo.Price, new.PublishedAt, new.Href)
	//
	//	return err

	return nil
}

func generateParamNums(from, to int) string {
	builder := strings.Builder{}
	builder.WriteByte('(')
	for i := from; i <= to; i++ {
		builder.WriteString(fmt.Sprintf("$%d ", i))
	}

	builder.WriteByte(')')

	return builder.String()
}
