package calculator

import (
	"errors"

	"game-slots/internal/models"
	"game-slots/internal/storage"
)

// Calculate осуществляет расчет выплат на основе заданной конфигурации катушек, линий и выплат.
func Calculate(store *storage.Storage, reelsName, linesName, payoutsName string) (*models.Result, error) {
	reelsConfig, found := store.GetReelConfiguration(reelsName)
	if !found {
		return nil, errors.New("reels configuration not found")
	}

	linesConfig, found := store.GetLineConfiguration(linesName)
	if !found {
		return nil, errors.New("lines configuration not found")
	}

	payoutsConfig, found := store.GetPayoutConfiguration(payoutsName)
	if !found {
		return nil, errors.New("payouts configuration not found")
	}

	total := 0
	var lineResults []models.LineResult

	for _, lineConfig := range linesConfig {
		var symbols []string
		for _, pos := range lineConfig.Positions {
			symbols = append(symbols, reelsConfig[pos.Row][pos.Col])
		}
		payout := getPayout(symbols, payoutsConfig)
		total += payout
		lineResults = append(lineResults, models.LineResult{
			Line:   lineConfig.Line,
			Payout: payout,
		})
	}

	return &models.Result{
		Lines: lineResults,
		Total: total,
	}, nil
}

// getPayout вычисляет выплату для данной последовательности символов, используя предоставленную конфигурацию выплат.
func getPayout(symbols []string, payoutsConfig []models.PayoutConfiguration) int {
	if len(symbols) == 0 {
		return 0
	}

	firstSymbol := symbols[0]
	count := 1
	for _, s := range symbols[1:] {
		if s == firstSymbol {
			count++
		} else {
			break
		}
	}

	for _, payoutConfig := range payoutsConfig {
		if payoutConfig.Symbol == firstSymbol {
			return payoutConfig.Payout[count-1]
		}
	}

	return 0
}
