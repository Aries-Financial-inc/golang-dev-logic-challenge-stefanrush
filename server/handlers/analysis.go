package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/helpers"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/models"
)

// OnPOSTAnalysis accepts an array of options contracts and returns the processed contracts
// that include the analysis results
func OnPOSTAnalysis(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v\n", r.Method, r.URL)

	var contracts []models.OptionsContract

	if err := json.NewDecoder(r.Body).Decode(&contracts); err != nil {
		log.Printf("Error parsing contracts from request body JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for idx := range contracts {
		contract := &contracts[idx]

		helpers.SanitizeContract(contract)
		if err := helpers.ValidateContract(contract); err != nil {
			log.Printf("Invalid contract: %v\n", err)
			continue
		}
		contract.IsValid = true
		helpers.AssignContractID(contract)

		contract.AnalysisResult = *helpers.AnalyzeContract(contract)

		if err := helpers.SaveContractAnalysis(contract); err != nil {
			log.Printf("Error saving contract analysis: %v\n", err)
		}
	}

	if err := json.NewEncoder(w).Encode(contracts); err != nil {
		log.Printf("Error encoding response body JSON: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// OnGETAnalysisResult returns the analysis result for the given contract ID
func OnGETAnalysisResult(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v\n", r.Method, r.URL)

	contractID := r.PathValue("contract_id")
	if contractID == "" {
		log.Printf("No contract ID provided\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contract, err := helpers.RetrieveContractAnalysis(contractID)
	if err != nil {
		if err == redis.Nil {
			log.Printf("Contract not found: %v\n", contractID)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Printf("Error retrieving contract analysis from database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(contract); err != nil {
		log.Printf("Error encoding response body JSON: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// OnGETAnalysisResultGraph renders a line graph of the analysis result for the given contract ID
func OnGETAnalysisResultGraph(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v\n", r.Method, r.URL)

	contractID := r.PathValue("contract_id")
	if contractID == "" {
		log.Printf("No contract ID provided\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contract, err := helpers.RetrieveContractAnalysis(contractID)
	if err != nil {
		if err == redis.Nil {
			log.Printf("Contract not found: %v\n", contractID)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Printf("Error retrieving contract analysis from database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	line := helpers.PlotLineGraph(contract)

	line.Render(w)
}
