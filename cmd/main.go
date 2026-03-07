package main

import (
	"minibank/infra"
	"minibank/seeder"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load("../.env")
	if env != nil {
		panic("Error loading .env file")
	}
	url := os.Getenv("DB_URL")
	client, ctx, cancel := infra.ConnectDB(url)
	accounts := seeder.GenAccounts(10)
	infra.InsertMany(client, ctx, "minibank", "accounts", accounts)
	defer infra.DisconnectDB(client, ctx, cancel)
}
