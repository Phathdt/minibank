-- name: ListTransactions :many
SELECT
	t.*
FROM
	transactions t
	JOIN accounts a ON t.account_id = a.id
	JOIN users u ON u.id = a.user_id
WHERE
	u.id = $1
ORDER BY t.id DESC;


-- name: ListTransactionsByAccount :many
SELECT
	t.*
FROM
	transactions t
	JOIN accounts a ON t.account_id = a.id
	JOIN users u ON u.id = a.user_id
WHERE
	u.id = $1 AND t.account_id = $2
ORDER BY t.id;

-- name: InsertTransaction :one
INSERT INTO transactions (account_id, amount, transaction_type)
		VALUES($1, $2, $3)
	RETURNING *;
