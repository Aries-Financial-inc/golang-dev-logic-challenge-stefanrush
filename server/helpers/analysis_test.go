package helpers

import (
	"math"
	"testing"
	"time"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/models"
)

var testContracts = []*models.OptionsContract{
	{
		Type:           models.Call,
		LongShort:      models.Long,
		StrikePrice:    250,
		ExpirationDate: time.Now().AddDate(0, 0, 30),
		Bid:            5.54,
		Ask:            6.61,
	},
	{
		Type:           models.Call,
		LongShort:      models.Short,
		StrikePrice:    250,
		ExpirationDate: time.Now().AddDate(0, 0, 30),
		Bid:            5.54,
		Ask:            6.61,
	},
	{
		Type:           models.Put,
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

var expectedAnalysisResults = []*models.AnalysisResult{
	{
		BidGraphData: []models.XYPoint{
			{X: 245, Y: -554},
			{X: 246, Y: -554},
			{X: 247, Y: -554},
			{X: 248, Y: -554},
			{X: 249, Y: -554},
			{X: 250, Y: -554},
			{X: 251, Y: -454},
			{X: 252, Y: -354},
			{X: 253, Y: -254},
			{X: 254, Y: -154},
			{X: 255, Y: -54},
		},
		BidPLData: models.PLData{
			MaxProfit:      math.MaxFloat64,
			MaxLoss:        554,
			BreakEvenPoint: 255.54,
		},
		AskGraphData: []models.XYPoint{
			{X: 245, Y: -661},
			{X: 246, Y: -661},
			{X: 247, Y: -661},
			{X: 248, Y: -661},
			{X: 249, Y: -661},
			{X: 250, Y: -661},
			{X: 251, Y: -561},
			{X: 252, Y: -461},
			{X: 253, Y: -361},
			{X: 254, Y: -261},
			{X: 255, Y: -161},
		},
		AskPLData: models.PLData{
			MaxProfit:      math.MaxFloat64,
			MaxLoss:        661,
			BreakEvenPoint: 256.61,
		},
	},
	{
		BidGraphData: []models.XYPoint{
			{X: 245, Y: 554},
			{X: 246, Y: 554},
			{X: 247, Y: 554},
			{X: 248, Y: 554},
			{X: 249, Y: 554},
			{X: 250, Y: 554},
			{X: 251, Y: 454},
			{X: 252, Y: 354},
			{X: 253, Y: 254},
			{X: 254, Y: 154},
			{X: 255, Y: 54},
		},
		BidPLData: models.PLData{
			MaxProfit:      554,
			MaxLoss:        math.MaxFloat64,
			BreakEvenPoint: 255.54,
		},
		AskGraphData: []models.XYPoint{
			{X: 245, Y: 661},
			{X: 246, Y: 661},
			{X: 247, Y: 661},
			{X: 248, Y: 661},
			{X: 249, Y: 661},
			{X: 250, Y: 661},
			{X: 251, Y: 561},
			{X: 252, Y: 461},
			{X: 253, Y: 361},
			{X: 254, Y: 261},
			{X: 255, Y: 161},
		},
		AskPLData: models.PLData{
			MaxProfit:      661,
			MaxLoss:        math.MaxFloat64,
			BreakEvenPoint: 256.61,
		},
	},
	{
		BidGraphData: []models.XYPoint{
			{X: 245, Y: -54},
			{X: 246, Y: -154},
			{X: 247, Y: -254},
			{X: 248, Y: -354},
			{X: 249, Y: -454},
			{X: 250, Y: -554},
			{X: 251, Y: -554},
			{X: 252, Y: -554},
			{X: 253, Y: -554},
			{X: 254, Y: -554},
			{X: 255, Y: -554},
		},
		BidPLData: models.PLData{
			MaxProfit:      244.46,
			MaxLoss:        554,
			BreakEvenPoint: 244.46,
		},
		AskGraphData: []models.XYPoint{
			{X: 245, Y: -161},
			{X: 246, Y: -261},
			{X: 247, Y: -361},
			{X: 248, Y: -461},
			{X: 249, Y: -561},
			{X: 250, Y: -661},
			{X: 251, Y: -661},
			{X: 252, Y: -661},
			{X: 253, Y: -661},
			{X: 254, Y: -661},
			{X: 255, Y: -661},
		},
		AskPLData: models.PLData{
			MaxProfit:      243.39,
			MaxLoss:        661,
			BreakEvenPoint: 243.39,
		},
	},
	{
		BidGraphData: []models.XYPoint{
			{X: 245, Y: 54},
			{X: 246, Y: 154},
			{X: 247, Y: 254},
			{X: 248, Y: 354},
			{X: 249, Y: 454},
			{X: 250, Y: 554},
			{X: 251, Y: 554},
			{X: 252, Y: 554},
			{X: 253, Y: 554},
			{X: 254, Y: 554},
			{X: 255, Y: 554},
		},
		BidPLData: models.PLData{
			MaxProfit:      554,
			MaxLoss:        24446,
			BreakEvenPoint: 244.46,
		},
		AskGraphData: []models.XYPoint{
			{X: 245, Y: 161},
			{X: 246, Y: 261},
			{X: 247, Y: 361},
			{X: 248, Y: 461},
			{X: 249, Y: 561},
			{X: 250, Y: 661},
			{X: 251, Y: 661},
			{X: 252, Y: 661},
			{X: 253, Y: 661},
			{X: 254, Y: 661},
			{X: 255, Y: 661},
		},
		AskPLData: models.PLData{
			MaxProfit:      661,
			MaxLoss:        24339,
			BreakEvenPoint: 243.39,
		},
	},
}

func TestAnalyzeContract(t *testing.T) {
	for i, contract := range testContracts {
		result := AnalyzeContract(contract)
		if !analysisResultsEqual(result, expectedAnalysisResults[i]) {
			t.Errorf("AnalyzeContract(%v) = %v; want %v", contract, result, expectedAnalysisResults[i])
		}
	}
}

func analysisResultsEqual(a, b *models.AnalysisResult) bool {
	if !graphDataEqual(a.BidGraphData, b.BidGraphData) {
		return false
	}

	if a.BidPLData != b.BidPLData {
		return false
	}

	if !graphDataEqual(a.AskGraphData, b.AskGraphData) {
		return false
	}

	if a.AskPLData != b.AskPLData {
		return false
	}

	return true
}

func graphDataEqual(a, b []models.XYPoint) bool {
	// Since b will be a subset of a, lookup and compare only the corresponding points from b in a
	for _, bXYPoint := range b {
		aY := findY(bXYPoint.X, a)
		if aY != bXYPoint.Y {
			return false
		}
	}
	return true
}

func findY(x float64, graphData []models.XYPoint) float64 {
	for _, point := range graphData {
		if point.X == x {
			return point.Y
		}
	}
	return math.MaxFloat64
}
