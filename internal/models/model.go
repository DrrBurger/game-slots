package models

type ReelConfiguration [][]string

type LineConfiguration struct {
	Line      int        `json:"line"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type PayoutConfiguration struct {
	Symbol string `json:"symbol"`
	Payout []int  `json:"payout"`
}

type LineResult struct {
	Line   int `json:"line"`
	Payout int `json:"payout"`
}

type Result struct {
	Lines []LineResult `json:"lines"`
	Total int          `json:"total"`
}
