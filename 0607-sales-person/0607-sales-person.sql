# Write your MySQL query statement below
SELECT SalesPerson.name FROM SalesPerson
WHERE SalesPerson.sales_id NOT IN (
    SELECT SalesPerson.sales_id FROM SalesPerson
    LEFT JOIN Orders ON SalesPerson.sales_id = Orders.sales_id
    LEFT JOIN Company ON Orders.com_id = Company.com_id
    WHERE Company.name = "RED"
)
