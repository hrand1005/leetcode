# Write your MySQL query statement below
# SELECT emp.name AS 'Employee'
# FROM Employee AS emp, Employee AS man
# WHERE emp.managerId = man.id AND emp.salary > man.salary;

SELECT emp.name AS 'Employee'
FROM Employee AS emp
JOIN Employee AS man ON emp.managerId = man.id
WHERE emp.salary > man.salary;