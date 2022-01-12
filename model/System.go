package model

var SystemSettlement bool

func InitSystem() {
	var systemPath System
	db.First(&systemPath, "text = ?", "SystemSettlement")
	if systemPath.ID == 0 {
		SystemSettlement = false
	} else {
		SystemSettlement = systemPath.Value
	}
}

type System struct {
	ID    uint   `gorm:"primary_key" json:"ID"`
	Text  string `gorm:"type:string;comment:Text" json:"text"`
	Value bool   `gorm:"type:boolean;comment:Value" json:"value"`
}

func ChangeSystemSettlementToTrue() {
	err = db.Model(&System{}).Where("text = ?", "SystemSettlement").Update("value", true).Error
	if err == nil {
		SystemSettlement = true
	}
}

func ChangeSystemSettlementToFalse() {
	err = db.Model(&System{}).Where("text = ?", "SystemSettlement").Update("value", false).Error
	if err == nil {
		SystemSettlement = false
	}
}
