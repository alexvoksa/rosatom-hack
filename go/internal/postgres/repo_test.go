package postgres

import (
	"context"
	"encoding/xml"
	"io/ioutil"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"

	"github.com/alexvoksa/rosatom-hack/go/internal/processor"
)

const dateLayout = "2006-07-21"

var logger = zap.NewNop().Sugar()

func TestRepo_UpsertTenders(t *testing.T) {
	db, err := pgx.Connect(context.Background(), "postgres://hackathon:hackathon@localhost:5432/hackathon")
	if err != nil {
		t.Fatal(err)
	}

	r := repo{db: db, logger: logger}

	xmlFile, err := readData("/Users/anantonenko/GolandProjects/staff/rosatom-hack/data/contract_1010502039621000006_66410533.xml")
	if err != nil {
		t.Fatal(err)
	}

	err = r.UpsertTenders(&xmlFile.Tender)
	if err != nil {
		t.Fatal(err)
	}

}

func TestRepo_UpsertSuppliers(t *testing.T) {
	db, err := pgx.Connect(context.Background(), "postgres://hackathon:hackathon@localhost:5432/hackathon")
	if err != nil {
		t.Fatal(err)
	}

	r := repo{db: db, logger: logger}

	xmlFile, err := readData("/Users/anantonenko/GolandProjects/staff/rosatom-hack/data/contract_1010502039621000006_66410533.xml")
	if err != nil {
		t.Fatal(err)
	}

	err = r.UpsertSuppliers(&xmlFile.Tender.Suppliers, xmlFile.Tender.PriceInfo.Price)
	if err != nil {
		t.Fatal(err)
	}

}

func readData(path string) (*processor.XmlFile, error) {
	var xmlFile processor.XmlFile
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, &xmlFile)
	if err != nil {
		return nil, err
	}

	xmlFile.Tender.PublishedAt, _ = time.Parse(dateLayout, xmlFile.Tender.PublishDate)

	for i := range xmlFile.Tender.Products.Product {
		product := &xmlFile.Tender.Products.Product[i]
		// если не смогли распарсить, то всё равно 0 будет
		product.VatRUR, _ = strconv.ParseFloat(product.VATRate, 10)
	}

	return &xmlFile, nil
}
