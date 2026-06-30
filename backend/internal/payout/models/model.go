package models 

import (
	"FLUX/internal/generation/payback"
)

type PayoutResult struct {
	UserID string
	Fill payback.Payback
	Gain float64
	IsWin bool
}

type Inter struct {
	
}