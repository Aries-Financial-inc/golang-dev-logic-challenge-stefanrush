package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/helpers"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/models"
)

func OnPOSTAnalysis(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v\n", r.Method, r.URL)

	var contracts []models.OptionsContract

	if err := json.NewDecoder(r.Body).Decode(&contracts); err != nil {
		log.Printf("Error parsing contracts from JSON: %v\n", err)
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
		log.Printf("Error encoding response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func OnGETAnalysisResult(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v\n", r.Method, r.URL)

	contractID := r.PathValue("contract_id")
	if contractID == "" {
		log.Printf("No contract ID provided\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contract, err := helpers.GetContractAnalysis(contractID)
	if err != nil {
		log.Printf("Error getting contract analysis: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(contract); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func OnGETAnalysisResultGraph(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v\n", r.Method, r.URL)

	contractID := r.PathValue("contract_id")
	if contractID == "" {
		log.Printf("No contract ID provided\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contract, err := helpers.GetContractAnalysis(contractID)
	if err != nil {
		log.Printf("Error getting contract analysis: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	line := helpers.PlotLineGraph(contract)

	line.Render(w)
}
