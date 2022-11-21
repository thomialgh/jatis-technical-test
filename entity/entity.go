package entity

import "time"

type Customers struct {
	CustomerID          int64  `db:"CustomerID"`
	CompanyName         string `db:"CompanyName"`
	FirstName           string `db:"FirstName"`
	LastName            string `db:"LastName"`
	BillingAddress      string `db:"BillingAddress"`
	City                string `db:"City"`
	StateOfProvince     string `db:"StateOfProvince"`
	ZIPCode             string `db:"ZIPCode"`
	Email               string `db:"Email"`
	CompanyWebsite      string `db:"CompanyWebsite"`
	PhoneNumber         string `db:"PhoneNumber"`
	FaxNumber           string `db:"FaxNumber"`
	ShipAddress         string `db:"ShipAddress"`
	ShipCity            string `db:"ShipCity"`
	ShipStateOrProvince string `db:"ShipStateOrProvince"`
	ShipZIPCode         string `db:"ShipZIPCode"`
	ShipPhoneNumber     string `db:"ShipPhoneNumber"`
}

type Employees struct {
	EmployeeID int64  `db:"EmployeeID"`
	FirstName  string `db:"FirstName"`
	LastName   string `db:"LastName"`
	Title      string `db:"Title"`
	WorkPhone  string `db:"WorkPhone"`
}

type ShippingMethods struct {
	ShippingMethodID int64  `db:"ShippingMethodID"`
	ShippingMethod   string `db:"ShippingMethod"`
}

type Orders struct {
	OrderID             int64     `db:"OrderID"`
	CustomerID          int64     `db:"CustomerID"`
	EmployeeID          int64     `db:"EmployeeID"`
	OrderDate           time.Time `db:"OrderDate"`
	PurchaseOrderNumber string    `db:"PurchaseOrderNumber"`
	ShipDate            time.Time `db:"ShipDate"`
	ShippingMethodID    int64     `db:"ShippingMethodID"`
	FreightCharge       float64   `db:"FreightCharge"`
	Taxes               float64   `db:"Taxes"`
	PaymentReceived     string    `db:"PaymentReceived"`
	Comment             string    `db:"Comment"`
}

type Products struct {
	ProductID   int64   `db:"ProductID"`
	ProductName string  `db:"ProductName"`
	UnitPrice   float64 `db:"UnitPrice"`
	InStock     string  `db:"InStock"`
}

type OrderDetails struct {
	OrderDetailID int64   `db:"OrderDetailID"`
	OrderID       int64   `db:"OrderID"`
	ProductID     int64   `db:"ProductID"`
	Quantity      int64   `db:"Quantity"`
	UnitPrice     float64 `db:"UnitPrice"`
	Discount      float64 `db:"Discount"`
}
