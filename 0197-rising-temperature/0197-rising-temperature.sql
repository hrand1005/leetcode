# Write your MySQL query statement below

SELECT A.id 
FROM Weather AS A, Weather AS B
WHERE A.temperature > B.temperature 
    AND TO_DAYS(A.recordDate) - TO_DAYS(B.recordDate) = 1;