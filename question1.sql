SELECT child."ID", child."UserName", parent."UserName" as "ParentUserName" 
FROM "USER" as "parent" RIGHT JOIN "USER" as "child" ON parent."ID" = child."Parent" 
ORDER BY child."UserName" ASC;