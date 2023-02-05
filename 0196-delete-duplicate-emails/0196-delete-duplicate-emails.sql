# Please write a DELETE statement and DO NOT write a SELECT statement.
# Write your MySQL query statement below

# WITH MinIds AS (
#     SELECT MIN(id) as 'MinId', Email
#     FROM Person
#     GROUP BY id
# )
# DELETE  
# FROM Person
# WHERE id NOT IN (SELECT MinIds.MinId FROM MinIds);

DELETE P1
FROM Person AS P1, Person AS P2
WHERE P1.Email = P2.Email AND P1.id > P2.id;
