package model

type ProductQuery struct {
	SourceType    uint   `json:"sourceType"`
	Subtype       uint   `json:"subtype"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
}
