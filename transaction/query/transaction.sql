-- name: ListTransactions :many
SELECT
	t.*
FROM
	transactions t
	JOIN accounts a ON t.from_account_id = a.id OR t.to_account_id = a.id
	JOIN users u ON u.id = a.user_id
WHERE
	u.id = $1
ORDER BY t.id DESC;


-- name: ListTransactionsByAccount :many
SELECT
	t.*
FROM
	transactions t
	JOIN accounts a ON t.from_account_id = a.id OR t.to_account_id = a.id
	JOIN users u ON u.id = a.user_id
WHERE
	u.id = $1 AND a.id = $2
ORDER BY t.id;

-- name: InsertTransaction :one
INSERT INTO transactions (from_account_id, to_account_id, amount)
		VALUES($1, $2, $3)
	RETURNING *;
