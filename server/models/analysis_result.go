package models

// An AnalysisResult contains the results of a profit/loss analysis on an options contract
type AnalysisResult struct {
	BidGraphData []XYPoint `json:"bid_graph_data"`
	BidPLData    PLData    `json:"bid_pl_data"`
	AskGraphData []XYPoint `json:"ask_graph_data"`
	AskPLData    PLData    `json:"ask_pl_data"`
}

// A PLData contains the profit/loss data for an options contract
type PLData struct {
	MaxProfit      float64 `json:"max_profit"`
	MaxLoss        float64 `json:"max_loss"`
	BreakEvenPoint float64 `json:"break_even_point"`
}

// An XYPoint contains an x and y coordinate
type XYPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
