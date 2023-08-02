package api

import (
	"context"
	"database/sql"
	db "github/adefemi/fingreat_backend/db/sqlc"
	"github/adefemi/fingreat_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Account struct {
	server *Server
}

func (a Account) router(server *Server) {
	a.server = server

	serverGroup := server.router.Group("/account", AuthenticatedMiddleware())
	serverGroup.POST("create", a.createAccount)
	serverGroup.GET("", a.getUserAccounts)
	serverGroup.POST("transfer", a.transfer)
}

type AccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (a *Account) createAccount(c *gin.Context) {
	userId, err := utils.GetActiveUser(c)
	if err != nil {
		return
	}

	acc := new(AccountRequest)

	if err := c.ShouldBindJSON(&acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateAccountParams{
		UserID:   int32(userId),
		Currency: acc.Currency,
		Balance:  0,
	}

	account, err := a.server.queries.CreateAccount(context.Background(), arg)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "you already have an account with this currency."})
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

func (a *Account) getUserAccounts(c *gin.Context) {
	userId, err := utils.GetActiveUser(c)
	if err != nil {
		return
	}

	accounts, err := a.server.queries.GetAccountByUserID(context.Background(), int32(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

type TransferRequest struct {
	ToAccountID   int32   `json:"to_account_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	FromAccountID int32   `json:"from_account_id" binding:"required"`
}

func (a *Account) transfer(c *gin.Context) {
	userId, err := utils.GetActiveUser(c)
	if err != nil {
		return
	}

	tr := new(TransferRequest)

	if err := c.ShouldBindJSON(&tr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := a.server.queries.GetAccountByID(context.Background(), int64(tr.FromAccountID))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't get account"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if account.UserID != int32(userId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't get account"})
		return
	}

	toAccount, err := a.server.queries.GetAccountByID(context.Background(), int64(tr.ToAccountID))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find account to send to"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if toAccount.Currency != account.Currency {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Currencies of account do not match"})
		return
	}

	if account.Balance < tr.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have enough balance"})
		return
	}

	txArg := db.CreateTransferParams{
		FromAccountID: tr.FromAccountID,
		ToAccountID:   tr.ToAccountID,
		Amount:        tr.Amount,
	}

	tx, err := a.server.queries.TransferTx(context.Background(), txArg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encountered issue with transaction"})
		return
	}

	c.JSON(http.StatusCreated, tx)
}
