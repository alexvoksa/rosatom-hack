package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"

	"github.com/alexvoksa/rosatom-hack/go/internal/processor"
)

type Repository interface {
	UpsertProducts(new *processor.Product) error
	UpsertCustomers(new *processor.Customer, orderAmount float64) error
	UpsertSuppliers(new *processor.Suppliers, orderPrice float64) error
	UpsertTenders(new *processor.Tender) error
}

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

const upsertSupplierContactStmt = `
INSERT INTO suppliers_contact (first_name, middle_name, last_name, email,phone, address) 
values ($1,$2, $3, $4, $5, $6) ON CONFLICT DO NOTHING;
`

func (r *repo) UpsertSuppliers(new *processor.Suppliers, orderPrice float64) error {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString(`INSERT INTO suppliers
	(ogrn, name, short_name, inn, kpp, registered_at, sold_amount) values `)

	for i := range new.Supplier {
		sqlBuilder.WriteString(formatSupplierForSuppliersTable(&new.Supplier[i], orderPrice))
		if i != len(new.Supplier)-1 {
			sqlBuilder.WriteString(", ")
		}
	}

	sqlBuilder.WriteString(`ON CONFLICT (ogrn) DO NOTHING;`)
	//UPDATE SET
	//successfull_tenders = suppliers.successfull_tenders + 1,
	//unsuccessfull_tenders = suppliers.unsuccessfull_tenders + 0;`

	_, err := r.db.Exec(context.Background(), sqlBuilder.String())
	if err != nil {
		return err
	}

	sqlBuilder.Reset()

	sqlBuilder.WriteString(`INSERT INTO supplier_contacts 
	(supplier_id, first_name, middle_name, last_name, email,phone, address) values `)
	for i := range new.Supplier {
		info := &new.Supplier[i].LegalEntity.OtherInfo
		sqlBuilder.WriteString(fmt.Sprintf("(%d, '%s', '%s', '%s', '%s', '%s', '%s')",
			new.Supplier[i].LegalEntity.EGRULInfo.OGRN,
			info.ContactInfo.FirstName, info.ContactInfo.MiddleName, info.ContactInfo.LastName,
			info.ContactEmail, info.ContactPhone, new.Supplier[i].LegalEntity.EGRULInfo.Address))
		if i != len(new.Supplier)-1 {
			sqlBuilder.WriteString(", ")
		}
	}

	sqlBuilder.WriteString("ON CONFLICT(id) DO NOTHING")

	_, err = r.db.Exec(context.Background(), sqlBuilder.String())

	return err
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

func formatSupplierForSuppliersTable(sup *processor.Supplier, price float64) string {
	egrulInfo := sup.LegalEntity.EGRULInfo

	_ = "(ogrn, name, short_name, inn, kpp, registered_at, sold_amount, okpo, oktmo_name, oktmo_code) values `)"

	return fmt.Sprintf("(%d, '%s', '%s', '%s', '%s', '%s', %f, '%s', '%s', '%s')",
		egrulInfo.OGRN, egrulInfo.FullName, egrulInfo.ShortName, egrulInfo.INN, egrulInfo.KPP, egrulInfo.RegistrationDate, price,
	)
}
