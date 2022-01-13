package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

type ProcurementPlan struct {
	ID               uint   `gorm:"primary_key" json:"ID"`
	No               string `gorm:"varchar(99);comment:合同编号" json:"no"`
	Customer         string `gorm:"varchar(99);comment:客户" json:"customer"`
	EmployeeUID      string `gorm:"type:varchar(32);comment:申请人UID;default:(-)" json:"employeeUID"`
	StartDate        XDate  `gorm:"type:date;comment:申请日期" json:"startDate"`
	Type             string `gorm:"varchar(32);comment:类别" json:"type"`
	Product          string `gorm:"varchar(99);comment:物品" json:"product"`
	Specification    string `gorm:"type:varchar(200);comment:规格" json:"specification"`
	UseNumber        int    `gorm:"type:int;comment:使用数量" json:"userNumber"`
	BuyNumber        int    `gorm:"type:int;comment:采购数量" json:"buyNumber"`
	Unit             string `gorm:"type:varchar(50);comment:单位" json:"unit"`
	Description      string `gorm:"type:varchar(200);comment:要求描述/链接" json:"description"`
	Remarks          string `gorm:"type:varchar(600);comment:备注" json:"remarks"`
	BuyDate          XDate  `gorm:"type:date;comment:购买日期" json:"buyDate"`
	ArriveDate       XDate  `gorm:"type:date;comment:到达日期" json:"arriveDate"`
	BuyRemarks       string `gorm:"type:varchar(600);comment:采购备注" json:"buyRemarks"`
	PayDate          XDate  `gorm:"type:date;comment:付款日期" json:"payDate"`
	PayUnit          string `gorm:"type:varchar(50);comment:付款单位" json:"payUnit"`
	PayRemarks       string `gorm:"type:varchar(600);comment:付款备注" json:"payRemarks"`
	WarehouseInDate  XDate  `gorm:"type:date;comment:入库日期" json:"warehouseInDate"`
	WarehouseOutDate XDate  `gorm:"type:date;comment:出库日期" json:"warehouseOutDate"`
	WarehouseRemarks string `gorm:"type:varchar(600);comment:仓库备注" json:"warehouseRemarks"`

	Employee Employee `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
}

func InsertProcurementPlan(procurementPlan *ProcurementPlan) (code int) {
	err = db.Create(&procurementPlan).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateProcurementPlan(procurementPlan *ProcurementPlan) (code int) {
	var maps = make(map[string]interface{})
	maps["customer"] = procurementPlan.Customer
	maps["start_date"] = procurementPlan.StartDate
	maps["type"] = procurementPlan.Type
	maps["product"] = procurementPlan.Product
	maps["specification"] = procurementPlan.Specification
	maps["use_number"] = procurementPlan.UseNumber
	maps["buy_number"] = procurementPlan.BuyNumber
	maps["unit"] = procurementPlan.Unit
	maps["description"] = procurementPlan.Description
	maps["remarks"] = procurementPlan.Remarks
	maps["buy_date"] = procurementPlan.BuyDate
	maps["arrive_date"] = procurementPlan.ArriveDate
	maps["buy_remarks"] = procurementPlan.BuyRemarks
	maps["pay_date"] = procurementPlan.PayDate
	maps["pay_unit"] = procurementPlan.PayUnit
	maps["pay_remarks"] = procurementPlan.PayRemarks
	maps["warehouse_in_date"] = procurementPlan.WarehouseInDate
	maps["warehouse_out_date"] = procurementPlan.WarehouseOutDate
	maps["warehouse_remarks"] = procurementPlan.WarehouseRemarks

	err = db.Model(&ProcurementPlan{}).Where("id = ?", procurementPlan.ID).Updates(maps).Error
	if err != nil {
		code = msg.ERROR
	} else {
		code = msg.SUCCESS
	}
	return
}

func SelectProcurementPlan(uid string) (procurementPlan ProcurementPlan, code int) {
	err = db.Preload("Employee").First(&procurementPlan, "uid = ?", uid).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		code = msg.ERROR
	} else {
		code = msg.SUCCESS
	}
	return
}

func SelectProcurementPlans(pageSize int, pageNo int, procurementPlan *ProcurementPlan) (procurementPlans []ProcurementPlan, code int, total int64) {
	var maps = make(map[string]interface{})
	tdb := db.Where(maps)
	if procurementPlan.No != "" {
		tdb = tdb.Where("no like ?", "%"+procurementPlan.No+"%")
	}
	if procurementPlan.Product != "" {
		tdb = tdb.Where("product like ?", "%"+procurementPlan.Product+"%")
	}
	err = tdb.Preload("Employee").Find(&procurementPlans).Count(&total).
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Error
	if err != nil {
		code = msg.ERROR
	} else {
		code = msg.SUCCESS
	}
	return
}
