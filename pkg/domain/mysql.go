package domain

import (
	"fmt"
	"jatis/entity"
	"strings"

	"github.com/jmoiron/sqlx"
)

func InsertBulkCustomersTx(tx *sqlx.Tx, d []entity.Customers) error {
	baseQuery := `
		INSERT INTO
			Jatis.Customers
		(
			CustomerID,
			CompanyName,
			FirstName,
			LastName,
			BillingAddress,
			City,
			StateOfProvince,
			ZIPCode,
			Email,
			CompanyWebsite,
			PhoneNumber,
			FaxNumber,
			ShipAddress,
			ShipCity,
			ShipStateOrProvince,
			ShipZIPCode,
			ShipPhoneNumber
		)
		VALUES
		%s
	`

	arrParams := make([]interface{}, 0, len(d))
	paramQueryArr := make([]string, 0, len(d))

	for _, v := range d {
		arrParams = append(arrParams, []interface{}{
			v.CustomerID,
			v.CompanyName,
			v.FirstName,
			v.LastName,
			v.BillingAddress,
			v.City,
			v.StateOfProvince,
			v.ZIPCode,
			v.Email,
			v.CompanyWebsite,
			v.PhoneNumber,
			v.FaxNumber,
			v.ShipAddress,
			v.ShipCity,
			v.ShipStateOrProvince,
			v.ShipZIPCode,
			v.ShipPhoneNumber,
		}...)
		paramQueryArr = append(paramQueryArr, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	}

	baseQuery = fmt.Sprintf(baseQuery, strings.Join(paramQueryArr, ",\n"))

	_, err := tx.Exec(baseQuery, arrParams...)
	if err != nil {
		return err
	}

	return nil
}

func InsertBulkEmployeesTx(tx *sqlx.Tx, d []entity.Employees) error {
	baseQuery := `
		INSERT INTO
			Jatis.Employees
		(
			EmployeeID,
			FirstName,
			LastName,
			Title,
			WorkPhone
		)
		VALUES
		%s
	`

	arrParams := make([]interface{}, 0, len(d))
	paramQueryArr := make([]string, 0, len(d))

	for _, v := range d {
		arrParams = append(arrParams, []interface{}{
			v.EmployeeID,
			v.FirstName,
			v.LastName,
			v.Title,
			v.WorkPhone,
		}...)
		paramQueryArr = append(paramQueryArr, "(?, ?, ?, ?, ?)")
	}

	baseQuery = fmt.Sprintf(baseQuery, strings.Join(paramQueryArr, ",\n"))

	_, err := tx.Exec(baseQuery, arrParams...)
	if err != nil {
		return err
	}

	return nil
}

func InsertShippingMethodsTx(tx *sqlx.Tx, d []entity.ShippingMethods) error {
	baseQuery := `
		INSERT INTO
			Jatis.Shipping_Methods
		(
			ShippingMethodID,
			ShippingMethod
		)
		VALUES
		%s
	`

	arrParams := make([]interface{}, 0, len(d))
	paramQueryArr := make([]string, 0, len(d))

	for _, v := range d {
		arrParams = append(arrParams, []interface{}{
			v.ShippingMethodID,
			v.ShippingMethod,
		}...)
		paramQueryArr = append(paramQueryArr, "(?, ?)")
	}

	baseQuery = fmt.Sprintf(baseQuery, strings.Join(paramQueryArr, ",\n"))

	_, err := tx.Exec(baseQuery, arrParams...)
	if err != nil {
		return err
	}

	return nil
}

func InsertOrdersTx(tx *sqlx.Tx, d []entity.Orders) error {
	baseQuery := `
		INSERT INTO
			Jatis.Orders
		(
			OrderID,
			CustomerID,
			EmployeeID,
			OrderDate,
			PurchaseOrderNumber,
			ShipDate,
			ShippingMethodID,
			FreightCharge,
			Taxes,
			PaymentReceived,
			Comment
		)
		VALUES
		%s
	`

	arrParams := make([]interface{}, 0, len(d))
	paramQueryArr := make([]string, 0, len(d))

	for _, v := range d {
		arrParams = append(arrParams, []interface{}{
			v.OrderID,
			v.CustomerID,
			v.EmployeeID,
			v.OrderDate,
			v.PurchaseOrderNumber,
			v.ShipDate,
			v.ShippingMethodID,
			v.FreightCharge,
			v.Taxes,
			v.PaymentReceived,
			v.Comment,
		}...)
		paramQueryArr = append(paramQueryArr, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	}

	baseQuery = fmt.Sprintf(baseQuery, strings.Join(paramQueryArr, ",\n"))

	_, err := tx.Exec(baseQuery, arrParams...)
	if err != nil {
		return err
	}

	return nil
}

func InsertProductsTx(tx *sqlx.Tx, d []entity.Products) error {
	baseQuery := `
		INSERT INTO
			Jatis.Products
		(
			ProductID,
			ProductName,
			UnitPrice,
			InStock
		)
		VALUES
		%s
	`

	arrParams := make([]interface{}, 0, len(d))
	paramQueryArr := make([]string, 0, len(d))

	for _, v := range d {
		arrParams = append(arrParams, []interface{}{
			v.ProductID,
			v.ProductName,
			v.UnitPrice,
			v.InStock,
		}...)
		paramQueryArr = append(paramQueryArr, "(?, ?, ?, ?)")
	}

	baseQuery = fmt.Sprintf(baseQuery, strings.Join(paramQueryArr, ",\n"))

	_, err := tx.Exec(baseQuery, arrParams...)
	if err != nil {
		return err
	}

	return nil
}

func InsertOrderDetailsTx(tx *sqlx.Tx, d []entity.OrderDetails) error {
	baseQuery := `
		INSERT INTO
			Jatis.Order_Details
		(
			OrderDetailID,
			OrderID,
			ProductID,
			Quantity,
			UnitPrice,
			Discount
		)
		VALUES
		%s
	`

	arrParams := make([]interface{}, 0, len(d))
	paramQueryArr := make([]string, 0, len(d))

	for _, v := range d {
		arrParams = append(arrParams, []interface{}{
			v.OrderDetailID,
			v.OrderID,
			v.ProductID,
			v.Quantity,
			v.UnitPrice,
			v.Discount,
		}...)
		paramQueryArr = append(paramQueryArr, "(?, ?, ?, ?, ?, ?)")
	}

	baseQuery = fmt.Sprintf(baseQuery, strings.Join(paramQueryArr, ",\n"))

	_, err := tx.Exec(baseQuery, arrParams...)
	if err != nil {
		return err
	}

	return nil
}

func GetOrder(db *sqlx.DB, orderID int64) (*entity.Orders, error) {
	query := `
		SELECT
			OrderID
			,CustomerID
			,EmployeeID
			,OrderDate
			,PurchaseOrderNumber
			,ShipDate
			,ShippingMethodID
			,FreightCharge
			,Taxes
			,PaymentReceived
			,Comment
		FROM Jatis.Orders
		WHERE OrderID = ?
	`

	var order entity.Orders

	err := db.QueryRowx(query, orderID).StructScan(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func GetCustomer(db *sqlx.DB, CustomeriD int64) (*entity.Customers, error) {
	var cust entity.Customers
	err := db.QueryRowx(`
		SELECT
			CustomerID
			,CompanyName
			,FirstName
			,LastName
			,BillingAddress
			,City
			,StateOfProvince
			,ZIPCode
			,Email
			,CompanyWebsite
			,PhoneNumber
			,FaxNumber
			,ShipAddress
			,ShipCity
			,ShipStateOrProvince
			,ShipZIPCode
			,ShipPhoneNumber
		FROM Jatis.Customers
		WHERE CustomerID = ?
	`, CustomeriD).StructScan(&cust)
	if err != nil {
		return nil, err
	}

	return &cust, nil
}

func GetEmployee(db *sqlx.DB, employeeID int64) (*entity.Employees, error) {
	var empl entity.Employees
	err := db.QueryRowx(`
		SELECT
			EmployeeID
			,FirstName
			,LastName
			,Title
			,WorkPhone
		FROM Jatis.Employees
		WHERE EmployeeID = ?
	`, employeeID).StructScan(&empl)
	if err != nil {
		return nil, err
	}

	return &empl, nil
}

func GetShippingMethods(db *sqlx.DB, ShippingMethodID int64) (*entity.ShippingMethods, error) {
	var shippingMethods entity.ShippingMethods
	err := db.QueryRowx(`
		SELECT
			ShippingMethodID
			,ShippingMethod
		FROM Jatis.Shipping_Methods
		WHERE ShippingMethodID = ?
	`, ShippingMethodID).StructScan(&shippingMethods)

	if err != nil {
		return nil, err
	}

	return &shippingMethods, nil
}

func GetOrderDetailsByOrderID(db *sqlx.DB, OrderID int64) ([]entity.OrderDetails, []int64, error) {
	var orderDetails []entity.OrderDetails
	rows, err := db.Queryx(`
		SELECT 
			OrderDetailID
			,OrderID
			,ProductID
			,Quantity
			,UnitPrice
			,Discount
		FROM Jatis.Order_Details
		WHERE OrderID = ?
	`, OrderID)

	if err != nil {
		return nil, nil, err
	}

	var productIDMap = make(map[int64]bool)
	var productIDs []int64

	for rows.Next() {
		var orderDetail entity.OrderDetails
		err = rows.StructScan(&orderDetail)
		if err != nil {
			return nil, nil, err
		}

		if _, ok := productIDMap[orderDetail.ProductID]; !ok {
			productIDs = append(productIDs, orderDetail.ProductID)
			productIDMap[orderDetail.ProductID] = true
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	return orderDetails, productIDs, nil
}

func GetProductsMap(db *sqlx.DB, productIds []int64) (map[int64]entity.Products, error) {
	productMap := make(map[int64]entity.Products)

	query := `
		SELECT
			ProductID
			,ProductName
			,UnitPrice
			,InStock
		FROM Jatis.Products
		WHERE ProductID in (%s)
	`
	args := make([]interface{}, 0, len(productIds))
	inMark := make([]string, 0, len(productIds))
	for _, v := range productIds {
		inMark = append(inMark, "?")
		args = append(args, v)
	}

	query = fmt.Sprintf(query, strings.Join(inMark, ", "))

	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product entity.Products
		err = rows.StructScan(&product)
		if err != nil {
			return nil, err
		}

		productMap[product.ProductID] = product
	}

	return productMap, nil

}
