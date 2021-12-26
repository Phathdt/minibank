-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1;

-- name: UpdateBalanceAccount :one
UPDATE
	accounts
SET
	balance = balance + @balance
WHERE
	id = @account_id
RETURNING *;
