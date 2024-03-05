package api

import (
	db "github/adefemi/fingreat_backend/db/sqlc"
	"time"
)

type AccountResponse struct {
	ID            int64     `json:"id"`
	UserID        int32     `json:"user_id"`
	Balance       float64   `json:"balance"`
	Currency      string    `json:"currency"`
	CreatedAt     time.Time `json:"created_at"`
	AccountNumber string    `json:"account_number"`
}

func (u AccountResponse) ToAccountResponse(account *db.Account) *AccountResponse {
	return &AccountResponse{
		ID:            account.ID,
		UserID:        account.UserID,
		Balance:       account.Balance,
		Currency:      account.Currency,
		CreatedAt:     account.CreatedAt,
		AccountNumber: account.AccountNumber.String,
	}
}

func (u AccountResponse) ToAccountResponses(accounts []db.Account) []AccountResponse {
	accountResponses := make([]AccountResponse, len(accounts))

	for i := range accounts {
		accountResponses[i] = *u.ToAccountResponse(&accounts[i])
	}

	return accountResponses
}
