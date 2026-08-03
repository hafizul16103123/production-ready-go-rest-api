# production-ready-go-rest-api
# Migration
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate create -ext sql -dir migrations -seq create_users
## Run migration cmd
migrate -path migrations -database "postgres://postgres:postgresql@localhost:5432/postgres?sslmode=disable" up

migrate -path migrations -database "postgres://postgres:postgresql@localhost:5432/postgres?sslmode=disable" down

migrate -path migrations -database "postgres://postgres:postgresql@localhost:5432/postgres?sslmode=disable" force 2

force 2 tells migrate "pretend 000001 and 000002 already ran, and the DB isn't dirty" without executing their SQL.Run after 000002.
## Roll Back Undo the last migration:
migrate -path migrations -database "..." down 1

## Best Practices
✅ One migration = one logical change.
✅ Never edit an already-applied migration.
✅ Keep Up and Down migrations in sync.
✅ Commit migrations to Git.\
## Common Beginner Mistakes
❌ Editing old migration files after they have been applied.
❌ Forgetting the Down migration.
❌ Combining unrelated schema changes into one migration.
❌ Running SQL manually in production.

# JWT
go get github.com/golang-jwt/jwt/v5

# SQL
## Command sequence
- SELECT FROM JOIN WHERE
- GROUP BY HAVING
- ORDER BY LIMIT OFFSET
### Example:Almost every query follows this order.
- SELECT columns
- FROM table
- JOIN another_table ON condition
- WHERE conditions
- GROUP BY columns
- HAVING aggregate_condition
- ORDER BY columns
- LIMIT number
- OFFSET number;

### Example
- SELECT 
    - d.name AS department, 
    - AVG(e.salary) AS avg_salary
- FROM departments d
- JOIN employees e 
   - ON d.id = e.department_id
- WHERE e.status = 'Active'
- GROUP BY d.name
- HAVING AVG(e.salary) > 75000
- ORDER BY avg_salary DESC
- LIMIT 1
- OFFSET 0;`

### Template 1 — SELECT
SELECT column1,
       column2,
       column3
FROM table_name;

### Template 2 — SELECT *
SELECT *
FROM table_name;

### Template 3 — WHERE
SELECT columns
FROM table
WHERE condition; // WHERE id = 5

### Template 4 — Multiple WHERE
SELECT columns
FROM table
WHERE condition1
  AND condition2
  OR condition3;

### Template 5 — IN
SELECT columns
FROM table
WHERE column IN (
    value1,
    value2,
    value3
); // WHERE id IN (1,2,5,8)

### Template 6 — BETWEEN
SELECT columns
FROM table
WHERE column BETWEEN start AND end;

### Template 7 — LIKE
SELECT columns
FROM table
WHERE column LIKE pattern; // WHERE name LIKE 'Ha%'

- Starts with
    - WHERE name LIKE 'Ha%'
- Ends with
    - WHERE name LIKE '%ul'
- Contains
    - WHERE name LIKE '%afi%'

### Template 8 — ORDER BY
SELECT columns
FROM table
ORDER BY column ASC; // ORDER BY column DESC;

### Template 9 — LIMIT
SELECT columns
FROM table
LIMIT n;

### Template 10 — Pagination
SELECT columns
FROM table
ORDER BY id
LIMIT page_size
OFFSET skip_rows;

### Template 11 — DISTINCT
SELECT DISTINCT column
FROM table;

### Template 12 — Aggregate
SELECT
    aggregate(column)
FROM table;

#### Examples

SELECT COUNT(*)
FROM users;

SELECT SUM(price)
FROM orders;

SELECT AVG(price)
FROM orders;

SELECT MAX(price)
FROM products;

SELECT MIN(price)
FROM products;

### Template 13 — GROUP BY
SELECT
    column,
    aggregate(column2)
FROM table
GROUP BY column;

Example

SELECT
    country,
    COUNT(*)
FROM users
GROUP BY country;


### Template 14 — HAVING
SELECT
    column,
    aggregate(column2)
FROM table
GROUP BY column
HAVING aggregate(column2) > value;

Example:

SELECT
    country,
    COUNT(*)
FROM users
GROUP BY country
HAVING COUNT(*) > 10;

### Template 15 — INNER JOIN
SELECT columns
FROM table1
INNER JOIN table2 ON table1.column = table2.column;

Example

SELECT
    users.name,
    orders.total
FROM users
INNER JOIN orders
ON users.id = orders.user_id;

### Template 16 — LEFT JOIN
SELECT columns
FROM table1
LEFT JOIN table2
ON table1.column = table2.column;

Example

SELECT
    users.name,
    orders.total
FROM users
LEFT JOIN orders
ON users.id = orders.user_id;

- Returns every user, even if they have no orders.

### Template 17 — INSERT
INSERT INTO table (
    column1,
    column2,
    column3
)
VALUES (
    value1,
    value2,
    value3
);
 ### Template 18 — INSERT Multiple
INSERT INTO table (
    column1,
    column2
)
VALUES
(value1,value2),
(value3,value4),
(value5,value6);
### Template 19 — UPDATE
UPDATE table
SET
    column1 = value1,
    column2 = value2
WHERE condition;

Example

UPDATE users
SET
    name='Hafiz',
    age=25
WHERE id=5;
### Template 20 — DELETE
DELETE
FROM table
WHERE condition;

Example

DELETE
FROM users
WHERE id=5;
### Template 21 — UPSERT (PostgreSQL)
INSERT INTO table (
    columns...
)
VALUES (...)
ON CONFLICT (unique_column)
DO UPDATE
SET
    column = EXCLUDED.column;

Example

INSERT INTO users (
    id,
    name
)
VALUES (
    1,
    'Hafiz'
)
ON CONFLICT (id)
DO UPDATE
SET
    name = EXCLUDED.name;

### Template 22 — RETURNING
INSERT INTO table (...)
VALUES (...)
RETURNING *;

or

UPDATE table
SET ...
WHERE ...
RETURNING *;

### Template 23 — CTE (WITH)
WITH cte_name AS (
    SELECT ...
)
SELECT *
FROM cte_name;

Example

WITH active_users AS (
    SELECT *
    FROM users
    WHERE active = true
)
SELECT *
FROM active_users;

### Template 24 — Subquery
SELECT columns
FROM table
WHERE column IN (
    SELECT column
    FROM another_table
);

### Template 25 — EXISTS
SELECT *
FROM table1
WHERE EXISTS (
    SELECT 1
    FROM table2
    WHERE ...
);

### Template 26 — CASE
SELECT
    column,
    CASE
        WHEN condition THEN value
        WHEN condition THEN value
        ELSE value
    END
FROM table;