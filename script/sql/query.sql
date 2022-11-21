-- List of customers located in Irvine city
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
WHERE City = 'Irvine City'

--  List of customers whose order is handled by an employee named Adam Barr
SELECT
    DISTINCT
    cust.CustomerID
    ,cust.CompanyName
    ,cust.FirstName
    ,cust.LastName
    ,cust.BillingAddress
    ,cust.City
    ,cust.StateOfProvince
    ,cust.ZIPCode
    ,cust.Email
    ,cust.CompanyWebsite
    ,cust.PhoneNumber
    ,cust.FaxNumber
    ,cust.ShipAddress
    ,cust.ShipCity
    ,cust.ShipStateOrProvince
    ,cust.ShipZIPCode
    ,cust.ShipPhoneNumber
FROM Jatis.Customers cust
INNER JOIN Jatis.Orders ON Orders.CustomerID = cust.CustomerID
INNER JOIN Jatis.Employees ON Employees.EmployeeID = Orders.EmployeeID
WHERE Employees.FirstName = 'Adam' AND Employees.LastName = 'Barr'

--  List of products which are ordered by "Contonso, Ltd" Company
SELECT
    DISTINCT
    Products.ProductID
    ,Products.ProductName
    ,Products.UnitPrice
    ,Products.InStock
FROM Jatis.Products
INNER JOIN Jatis.Order_Details ON Order_Details.ProductID = Products.ProductID
INNER JOIN Jatis.Orders ON Orders.OrderID = Order_Details.OrderID
INNER JOIN Jatis.Customers ON Customers.CustomerID = Orders.CustomersID
WHERE Customers.CompanyName = 'Contonso, Ltd'

-- List of transactions (orders) which has "UPS Ground" as shipping method
SELECT
    Orders.OrderID
    ,Orders.CustomerID
    ,Orders.EmployeeID
    ,Orders.OrderDate
    ,Orders.PurchaseOrderNumber
    ,Orders.ShipDate
    ,Orders.ShippingMethodID
    ,Orders.FreightCharge
    ,Orders.Taxes
    ,Orders.PaymentReceived
    ,Orders.Comment
FROM Jatis.Orders
INNER JOIN Jatis.Shipping_Methods ON Orders.ShippingMethodID = Shipping_Methods.ShippingMethodID
WHERE Shipping_Methods.ShippingMethod = 'UPS Ground'


-- List of total cost (including tax and freight charge) for every order sorted by ship date
SELECT
    Orders.OrderID
    ,Orders.FreightCharge + Orders.Taxes as Total_Cost
    ,Orders.ShipDate
FROM Jatis.Orders
ORDER By ShipDate
