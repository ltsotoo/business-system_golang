package middleware

import (
	"business-system_golang/model"
	"fmt"

	"github.com/robfig/cron/v3"
)

func UpdateAllEmployeesMoney() {
	cronTask := cron.New()

	_, err := cronTask.AddFunc("0 0 1 * *", model.UpdateAllAddMoney)
	if err != nil {
		fmt.Println(err)
	}

	cronTask.Start()
}
