package payout

//Егор: добавил расчет выплат по ставкам

import (
	"FLUX/internal/bet/models"
	"FLUX/internal/generation/payback"
)

const (
	colorMultiplier  = 2.0
	numberMultiplier = 36.0
)

type PayoutResult struct {
	UserID string
	Win    float64
	IsWin  bool
}

func Calculate(bet models.Bet, result payback.Payback) PayoutResult {
	var win float64

	if string(bet.Fill.Color) == result.Color {
		win += bet.Stake * colorMultiplier
	}

	if bet.Fill.Number == result.Number {
		win += bet.Stake * numberMultiplier
	}

	return PayoutResult{
		UserID: bet.UserID,
		Win:    win,
		IsWin:  win > 0,
	}
}

func CalculateMany(bets []models.Bet, result payback.Payback) []PayoutResult {
	results := make([]PayoutResult, 0, len(bets))
	for _, bet := range bets {
		results = append(results, Calculate(bet, result))
	}
	return results
}
