# Write your MySQL query statement below
# WITH result AS (
#     SELECT Courses.class, COUNT(Courses.class) AS students
#     FROM Courses
#     GROUP BY Courses.class
# )
# SELECT result.class FROM result
# WHERE result.students >= 5;

SELECT class FROM Courses
GROUP BY class
HAVING COUNT(DISTINCT student) >= 5;