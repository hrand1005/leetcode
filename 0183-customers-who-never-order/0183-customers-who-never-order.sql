# Write your MySQL query statement below
# SELECT name AS 'Customers'
# FROM Customers
# LEFT JOIN Orders ON Customers.id = Orders.customerId
# WHERE customerId IS NULL;

# SELECT A.name AS 'Customers'
# FROM Customers AS A
# LEFT JOIN Orders AS B ON A.id = B.customerId
# WHERE B.customerId IS NULL;

SELECT name AS 'Customers'
FROM Customers
WHERE NOT EXISTS (SELECT * FROM Orders WHERE Orders.customerId = Customers.Id LIMIT 1);