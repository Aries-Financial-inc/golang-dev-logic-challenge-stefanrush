package helpers

import (
	"math"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/models"
)

// AnalyzeContract runs a profit/loss analysis on the given contract and returns the result
func AnalyzeContract(contract *models.OptionsContract) *models.AnalysisResult {
	var result *models.AnalysisResult

	switch contract.Type {
	case models.Call:
		switch contract.LongShort {
		case models.Long:
			result = analyzeLongCall(contract)
		case models.Short:
			result = analyzeShortCall(contract)
		}
	case models.Put:
		switch contract.LongShort {
		case models.Long:
			result = analyzeLongPut(contract)
		case models.Short:
			result = analyzeShortPut(contract)
		}
	}

	return result
}

const (
	stockPriceRangeInterval float64 = 1
	stockPriceRangeWindow   float64 = 0.5
)

// analyzeLongCall runs a profit/loss analysis on a long call contract
func analyzeLongCall(contract *models.OptionsContract) *models.AnalysisResult {
	stock_price_range_start := math.Floor(contract.StrikePrice * stockPriceRangeWindow)
	stock_price_range_end := math.Ceil(contract.StrikePrice / stockPriceRangeWindow)

	analyze := func(purchase_price float64) ([]models.XYPoint, models.PLData) {
		graphData := []models.XYPoint{}

		current_stock_price := stock_price_range_start
		for current_stock_price <= stock_price_range_end {
			profit := -purchase_price * 100
			if current_stock_price >= contract.StrikePrice {
				profit = (current_stock_price-contract.StrikePrice)*100 - purchase_price*100
			}

			graphData = append(graphData, models.XYPoint{
				X: current_stock_price,
				Y: roundToHundredth(profit),
			})

			current_stock_price += stockPriceRangeInterval
		}

		plData := models.PLData{
			MaxProfit:      math.MaxFloat64,
			MaxLoss:        roundToHundredth(purchase_price * 100),
			BreakEvenPoint: roundToHundredth(contract.StrikePrice + purchase_price),
		}

		return graphData, plData
	}

	bidGraphData, bidPLData := analyze(contract.Bid)
	askGraphData, askPLData := analyze(contract.Ask)

	return &models.AnalysisResult{
		BidGraphData: bidGraphData,
		BidPLData:    bidPLData,
		AskGraphData: askGraphData,
		AskPLData:    askPLData,
	}
}

// analyzeShortCall runs a profit/loss analysis on a short call contract
func analyzeShortCall(contract *models.OptionsContract) *models.AnalysisResult {
	stock_price_range_start := math.Floor(contract.StrikePrice * stockPriceRangeWindow)
	stock_price_range_end := math.Ceil(contract.StrikePrice / stockPriceRangeWindow)

	analyze := func(purchase_price float64) ([]models.XYPoint, models.PLData) {
		graphData := []models.XYPoint{}

		current_stock_price := stock_price_range_start
		for current_stock_price <= stock_price_range_end {
			profit := purchase_price * 100
			if current_stock_price >= contract.StrikePrice {
				profit = purchase_price*100 - (current_stock_price-contract.StrikePrice)*100
			}

			graphData = append(graphData, models.XYPoint{
				X: current_stock_price,
				Y: roundToHundredth(profit),
			})

			current_stock_price += stockPriceRangeInterval
		}

		plData := models.PLData{
			MaxProfit:      roundToHundredth(purchase_price * 100),
			MaxLoss:        math.MaxFloat64,
			BreakEvenPoint: roundToHundredth(contract.StrikePrice + purchase_price),
		}

		return graphData, plData
	}

	bidGraphData, bidPLData := analyze(contract.Bid)
	askGraphData, askPLData := analyze(contract.Ask)

	return &models.AnalysisResult{
		BidGraphData: bidGraphData,
		BidPLData:    bidPLData,
		AskGraphData: askGraphData,
		AskPLData:    askPLData,
	}
}

// analyzeLongPut runs a profit/loss analysis on a long put contract
func analyzeLongPut(contract *models.OptionsContract) *models.AnalysisResult {
	stock_price_range_start := math.Floor(contract.StrikePrice * stockPriceRangeWindow)
	stock_price_range_end := math.Ceil(contract.StrikePrice / stockPriceRangeWindow)

	analyze := func(purchase_price float64) ([]models.XYPoint, models.PLData) {
		graphData := []models.XYPoint{}

		current_stock_price := stock_price_range_start
		for current_stock_price <= stock_price_range_end {
			profit := -purchase_price * 100
			if current_stock_price < contract.StrikePrice {
				profit = (contract.StrikePrice-current_stock_price)*100 - purchase_price*100
			}

			graphData = append(graphData, models.XYPoint{
				X: current_stock_price,
				Y: roundToHundredth(profit),
			})

			current_stock_price += stockPriceRangeInterval
		}

		plData := models.PLData{
			MaxProfit:      roundToHundredth(contract.StrikePrice - purchase_price),
			MaxLoss:        roundToHundredth(purchase_price * 100),
			BreakEvenPoint: roundToHundredth(contract.StrikePrice - purchase_price),
		}

		return graphData, plData
	}

	bidGraphData, bidPLData := analyze(contract.Bid)
	askGraphData, askPLData := analyze(contract.Ask)

	return &models.AnalysisResult{
		BidGraphData: bidGraphData,
		BidPLData:    bidPLData,
		AskGraphData: askGraphData,
		AskPLData:    askPLData,
	}
}

// analyzeShortPut runs a profit/loss analysis on a short put contract
func analyzeShortPut(contract *models.OptionsContract) *models.AnalysisResult {
	stock_price_range_start := math.Floor(contract.StrikePrice * stockPriceRangeWindow)
	stock_price_range_end := math.Ceil(contract.StrikePrice / stockPriceRangeWindow)

	analyze := func(purchase_price float64) ([]models.XYPoint, models.PLData) {
		graphData := []models.XYPoint{}

		current_stock_price := stock_price_range_start
		for current_stock_price <= stock_price_range_end {
			profit := purchase_price * 100
			if current_stock_price <= contract.StrikePrice {
				profit = purchase_price*100 - (contract.StrikePrice-current_stock_price)*100
			}

			graphData = append(graphData, models.XYPoint{
				X: current_stock_price,
				Y: roundToHundredth(profit),
			})

			current_stock_price += stockPriceRangeInterval
		}

		plData := models.PLData{
			MaxProfit:      roundToHundredth(purchase_price * 100),
			MaxLoss:        (contract.StrikePrice - purchase_price) * 100,
			BreakEvenPoint: roundToHundredth(contract.StrikePrice - purchase_price),
		}

		return graphData, plData
	}

	bidGraphData, bidPLData := analyze(contract.Bid)
	askGraphData, askPLData := analyze(contract.Ask)

	return &models.AnalysisResult{
		BidGraphData: bidGraphData,
		BidPLData:    bidPLData,
		AskGraphData: askGraphData,
		AskPLData:    askPLData,
	}
}

// roundToHundredth rounds a float64 to the nearest hundredth
func roundToHundredth(x float64) float64 {
	return math.Round(x*100) / 100
}
