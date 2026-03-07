// seeder/seeder.go
package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
)

type Account struct {
	AccountID     string `json:"account_id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	AccountNumber string `json:"account_number"`
	Type          string `json:"type"`
	Dob           string `json:"dob"`
	Balance       int64  `json:"balance"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func weightedChoice(choices map[string]float64) string {
	total := 0.0
	for _, w := range choices {
		total += w
	}
	r := rand.Float64() * total
	for k, w := range choices {
		r -= w
		if r <= 0 {
			return k
		}
	}
	return ""
}

func genAccounts(amount int) []Account {
	accounts := make([]Account, amount)
	for i := range accounts {
		dob := gofakeit.DateRange(
			time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2005, 12, 31, 0, 0, 0, 0, time.UTC),
		)
		createdAt := gofakeit.DateRange(
			dob.AddDate(18, 0, 0),
			time.Now(),
		)
		updatedAt := gofakeit.DateRange(
			createdAt,
			time.Now(),
		)

		accounts[i] = Account{
			AccountID:     gofakeit.UUID(),
			FullName:      gofakeit.Name(),
			Email:         gofakeit.Email(),
			Password:      gofakeit.Password(true, true, true, false, false, 32),
			Phone:         gofakeit.Numerify("+84#########"),
			AccountNumber: gofakeit.Numerify("############"),
			Type:          weightedChoice(map[string]float64{"checking": 0.7, "savings": 0.3}),
			Status:        weightedChoice(map[string]float64{"active": 0.8, "frozen": 0.1, "closed": 0.1}),
			Balance:       int64(gofakeit.IntRange(0, 10000000)),
			Dob:           dob.Format("2006-01-02"),
			CreatedAt:     createdAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:     updatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}
	return accounts
}

func main() {
	env := godotenv.Load("../.env")
	if env != nil {
		panic("Error loading .env file")
	}
	url := os.Getenv("DB_URL")
	client, ctx, cancel := connectDB(url)
	accounts := genAccounts(10)
	insertMany(client, ctx, "minibank", "accounts", accounts)
	defer disconnectDB(client, ctx, cancel)
}
