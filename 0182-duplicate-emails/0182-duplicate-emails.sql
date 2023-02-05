# Write your MySQL query statement below
# WITH email_count AS (
#     SELECT Email, COUNT(Email) AS 'Count'
#     FROM Person
#     GROUP BY Email
# )
# SELECT DISTINCT person.Email
# FROM Person
# JOIN email_count ON person.Email = email_count.Email
# WHERE email_count.Count > 1;

SELECT Email
FROM Person
GROUP BY Email
HAVING COUNT(*) > 1;