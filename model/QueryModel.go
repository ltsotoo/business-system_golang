package model

type ContractQuery struct {
	AreaUID     string `json:"areaUID"`
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
	SourceTypeUID string `json:"sourceTypeUID"`
	SubtypeUID    string `json:"subtypeUID"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
}

type SupplierQuery struct {
	Name    string `json:"name"`
	Linkman string `json:"linkman"`
	Phone   string `json:"phone"`
}

type EmployeeQuery struct {
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	OfficeUID     string `json:"officeUID"`
	DepartmentUID string `json:"departmentUID"`
}

type AreaQuery struct {
	Name       string `json:"name"`
	OfficeName string `json:"officeName"`
}
