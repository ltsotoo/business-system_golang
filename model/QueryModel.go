package model

type ContractQuery struct {
	AreaID      string `json:"areaID"`
	No          string `json:"no"`
	CompanyName string `json:"companyName"`
}

type CustomerQuery struct {
	AreaUID       string `json:"areaUID"`
	CompanyUID    string `json:"companyUID"`
	ResearchGroup string `json:"researchGroup"`
	Name          string `json:"name"`
}

type ProductQuery struct {
	SourceTypeID  string `json:"sourceTypeID"`
	SubtypeID     string `json:"subtypeID"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
}

type SupplierQuery struct {
	Name    string `json:"name"`
	Linkman string `json:"linkman"`
	Phone   string `json:"phone"`
}

type DictionaryQuery struct {
	TypeModule string `json:"typeModule"`
	TypeName   string `json:"typeName"`
	TypeUID    string `json:"typeUID"`
	PUID       string `json:"pUID"`
}
