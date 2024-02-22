package main

import (
	"github/adefemi/fingreat_backend/utils"
)

func main() {
	// server := api.NewServer(".")
	// server.Start(8001)
	// scripts.AddAccountNumbers(".")
	utils.GenerateAccountNumber(1, "USD")
}
