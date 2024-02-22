package scripts

import (
	"database/sql"
	"fmt"
	"github/adefemi/fingreat_backend/utils"

	_ "github.com/lib/pq"
)

func AddAccountNumbers(envPath string) {
	config, err := utils.LoadConfig(envPath)
	if err != nil {
		panic(fmt.Sprintf("Couldn't load config: %v", err))
	}
	db_source := utils.GetDBSource(config, config.DB_name)
	conn, err := sql.Open(config.DBdriver, db_source)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	accounts, err := conn.Query("SELECT id, currency FROM accounts WHERE account_number IS NULL")
	if err != nil {
		panic(fmt.Sprintf("Could not query accounts: %v", err))
	}

	for accounts.Next() {
		var id int64
		var currency string

		err = accounts.Scan(&id, &currency)

		if err != nil {
			panic(fmt.Sprintf("Could not scan accounts: %v", err))
		}

		accountNumber, err := utils.GenerateAccountNumber(id, currency)
		if err != nil {
			panic(err)
		}

		_, err = conn.Exec("UPDATE accounts SET account_number = $1 WHERE id = $2", accountNumber, id)

		if err != nil {
			panic(fmt.Sprintf("Could not update accounts: %v", err))
		}
	}

	fmt.Println("Account numbers added successfully")
}
