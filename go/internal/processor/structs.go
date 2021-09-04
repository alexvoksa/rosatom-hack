package processor

import (
	"encoding/xml"
	"time"
)

type Tender struct {
	XMLName  xml.Name `xml:"export"`
	Contract Contract `xml:"contract"`
}

type Contract struct {
	ID          string    `xml:"id"`
	RegNum      string    `xml:"regNum"`
	Number      string    `xml:"number"`
	PublishDate string    `xml:"publishDate"`
	Customer    Customer  `xml:"customer"`
	Currency    Currency  `xml:"currency"`
	Products    Products  `xml:"products"`
	Suppliers   Suppliers `xml:"suppliersInfo"`
	Href        string    `xml:"href"`
}

type Currency struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}

type Customer struct {
	RegNumber string `xml:"regNum"`
	FullName  string `xml:"fullName"`
	INN       string `xml:"inn"`
	KPP       string `xml:"kpp"`
}

type Products struct {
	Product []Product `xml:"product"`
}

type Suppliers struct {
	Supplier []Supplier `xml:"supplierInfo"`
}

type Product struct {
	SID           int64   `xml:"sid"`
	GUID          string  `xml:"guid"`
	OKPD          OKPD    `xml:"OKPD2"`
	Name          string  `xml:"name"`
	HierarchyType string  `xml:"hierarchyType"`
	Type          string  `xml:"type"`
	OKEI          OKEI    `xml:"OKEI"`
	Price         float64 `xml:"price"`
	PriceRUR      float64 `xml:"priceRUR"`
	Quantity      float64 `xml:"quantity"`
	Sum           float64 `xml:"sum"`
	SumRUR        float64 `xml:"sumRUR"`
	VATRate       string  `xml:"VATRate"`
}
type OKPD struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}

type OKEI struct {
	Code             string `xml:"code"`
	NationalCode     string `xml:"nationalCode"`
	TrueNationalCode string `xml:"trueNationalCode"`
	FullName         string `xml:"fullName"`
	NationalName     string
}

type Supplier struct {
	LegalEntity LegalEntity `xml:"legalEntityRF"`
}

type LegalEntity struct {
	EGRULInfo EGRULInfo `xml:"EGRULInfo"`
	OtherInfo OtherInfo `xml:"otherInfo"`
}

type EGRULInfo struct {
	OGRN             int64     `xml:"OGRN"`
	LegalForm        LegalForm `xml:"legalForm"`
	FullName         string    `xml:"fullName"`
	ShortName        string    `xml:"shortName"`
	FirmName         string    `xml:"firmName"`
	INN              string    `xml:"INN"`
	KPP              string    `xml:"KPP"`
	RegistrationDate string    `xml:"registrationDate"`
	// после анмаршаллинга заполнить это поле, чтобы не ловить ошибки во время анмаршалинга
	registrationDate time.Time
	Address          string `xml:"address"`
}

type LegalForm struct {
	Code         int64  `xml:"code"`
	SingularName string `xml:"singularName"`
}

type OtherInfo struct {
	Status       int64  `xml:"status"`
	OKPO         int64  `xml:"OKPO"`
	OKTMO        OKTMO  `xml:"OKTMO"`
	ContactPhone string `xml:"contactPhone"`
	ContactEmail string `xml:"contactEMail"`
}

type OKTMO struct {
	Code int64  `xml:"code"`
	Name string `xml:"name"`
}
