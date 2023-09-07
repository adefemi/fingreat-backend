-- name: CreateMoneyRecord :one 
INSERT INTO money_records (
    user_id, reference, status, amount
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetMoneyRecordByReference :one
SELECT * FROM money_records WHERE reference = $1;

-- name: GetMoneyRecordsByStatus :many
SELECT * FROM money_records WHERE status = $1;

-- name: DeleteMoneyRecordById :exec
DELETE FROM money_records WHERE id = $1;