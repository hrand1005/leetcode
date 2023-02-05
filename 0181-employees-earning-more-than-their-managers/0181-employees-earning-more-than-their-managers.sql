# Write your MySQL query statement below
SELECT emp.name AS 'Employee'
FROM Employee AS emp, Employee AS man
WHERE emp.managerId = man.id AND emp.salary > man.salary;