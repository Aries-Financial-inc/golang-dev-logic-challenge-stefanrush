package helpers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/models"
)

const dbDSN = "localhost:6379"

var ctx = context.Background()

// SaveContractAnalysis saves the given options contract to the database
func SaveContractAnalysis(contract *models.OptionsContract) error {
	db := redis.NewClient(&redis.Options{Addr: dbDSN})

	key := fmt.Sprintf("contract:%s", contract.ID)

	contractJSON, err := json.Marshal(contract)
	if err != nil {
		return err
	}

	if err := db.Set(ctx, key, contractJSON, 0).Err(); err != nil {
		return err
	}

	return nil
}

// RetrieveContractAnalysis retrieves the options contract with the given ID from the database
func RetrieveContractAnalysis(contractID string) (*models.OptionsContract, error) {
	db := redis.NewClient(&redis.Options{Addr: dbDSN})

	key := fmt.Sprintf("contract:%s", contractID)

	contract := &models.OptionsContract{}

	val, err := db.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(val), contract); err != nil {
		return nil, err
	}

	return contract, nil
}
