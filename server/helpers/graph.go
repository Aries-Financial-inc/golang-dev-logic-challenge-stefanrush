package helpers

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/models"
)

// PlotLineGraph creates a line graph for the given options contract analysis
func PlotLineGraph(contract *models.OptionsContract) *charts.Line {
	line := charts.NewLine()

	var longShortText string
	switch contract.LongShort {
	case models.Long:
		longShortText = "Long"
	case models.Short:
		longShortText = "Short"
	}

	var typeText string
	switch contract.Type {
	case models.Call:
		typeText = "Call"
	case models.Put:
		typeText = "Put"
	}

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: fmt.Sprintf("%v", contract.ID),
			Subtitle: fmt.Sprintf(
				"%v %v, Strike Price = $%v, Expiration Date = %v\n"+
					"Bid: Price = $%v, Max P = $%v, Max L = $%v, B/E = $%v\n"+
					"Ask: Price = $%v, Max P = $%v, Max L = $%v, B/E = $%v",
				longShortText,
				typeText,
				contract.StrikePrice,
				contract.ExpirationDate.Format("2006-01-02"),
				contract.Bid,
				contract.AnalysisResult.BidPLData.MaxProfit,
				contract.AnalysisResult.BidPLData.MaxLoss,
				contract.AnalysisResult.BidPLData.BreakEvenPoint,
				contract.Ask,
				contract.AnalysisResult.AskPLData.MaxProfit,
				contract.AnalysisResult.AskPLData.MaxLoss,
				contract.AnalysisResult.AskPLData.BreakEvenPoint,
			),
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         "Stock Price",
			NameLocation: "center",
			NameGap:      36,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "P/L",
			NameLocation: "center",
			NameGap:      60,
		}),
		charts.WithGridOpts(opts.Grid{Top: "100"}),
	)

	line.SetXAxis(getXPoints(contract.AnalysisResult.BidGraphData)).
		AddSeries("Bid", getYPoints(contract.AnalysisResult.BidGraphData)).
		AddSeries("Ask", getYPoints(contract.AnalysisResult.AskGraphData))

	return line
}

// getXPoints returns the X points as a slice of float64s from the given XY points
func getXPoints(xyPoints []models.XYPoint) []float64 {
	points := make([]float64, 0, len(xyPoints))
	for _, point := range xyPoints {
		points = append(points, point.X)
	}
	return points
}

// getYPoints returns the Y points as a slice echarts LineData structs from the given XY points
func getYPoints(xyPoints []models.XYPoint) []opts.LineData {
	points := make([]opts.LineData, 0, len(xyPoints))
	for _, point := range xyPoints {
		points = append(points, opts.LineData{Value: point.Y})
	}
	return points
}
