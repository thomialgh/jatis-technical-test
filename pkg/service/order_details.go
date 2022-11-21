package service

import (
	"database/sql"
	"fmt"
	"jatis/entity"
	"jatis/pkg/domain"
	"jatis/pkg/repo"
)

var (
	ErrOrderNotFound = fmt.Errorf("Order Not Found")
)

func GetOrderDetails(orderID int64) (*entity.OrdersResponse, error) {
	var res = new(entity.OrdersResponse)
	order, err := domain.GetOrder(repo.DB, orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}

	res.OrderID = order.OrderID
	res.OrderDate = order.OrderDate
	res.PurchaseOrderNumber = order.PurchaseOrderNumber
	res.ShipDate = order.ShipDate
	res.FreightCharge = order.FreightCharge
	res.Taxes = order.Taxes
	res.PaymentReceived = order.PaymentReceived
	res.Comment = order.Comment

	customer, err := domain.GetCustomer(repo.DB, order.CustomerID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	res.Customers = entity.CustomerResponse{
		CustomerID:          customer.CustomerID,
		CompanyName:         customer.CompanyName,
		FirstName:           customer.FirstName,
		LastName:            customer.LastName,
		BillingAddress:      customer.BillingAddress,
		City:                customer.City,
		StateOfProvince:     customer.StateOfProvince,
		ZIPCode:             customer.ZIPCode,
		Email:               customer.Email,
		CompanyWebsite:      customer.CompanyWebsite,
		PhoneNumber:         customer.PhoneNumber,
		FaxNumber:           customer.FaxNumber,
		ShipAddress:         customer.ShipAddress,
		ShipCity:            customer.ShipCity,
		ShipStateOrProvince: customer.ShipStateOrProvince,
		ShipZIPCode:         customer.ShipZIPCode,
		ShipPhoneNumber:     customer.ShipPhoneNumber,
	}

	employee, err := domain.GetEmployee(repo.DB, order.EmployeeID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	res.Employees = entity.EmployeeResponse{
		EmployeeID: employee.EmployeeID,
		FirstName:  employee.FirstName,
		LastName:   employee.LastName,
		Title:      employee.Title,
		WorkPhone:  employee.WorkPhone,
	}

	shippingMethod, err := domain.GetShippingMethods(repo.DB, order.ShippingMethodID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	res.ShippingMethod = entity.ShippingMethodsResponse{
		ShippingMethodID: shippingMethod.ShippingMethodID,
		ShippingMethod:   shippingMethod.ShippingMethod,
	}

	detailOrder, productIds, err := domain.GetOrderDetailsByOrderID(repo.DB, orderID)
	if err != nil {
		return nil, err
	}

	mProduct, err := domain.GetProductsMap(repo.DB, productIds)
	if err != nil {
		return nil, err
	}

	for _, v := range detailOrder {
		od := entity.OrderDetailResponse{
			OrderDetailID: v.OrderDetailID,
			OrderID:       v.OrderID,
			ProductID:     v.ProductID,
			Quantity:      v.Quantity,
			UnitPrice:     v.UnitPrice,
			Discount:      v.Discount,
		}

		od.Products = entity.ProductResponse{
			ProductID:   mProduct[v.ProductID].ProductID,
			ProductName: mProduct[v.ProductID].ProductName,
			UnitPrice:   mProduct[v.ProductID].UnitPrice,
			InStock:     mProduct[v.ProductID].InStock,
		}

		res.OrderDetails = append(res.OrderDetails, od)
		res.TotalPayment += mProduct[v.ProductID].UnitPrice*float64(od.Quantity) - od.Discount
	}

	res.TotalPayment += res.Taxes

	return res, nil

}
