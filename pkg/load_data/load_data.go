package load_data

import (
	"encoding/csv"
	"fmt"
	"jatis/entity"
	"os"
	"path"
	"strconv"
	"time"

	"jatis/pkg/domain"
	"jatis/pkg/repo"
)

func InsertDB(dirName string) error {
	var customers []entity.Customers
	var employees []entity.Employees
	var shippingMethods []entity.ShippingMethods
	var orders []entity.Orders
	var products []entity.Products
	var orderDetails []entity.OrderDetails

	dir, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}

	// Read all file in directory
	for _, v := range dir {
		if v.IsDir() {
			continue // continue if directory has another directory
		}

		// readFile
		filePath := path.Join(dirName, v.Name())
		f, err := os.Open(filePath)
		if err != nil {
			return err
		}

		defer f.Close()

		reader := csv.NewReader(f)
		records, err := reader.ReadAll()
		if err != nil {
			return err
		}

		// if file has zero record, return error
		if len(records) == 0 {
			return fmt.Errorf("File : %s with zero record", v.Name())
		}

		headerArr := records[0] // get first row to get all header name

		headerMap := make(map[string]int)

		for i, v := range headerArr {
			headerMap[v] = i
		}

		for _, record := range records[1:] {
			// handle data based on filename
			switch v.Name() {
			case "customers.csv":
				customer, err := parseCustomerRecord(headerMap, record)
				if err != nil {
					return err
				}

				customers = append(customers, customer)
			case "employees.csv":
				employee, err := parseEmployesRecrod(headerMap, record)
				if err != nil {
					return err
				}

				employees = append(employees, employee)

			case "shipping_methods.csv":
				shippingMethod, err := parseShippingMethdosRecords(headerMap, record)
				if err != nil {
					return err
				}

				shippingMethods = append(shippingMethods, shippingMethod)

			case "orders.csv":
				order, err := parseOrdersRecord(headerMap, record)
				if err != nil {
					return err
				}

				orders = append(orders, order)

			case "products.csv":
				product, err := parseProductsRecord(headerMap, record)
				if err != nil {
					return err
				}

				products = append(products, product)

			case "order_details.csv":
				orderDetail, err := parseOrderDetailsRecord(headerMap, record)
				if err != nil {
					return err
				}

				orderDetails = append(orderDetails, orderDetail)
			}

		}

	}

	// using transaction to store data
	tx, err := repo.DB.Beginx()
	if err != nil {
		return err
	}

	err = domain.InsertBulkCustomersTx(tx, customers)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = domain.InsertBulkEmployeesTx(tx, employees)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = domain.InsertShippingMethodsTx(tx, shippingMethods)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = domain.InsertOrdersTx(tx, orders)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = domain.InsertProductsTx(tx, products)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = domain.InsertOrderDetailsTx(tx, orderDetails)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func parseCustomerRecord(headerMap map[string]int, record []string) (entity.Customers, error) {
	var res entity.Customers
	var err error

	res.CustomerID, err = strconv.ParseInt(record[headerMap["CustomerID"]], 10, 64)
	if err != nil {
		return res, err
	}

	res.CompanyName = record[headerMap["CompanyName"]]
	res.FirstName = record[headerMap["FileName"]]
	res.LastName = record[headerMap["LastName"]]
	res.BillingAddress = record[headerMap["BillingAddress"]]
	res.City = record[headerMap["City"]]
	res.StateOfProvince = record[headerMap["StateOfProvince"]]
	res.ZIPCode = record[headerMap["ZIPCode"]]
	res.Email = record[headerMap["Email"]]
	res.CompanyWebsite = record[headerMap["CompanyWebsite"]]
	res.PhoneNumber = record[headerMap["PhoneNumber"]]
	res.FaxNumber = record[headerMap["FaxNumber"]]
	res.ShipAddress = record[headerMap["ShipAddress"]]
	res.ShipCity = record[headerMap["ShipCity"]]
	res.ShipStateOrProvince = record[headerMap["ShipStateOrProvince"]]
	res.ShipZIPCode = record[headerMap["ShipZIPCode"]]
	res.ShipPhoneNumber = record[headerMap["ShipPhoneNumber"]]

	return res, nil
}

func parseEmployesRecrod(headerMap map[string]int, record []string) (entity.Employees, error) {
	var res entity.Employees
	var err error

	res.EmployeeID, err = strconv.ParseInt(record[headerMap["EmployeeID"]], 10, 64)
	if err != nil {
		return res, err
	}

	res.FirstName = record[headerMap["FirstName"]]
	res.LastName = record[headerMap["LastName"]]
	res.Title = record[headerMap["Title"]]
	res.WorkPhone = record[headerMap["WorkPhone"]]

	return res, nil

}

func parseShippingMethdosRecords(headerMap map[string]int, record []string) (entity.ShippingMethods, error) {
	var res entity.ShippingMethods
	var err error

	res.ShippingMethodID, err = strconv.ParseInt(record[headerMap["ShippingMethodID"]], 10, 64)
	if err != nil {
		return res, err
	}

	res.ShippingMethod = record[headerMap["ShippingMethod"]]

	return res, nil
}

func parseOrdersRecord(headerMap map[string]int, record []string) (entity.Orders, error) {
	var res entity.Orders
	var err error

	res.OrderID, err = strconv.ParseInt(record[headerMap["OrderID"]], 10, 64)
	if err != nil {
		return res, err
	}

	res.CustomerID, err = strconv.ParseInt(record[headerMap["CustomerID"]], 10, 64)
	if err != nil {
		return res, err
	}

	res.EmployeeID, err = strconv.ParseInt(record[headerMap["EmployeeID"]], 10, 64)
	if err != nil {
		return res, err
	}

	res.OrderDate, err = time.Parse("2006-01-02", record[headerMap["OrderDate"]])
	if err != nil {
		return res, err
	}

	res.PurchaseOrderNumber = record[headerMap["PurchaseOrderNumber"]]

	res.ShipDate, err = time.Parse("2006-01-02", record[headerMap["ShipDate"]])
	if err != nil {
		return res, err
	}

	res.ShippingMethodID, err = strconv.ParseInt(record[headerMap["ShippingMethodID"]], 10, 64)
	if err != nil {
		return res, err
	}

	res.FreightCharge, err = strconv.ParseFloat(record[headerMap["FreightCharge"]], 64)
	if err != nil {
		return res, err
	}

	res.Taxes, err = strconv.ParseFloat(record[headerMap["Taxes"]], 64)
	if err != nil {
		return res, err
	}

	res.PaymentReceived = record[headerMap["PaymentReceived"]]

	res.Comment = record[headerMap["Comment"]]

	return res, nil
}

func parseProductsRecord(headerMap map[string]int, record []string) (entity.Products, error) {
	var res entity.Products
	var err error

	res.ProductID, err = strconv.ParseInt(record[headerMap["OrderID"]], 10, 64)
	if err != nil {
		return res, err
	}

	res.ProductName = record[headerMap["ProductName"]]
	res.UnitPrice, err = strconv.ParseFloat(record[headerMap["UnitPrice"]], 64)
	if err != nil {
		return res, nil
	}
	res.InStock = record[headerMap["InStock"]]

	return res, nil

}

func parseOrderDetailsRecord(headerMap map[string]int, record []string) (entity.OrderDetails, error) {
	var res entity.OrderDetails
	var err error

	res.OrderDetailID, err = strconv.ParseInt(record[headerMap["OrderDetailID"]], 10, 64)
	if err != nil {
		return res, err
	}
	res.OrderID, err = strconv.ParseInt(record[headerMap["OrderID"]], 10, 64)
	if err != nil {
		return res, err
	}
	res.ProductID, err = strconv.ParseInt(record[headerMap["ProductID"]], 10, 64)
	if err != nil {
		return res, err
	}
	res.Quantity, err = strconv.ParseInt(record[headerMap["Quantity"]], 10, 64)
	if err != nil {
		return res, err
	}
	res.UnitPrice, err = strconv.ParseFloat(record[headerMap["UnitPrice"]], 64)
	if err != nil {
		return res, err
	}
	res.Discount, err = strconv.ParseFloat(record[headerMap["Discount"]], 64)
	if err != nil {
		return res, err
	}

	return res, nil
}
