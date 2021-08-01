package model

type ContractQuery struct {
	AreaID      uint   `json:"areaID"`
	No          string `json:"no"`
	CompanyName string `json:"companyName"`
}

type CustomerQuery struct {
	AreaID          uint   `json:"areaID"`
	CompanyID       uint   `json:"companyID"`
	ResearchGroupID uint   `json:"researchGroupID"`
	Name            string `json:"name"`
}

type ProductQuery struct {
	SourceTypeID  uint   `json:"sourceTypeID"`
	SubtypeID     uint   `json:"subtypeID"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
}

type SupplierQuery struct {
	Name    string `json:"name"`
	Linkman string `json:"linkman"`
	Phone   string `json:"phone"`
}
