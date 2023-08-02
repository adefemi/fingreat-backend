// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: accounts.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    user_id,
    currency,
    balance
) VALUES ($1, $2, $3) RETURNING id, user_id, balance, currency, created_at
`

type CreateAccountParams struct {
	UserID   int32   `json:"user_id"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.UserID, arg.Currency, arg.Balance)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const deleteAllAccount = `-- name: DeleteAllAccount :exec
DELETE FROM accounts
`

func (q *Queries) DeleteAllAccount(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllAccount)
	return err
}

const getAccountByID = `-- name: GetAccountByID :one
SELECT id, user_id, balance, currency, created_at FROM accounts WHERE id = $1
`

func (q *Queries) GetAccountByID(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByID, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountByUserID = `-- name: GetAccountByUserID :many
SELECT id, user_id, balance, currency, created_at FROM accounts WHERE user_id = $1
`

func (q *Queries) GetAccountByUserID(ctx context.Context, userID int32) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getAccountByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, user_id, balance, currency, created_at FROM accounts ORDER BY id
LIMIT $1 OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccountBalance = `-- name: UpdateAccountBalance :one
UPDATE accounts SET balance = $1 WHERE id = $2 RETURNING id, user_id, balance, currency, created_at
`

type UpdateAccountBalanceParams struct {
	Balance float64 `json:"balance"`
	ID      int64   `json:"id"`
}

func (q *Queries) UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountBalance, arg.Balance, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const updateAccountBalanceNew = `-- name: UpdateAccountBalanceNew :one
UPDATE accounts SET balance = balance + $1 WHERE id = $2
RETURNING id, user_id, balance, currency, created_at
`

type UpdateAccountBalanceNewParams struct {
	Amount float64 `json:"amount"`
	ID     int64   `json:"id"`
}

func (q *Queries) UpdateAccountBalanceNew(ctx context.Context, arg UpdateAccountBalanceNewParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountBalanceNew, arg.Amount, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}
