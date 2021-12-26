-- name: GetAccount :one
SELECT *
FROM accounts
WHERE user_id = $1
  AND id = $2 LIMIT 1;

-- name: UpdateBalanceAccount :one
UPDATE
	accounts
SET
	balance = balance + @balance
WHERE
	id = @account_id
RETURNING *;
