package models

import (
	"fmt"
)

type BetColor string  

const (
	Red BetColor = "red"
	Black BetColor = "black"
	Green BetColor = "green"
)

type BetFill struct {
	Color BetColor 
	Number int 
}
type Bet struct {
	ID int 
	UserID string 
	Fill BetFill
	Stake float64
	Result bool 
}

func (b *Bet) Validate() error {
	if b.ID < 0 {
		return fmt.Errorf("invalid ID")
	} 
	if b.UserID == "" {
		return fmt.Errorf("invalid UserID")
	}
	if b.Stake < 1 {
		return fmt.Errorf("invalid Stake")
	}
	return nil
}