package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 课题组 Model
type ResearchGroup struct {
	gorm.Model
	CompanyID uint   `gorm:"type:int;comment:公司ID;not null" json:"companyID"`
	Name      string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
}

func SelectResearchGroupsByCompanyID(companyID int) (researchGroups []ResearchGroup, code int) {
	err = db.Where("company_id = ?", companyID).Find(&researchGroups).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return researchGroups, msg.SUCCESS
}
