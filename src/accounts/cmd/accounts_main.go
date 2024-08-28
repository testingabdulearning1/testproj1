package main

import (
	"log"
	"school-management-app/common/config"
	"school-management-app/common/utils"
	accounts "school-management-app/src/accounts/config"
)

func main() {
	utils.ImportEnv()
	config.LoadConfig()

	log.Println(config.SERVER_TYPE)
	log.Println()

	if config.SERVER_TYPE == "accounts" {
		accounts.ConnectAccounts()
	}
}
