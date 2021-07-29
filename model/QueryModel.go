package model

type CustomerQuery struct {
	AreaID          uint   `json:"areaID"`
	CompanyID       uint   `json:"companyID"`
	ResearchGroupID uint   `json:"researchGroupID"`
	Name            string `json:"name"`
}

type ProductQuery struct {
	SourceType    uint   `json:"sourceType"`
	Subtype       uint   `json:"subtype"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
}
