package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/models"
)

// SanitizeContract normalizes the contract type and long/short values
func SanitizeContract(contract *models.OptionsContract) {
	contract.Type = models.OptionsType(strings.ToLower(string(contract.Type)))
	contract.LongShort = models.OptionsLongShort(strings.ToLower(string(contract.LongShort)))
}

// ValidateContract returns an error if the contract is invalid, otherwise nil
func ValidateContract(contract *models.OptionsContract) error {
	errs := []error{}

	contract.Type = models.OptionsType(strings.ToLower(string(contract.Type)))
	switch contract.Type {
	case models.Call:
	case models.Put:
	default:
		errs = append(errs, fmt.Errorf("Invalid contract type: %v", contract.Type))
	}

	contract.LongShort = models.OptionsLongShort(strings.ToLower(string(contract.LongShort)))
	switch contract.LongShort {
	case models.Long:
	case models.Short:
	default:
		errs = append(errs, fmt.Errorf("Invalid contract long/short: %v", contract.LongShort))
	}

	if contract.StrikePrice <= 0 {
		errs = append(errs, fmt.Errorf("Invalid contract strike price: %v", contract.StrikePrice))
	}

	if contract.Ask <= 0 {
		errs = append(errs, fmt.Errorf("Invalid contract ask price: %v", contract.Ask))
	}

	if contract.Bid <= 0 {
		errs = append(errs, fmt.Errorf("Invalid contract bid price: %v", contract.Bid))
	}

	if contract.ExpirationDate.IsZero() {
		errs = append(
			errs,
			fmt.Errorf("Invalid contract expiration date: %v", contract.ExpirationDate),
		)
	}

	return errors.Join(errs...)
}

// AssignContractID assigns a new random UUID to the given contract
func AssignContractID(contract *models.OptionsContract) {
	contract.ID = uuid.New()
}
