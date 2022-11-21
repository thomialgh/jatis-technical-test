CREATE DATABASE Jatis;
CREATE TABLE Jatis.Customers (
    CustomerID BIGINT NOT NULL AUTO_INCREMENT
    ,CompanyName VARCHAR(50)
    ,FirstName VARCHAR(30)
    ,LastName VARCHAR(50)
    ,BillingAddress VARCHAR(255)
    ,City VARCHAR(50)
    ,StateOfProvince VARCHAR(20)
    ,ZIPCode VARCHAR(20)
    ,Email VARCHAR(75)
    ,CompanyWebsite VARCHAR(200)
    ,PhoneNumber VARCHAR(30)
    ,FaxNumber VARCHAR(30)
    ,ShipAddress VARCHAR(255)
    ,ShipCity VARCHAR(50)
    ,ShipStateOrProvince VARCHAR(50)
    ,ShipZIPCode VARCHAR(20)
    ,ShipPhoneNumber VARCHAR(30)
    ,PRIMARY KEY (CustomerID)
);

CREATE TABLE Jatis.Employees (
    EmployeeID BIGINT NOT NULL AUTO_INCREMENT
    ,FirstName VARCHAR(50)
    ,LastName VARCHAR(50)
    ,Title VARCHAR(50)
    ,WorkPhone VARCHAR(30)
    ,PRIMARY KEY (EmployeeID)
);

CREATE TABLE Jatis.Shipping_Methods (
    ShippingMethodID INT NOT NULL AUTO_INCREMENT
    ,ShippingMethod VARCHAR(20)
    ,PRIMARY KEY (ShippingMethodID)
);

CREATE TABLE Orders (
    OrderID BIGINT NOT NULL AUTO_INCREMENT
    ,CustomerID BIGINT NOT NULL
    ,EmployeeID BIGINT NOT NULL
    ,OrderDate DATE
    ,PurchaseOrderNumber VARCHAR(30)
    ,ShipDate DATE
    ,ShippingMethodID INT NOT NULL
    ,FreightCharge FLOAT
    ,Taxes FLOAT
    ,PaymentReceived CHAR(1)
    ,Comment VARCHAR(150)
    ,PRIMARY KEY (OrderID)
    ,FOREIGN KEY (CustomerID) REFERENCES Jatis.Customers(CustomerID)
    ,FOREIGN KEY (EmployeeID) REFERENCES Jatis.Employees(EmployeeID)
    ,FOREIGN KEY (ShippingMethodID) REFERENCES Jatis.Shipping_Methods(ShippingMethodID)
);

CREATE TABLE Jatis.Products (
    ProductID INT NOT NULL AUTO_INCREMENT
    ,ProductName VARCHAR(20)
    ,UnitPrice FLOAT
    ,InStock CHAR(1)
    ,PRIMARY KEY (ProductID)
);

CREATE TABLE Order_Details (
    OrderDetailID BIGINT NOT NULL AUTO_INCREMENT
    ,OrderID BIGINT NOT NULL
    ,ProductID INT NOT NULL
    ,Quantity INT
    ,UnitPrice FLOAT
    ,Discount FLOAT
    ,PRIMARY KEY (OrderDetailID)
    ,FOREIGN KEY (OrderID) REFERENCES Jatis.Orders(OrderID)
    ,FOREIGN KEY (ProductID) REFERENCES Jatis.Products(ProductID)
);





