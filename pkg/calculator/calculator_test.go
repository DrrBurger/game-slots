package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"game-slots/internal/models"
	"game-slots/internal/storage"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name           string
		reelsConfig    [][]string
		linesConfig    []models.LineConfiguration
		payoutsConfig  []models.PayoutConfiguration
		expectedResult *models.Result
		expectedError  string
	}{
		{
			name: "basic test",
			reelsConfig: [][]string{
				{"A", "B", "C", "D", "E"},
				{"F", "A", "F", "B", "C"},
				{"D", "E", "A", "G", "A"},
			},
			linesConfig: []models.LineConfiguration{
				{
					Line: 1,
					Positions: []models.Position{
						{Row: 0, Col: 0},
						{Row: 1, Col: 1},
						{Row: 2, Col: 2},
						{Row: 1, Col: 3},
						{Row: 0, Col: 4},
					},
				},
				{
					Line: 2,
					Positions: []models.Position{
						{Row: 2, Col: 0},
						{Row: 1, Col: 1},
						{Row: 0, Col: 2},
						{Row: 1, Col: 3},
						{Row: 2, Col: 4},
					},
				},
				{
					Line: 3,
					Positions: []models.Position{
						{Row: 1, Col: 0},
						{Row: 2, Col: 1},
						{Row: 1, Col: 2},
						{Row: 0, Col: 3},
						{Row: 1, Col: 4},
					},
				},
			},
			payoutsConfig: []models.PayoutConfiguration{
				{"A", []int{0, 0, 50, 100, 200}},
				{"B", []int{0, 0, 40, 80, 160}},
				{"C", []int{0, 0, 30, 60, 120}},
				{"D", []int{0, 0, 20, 40, 80}},
				{"E", []int{0, 0, 10, 20, 40}},
				{"F", []int{0, 0, 5, 10, 20}},
				{"G", []int{0, 0, 2, 5, 10}},
			},
			expectedResult: &models.Result{
				Lines: []models.LineResult{
					{Line: 1, Payout: 50},
					{Line: 2, Payout: 0},
					{Line: 3, Payout: 0},
				},
				Total: 50,
			},
			expectedError: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := storage.NewStorage()
			store.SaveReelConfiguration("reels", test.reelsConfig)
			store.SaveLineConfiguration("lines", test.linesConfig)
			store.SavePayoutConfiguration("payouts", test.payoutsConfig)

			result, err := Calculate(store, "reels", "lines", "payouts")
			if test.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedResult, result)
			}
		})
	}
}
