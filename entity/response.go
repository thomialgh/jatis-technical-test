package entity

import "time"

type OrdersResponse struct {
	OrderID             int64                   `json:"OrderID"`
	Customers           CustomerResponse        `json:"Customers"`
	Employees           EmployeeResponse        `json:"Employees"`
	OrderDate           time.Time               `json:"OrderDate"`
	PurchaseOrderNumber string                  `json:"PurchaseOrderNumber"`
	ShipDate            time.Time               `json:"ShipDate"`
	ShippingMethod      ShippingMethodsResponse `json:"ShippingMethod"`
	FreightCharge       float64                 `json:"FreightCharge"`
	Taxes               float64                 `json:"Taxes"`
	PaymentReceived     string                  `json:"PaymentReceived"`
	Comment             string                  `json:"Comment"`
	OrderDetails        []OrderDetailResponse   `json:"OrderDetails"`
	TotalPayment        float64                 `json:"TotalPayment"`
}

type CustomerResponse struct {
	CustomerID          int64  `json:"CustomerID"`
	CompanyName         string `json:"CompanyName"`
	FirstName           string `json:"FirstName"`
	LastName            string `json:"LastName"`
	BillingAddress      string `json:"BillingAddress"`
	City                string `json:"City"`
	StateOfProvince     string `json:"StateOfProvince"`
	ZIPCode             string `json:"ZIPCode"`
	Email               string `json:"Email"`
	CompanyWebsite      string `json:"CompanyWebsite"`
	PhoneNumber         string `json:"PhoneNumber"`
	FaxNumber           string `json:"FaxNumber"`
	ShipAddress         string `json:"ShipAddress"`
	ShipCity            string `json:"ShipCity"`
	ShipStateOrProvince string `json:"ShipStateOrProvince"`
	ShipZIPCode         string `json:"ShipZIPCode"`
	ShipPhoneNumber     string `json:"ShipPhoneNumber"`
}

type EmployeeResponse struct {
	EmployeeID int64  `json:"EmployeeID"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Title      string `json:"Title"`
	WorkPhone  string `json:"WorkPhone"`
}

type ShippingMethodsResponse struct {
	ShippingMethodID int64  `json:"ShippingMethodID"`
	ShippingMethod   string `json:"ShippingMethod"`
}

type OrderDetailResponse struct {
	OrderDetailID int64           `json:"OrderDetailID"`
	OrderID       int64           `json:"-"`
	ProductID     int64           `json:"-"`
	Quantity      int64           `json:"Quantity"`
	UnitPrice     float64         `json:"UnitPrice"`
	Discount      float64         `json:"Discount"`
	Products      ProductResponse `json:"products"`
}

type ProductResponse struct {
	ProductID   int64   `json:"ProductID"`
	ProductName string  `json:"ProductName"`
	UnitPrice   float64 `json:"UnitPrice"`
	InStock     string  `json:"InStock"`
}
