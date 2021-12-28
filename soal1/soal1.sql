SELECT u.ID, u.UserName , u2.UserName AS "ParentUserName"
FROM USERS u 
LEFT JOIN USERS u2 ON u.Parent = u2.ID