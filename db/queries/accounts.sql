-- name: CreateAccount :one
INSERT INTO accounts (
    user_id,
    currency,
    balance
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = $1;

-- name: GetAccountByUserID :many
SELECT * FROM accounts WHERE user_id = $1 ORDER BY created_at ASC;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateAccountBalance :one
UPDATE accounts SET balance = $1 WHERE id = $2 RETURNING *;

-- name: UpdateAccountNumber :one
UPDATE accounts SET account_number = $1 WHERE id = $2 RETURNING *;

-- name: UpdateAccountBalanceNew :one
UPDATE accounts SET balance = balance + sqlc.arg(amount) WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name: DeleteAllAccount :exec
DELETE FROM accounts;

-- name: GetAccountByAccountNumber :one
SELECT accounts.*, users.email 
FROM accounts 
JOIN users ON accounts.user_id = users.id 
WHERE account_number = $1;