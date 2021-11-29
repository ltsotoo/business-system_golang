package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"
)

type BidBond struct {
	BaseModel
	UID              string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	EmployeeUID      string  `gorm:"type:varchar(32);comment:添加记录财务UID" json:"employee"`
	SalesmanUID      string  `gorm:"type:varchar(32);comment:业务员UID" json:"salesmanUID"`
	FinalEmployeeUID string  `gorm:"type:varchar(32);comment:完成回款财务UID" json:"finalEmployeeUID"`
	Money            float64 `gorm:"type:decimal(20,6);comment:保证金金额(元)" json:"money"`
	Remarks          string  `gorm:"type:varchar(600);comment:备注" json:"remarks"`
	Status           int     `gorm:"type:int;comment:状态(1:待退还 2:完成)" json:"status"`
	IsDelete         bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Employee      Employee `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Salesman      Employee `gorm:"foreignKey:SalesmanUID;references:UID" json:"salesman"`
	FinalEmployee Employee `gorm:"foreignKey:FinalEmployeeUID;references:UID" json:"finalEmployee"`
}

func InsertBidBond(bidBond *BidBond) (code int) {
	bidBond.UID = uid.Generate()
	err = db.Create(&bidBond).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteBidBond(uid string) (code int) {
	err = db.Model(&BidBond{}).Where("uid = ?", uid).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateBidBond(bidBond *BidBond) (code int) {
	var maps = make(map[string]interface{})
	maps["remarks"] = bidBond.Remarks
	maps["money"] = bidBond.Money
	maps["salesman_uid"] = bidBond.SalesmanUID

	err = db.Model(&BidBond{}).Where("uid = ?", bidBond.UID).Updates(maps).Error

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func ApproveBidBond(uid string) (code int) {
	err = db.Model(&BidBond{}).Where("uid = ?", uid).Update("status", 2).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectBidBonds(pageSize int, pageNo int, bidBond *BidBond) (bidBonds []BidBond, code int, total int64) {
	var maps = make(map[string]interface{})
	maps["is_delete"] = false

	if bidBond.Status != 0 {
		maps["status"] = bidBond.Status
	}

	err = db.Where(maps).Find(&bidBonds).Count(&total).
		Preload("Salesman.Office").Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&bidBonds).Error
	if err != nil {
		return bidBonds, msg.ERROR, total
	}
	return bidBonds, msg.SUCCESS, total

}
