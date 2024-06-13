package helpers

import (
	"testing"
	"time"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/models"
)

var validContracts = []*models.OptionsContract{
	{
		Type:           models.Call,
		LongShort:      models.Long,
		StrikePrice:    250,
		ExpirationDate: time.Now().AddDate(0, 0, 30),
		Bid:            5.54,
		Ask:            6.61,
	},
	{
		Type:           models.Put,
		LongShort:      models.Short,
		StrikePrice:    250,
		ExpirationDate: time.Now().AddDate(0, 0, 30),
		Bid:            5.54,
		Ask:            6.61,
	},
}

var invalidContracts = []*models.OptionsContract{
	{
		Type:           "cell",
		LongShort:      models.Short,
		StrikePrice:    250,
		ExpirationDate: time.Now().AddDate(0, 0, 30),
		Bid:            5.54,
		Ask:            6.61,
	},
	{
		Type:           models.Call,
		LongShort:      "lung",
		StrikePrice:    250,
		ExpirationDate: time.Now().AddDate(0, 0, 30),
		Bid:            5.54,
		Ask:            6.61,
	},
	{
		Type:           models.Call,
		LongShort:      models.Long,
		StrikePrice:    -1,
		ExpirationDate: time.Now().AddDate(0, 0, 30),
		Bid:            5.54,
		Ask:            6.61,
	},
	{
		Type:           models.Call,
		LongShort:      models.Long,
		StrikePrice:    250,
		ExpirationDate: time.Now().AddDate(0, 0, 30),
		Bid:            -1,
		Ask:            6.61,
	},
	{
		Type:           models.Call,
		LongShort:      models.Long,
		StrikePrice:    250,
		ExpirationDate: time.Now().AddDate(0, 0, 30),
		Bid:            5.54,
		Ask:            -1,
	},
}

func TestValidateContract(t *testing.T) {
	for _, contract := range validContracts {
		if err := ValidateContract(contract); err != nil {
			t.Errorf("Expected no error for valid contract, got: %v", err)
		}
	}

	for _, contract := range invalidContracts {
		if err := ValidateContract(contract); err == nil {
			t.Errorf("Expected error for invalid contract, got nil")
		}
	}
}
