// seeder/seeder.go
package seeder

import (
	"math/rand"
	"minibank/domain"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)



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

func GenAccounts(amount int) []domain.Account {
	accounts := make([]domain.Account, amount)
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

		accounts[i] = domain.Account{
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
