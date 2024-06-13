package models

import (
	"time"

	"github.com/google/uuid"
)

// An OptionsContract represents an options contract
type OptionsContract struct {
	ID             uuid.UUID        `json:"id"`
	Type           OptionsType      `json:"type"`
	LongShort      OptionsLongShort `json:"long_short"`
	StrikePrice    float64          `json:"strike_price"`
	ExpirationDate time.Time        `json:"expiration_date"`
	Bid            float64          `json:"bid"`
	Ask            float64          `json:"ask"`
	IsValid        bool             `json:"is_valid"`
	AnalysisResult AnalysisResult   `json:"analysis_result"`
}

type OptionsType string

const (
	Call OptionsType = "call"
	Put  OptionsType = "put"
)

type OptionsLongShort string

const (
	Long  OptionsLongShort = "long"
	Short OptionsLongShort = "short"
)
