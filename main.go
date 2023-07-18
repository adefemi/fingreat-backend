package main

import (
	"github/adefemi/fingreat_backend/api"
)

func main() {
	server := api.NewServer(".")
	server.Start(3000)
}
