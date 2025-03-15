package models

import (
	"errors"
	"time"
)

type Stock struct {
	Ticker     string    `json:"ticker"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Company    string    `json:"company"`
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	Time       time.Time `json:"time"`
}

func (s *Stock) Validate() error {
	if s.Ticker == "" || s.Company == "" || s.Action == "" {
		return errors.New("ticker, company, and action fields cannot be empty")
	}
	return nil
}
