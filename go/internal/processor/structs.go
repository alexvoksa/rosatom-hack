package processor

import (
	"encoding/xml"
	"time"
)

type XmlFile struct {
	XMLName xml.Name `xml:"export"`
	Tender  Tender   `xml:"contract"`
}

type Tender struct {
	ID          string `xml:"id"`
	RegNumber   string `xml:"regNum"`
	Number      string `xml:"number"`
	PublishDate string `xml:"publishDate"`
	PublishedAt time.Time
	Customer    Customer  `xml:"customer"`
	Products    Products  `xml:"products"`
	Suppliers   Suppliers `xml:"suppliersInfo"`
	Href        string    `xml:"href"`
	PriceInfo   struct {
		Price float64 `xml:"priceRUR"`
		Taxes float64 `xml:"priceVATRUR"`
	} `xml:"priceInfo"`
}

type Customer struct {
	RegNumber        string `xml:"regNum"`
	FullName         string `xml:"fullName"`
	INN              string `xml:"inn"`
	KPP              string `xml:"kpp"`
	RegistrationDate string `xml:"registrationDate"`
	// заполняем после парсинга
	RegisteredAt time.Time
	OKPO         string `xml:"OKPO"`
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
	VatRUR        float64
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
	NationalName     string `xml:"nationalName"`
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
	RegisteredAt time.Time
	Address      string `xml:"address"`
}

type LegalForm struct {
	Code         int64  `xml:"code"`
	SingularName string `xml:"singularName"`
}

type ContactInfo struct {
	LastName   string `xml:"lastName"`
	FirstName  string `xml:"firstName"`
	MiddleName string `xml:"middleName"`
}

type OtherInfo struct {
	Status       int64       `xml:"status"`
	OKPO         string      `xml:"OKPO"`
	OKTMO        OKTMO       `xml:"OKTMO"`
	ContactPhone string      `xml:"contactPhone"`
	ContactEmail string      `xml:"contactEMail"`
	ContactInfo  ContactInfo `xml:"contactInfo"`
}

type OKTMO struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}
