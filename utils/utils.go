package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

var Currencies = map[string]string{
	"USD": "USD",
	"NGN": "NGN",
	"GBP": "GBP",
}

func IsValidCurrency(currency string) bool {
	if _, ok := Currencies[currency]; ok {
		return true
	}
	return false
}

func GetActiveUser(c *gin.Context) (int64, error) {
	value, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to access resources"})
		return 0, fmt.Errorf("error occurred ")
	}

	userId, ok := value.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Encountered an issue"})
		return 0, fmt.Errorf("error occurred ")
	}

	return userId, nil
}

func HandleError(err error, c *gin.Context, gValid galidator.Validator) interface{} {
	if c.Request.ContentLength == 0 {
		return "provide body"
	}

	if e, ok := err.(*json.UnmarshalTypeError); ok {
		if e.Field == "" {
			return "provide a json body"
		}
		msg := fmt.Sprintf("Invalid value for field '%s'. Expected a value of type '%s'", e.Field, e.Type)
		return msg
	}

	return gValid.DecryptErrors(err)
}

func GetDBSource(config *Config, dbName string) string {
	// return the structure postgres://root:secret@localhost:5432/
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DB_username, config.DB_password, config.DB_host, config.DB_port, dbName)
}
