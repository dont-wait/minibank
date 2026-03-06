// seeder/seeder.go
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

type Account struct {
	AccountID     string  `json:"account_id"`
	FullName      string  `json:"full_name"`
	Email         string  `json:"email"`
	Password      string  `json:"password"`
	Phone         string  `json:"phone"`
	AccountNumber string  `json:"account_number"`
	Type          string  `json:"type"`
	Balance       float64 `json:"balance"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
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
		accounts[i] = Account{
			AccountID:     gofakeit.UUID(),
			FullName:      gofakeit.Name(),
			Email:         gofakeit.Email(),
			Password:      gofakeit.Password(true, true, true, false, false, 32),
			Phone:         gofakeit.Phone(),
			AccountNumber: gofakeit.Numerify("############"),
			Type:          weightedChoice(map[string]float64{"checking": 0.7, "savings": 0.3}),
			Balance:       gofakeit.Float64Range(0, 10_000_000),
			Status:        weightedChoice(map[string]float64{"active": 0.8, "frozen": 0.1, "closed": 0.1}),
			CreatedAt:     gofakeit.Date().Format("2006-01-02T15:04:05Z"),
			UpdatedAt:     gofakeit.Date().Format("2006-01-02T15:04:05Z"),
		}
	}
	return accounts
}

func main() {
	gofakeit.Seed(0xff)
	accounts := genAccounts(10)
	out, _ := json.MarshalIndent(accounts, "", "  ")
	fmt.Println(string(out))
}
