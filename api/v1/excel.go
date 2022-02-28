package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func UploadExcelToProcurementPlan(c *gin.Context) {

	no := c.Query("no")
	customer := c.Query("customer")
	startDate := c.Query("startDate")

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		println(err.Error())
		msg.Message(c, msg.ERROR, nil)
	}
	excel, err := excelize.OpenReader(file)
	if err != nil {
		println(err.Error())
		msg.Message(c, msg.ERROR, nil)
	}
	rows, err := excel.GetRows("Sheet1")
	if err != nil {
		println(err.Error())
		msg.Message(c, msg.ERROR, nil)
	}
	rowsLength := len(rows)
	procurementPlans := make([]model.ProcurementPlan, rowsLength)

	for i, row := range rows {
		var procurementPlan model.ProcurementPlan
		procurementPlan.No = no
		procurementPlan.Customer = customer
		procurementPlan.EmployeeUID = c.MustGet("employeeUID").(string)
		if startDate != "" {
			procurementPlan.StartDate.Time, _ = time.Parse("2006-01-02", startDate)
		}
		rowLength := len(row)
		switch {
		case rowLength == 2:
			procurementPlan.Type = row[1]
		case rowLength == 3:
			procurementPlan.Type = row[1]
			procurementPlan.Product = row[2]
		case rowLength == 4:
			procurementPlan.Type = row[1]
			procurementPlan.Product = row[2]
			row3 := row[3]
			if len(row[3]) > 200 {
				row3 = row3[:200]
			}
			procurementPlan.Specification = row3
		case rowLength == 5:
			procurementPlan.Type = row[1]
			procurementPlan.Product = row[2]
			row3 := row[3]
			if len(row[3]) > 200 {
				row3 = row3[:200]
			}
			procurementPlan.Specification = row3
			row4, row4Err := strconv.Atoi(row[4])
			if row4Err != nil {
				row4 = 0
			}
			procurementPlan.UseNumber = row4
		case rowLength == 6:
			procurementPlan.Type = row[1]
			procurementPlan.Product = row[2]
			row3 := row[3]
			if len(row[3]) > 200 {
				row3 = row3[:200]
			}
			procurementPlan.Specification = row3
			row4, row4Err := strconv.Atoi(row[4])
			if row4Err != nil {
				row4 = 0
			}
			procurementPlan.UseNumber = row4
			row5, row5Err := strconv.Atoi(row[5])
			if row5Err != nil {
				row5 = 0
			}
			procurementPlan.BuyNumber = row5
		case rowLength == 7:
			procurementPlan.Type = row[1]
			procurementPlan.Product = row[2]
			row3 := row[3]
			if len(row[3]) > 200 {
				row3 = row3[:200]
			}
			procurementPlan.Specification = row3
			row4, row4Err := strconv.Atoi(row[4])
			if row4Err != nil {
				row4 = 0
			}
			procurementPlan.UseNumber = row4
			row5, row5Err := strconv.Atoi(row[5])
			if row5Err != nil {
				row5 = 0
			}
			procurementPlan.BuyNumber = row5
			procurementPlan.Unit = row[6]
		case rowLength == 8:
			procurementPlan.Type = row[1]
			procurementPlan.Product = row[2]
			row3 := row[3]
			if len(row[3]) > 200 {
				row3 = row3[:200]
			}
			procurementPlan.Specification = row3
			row4, row4Err := strconv.Atoi(row[4])
			if row4Err != nil {
				row4 = 0
			}
			procurementPlan.UseNumber = row4
			row5, row5Err := strconv.Atoi(row[5])
			if row5Err != nil {
				row5 = 0
			}
			procurementPlan.BuyNumber = row5
			procurementPlan.Unit = row[6]
			row7 := row[7]
			if len(row[7]) > 200 {
				row7 = row7[:200]
			}
			procurementPlan.Description = row7
		}
		procurementPlans[i] = procurementPlan
	}
	code = model.InsertProcurementPlans(procurementPlans)
	msg.Message(c, code, nil)
}
