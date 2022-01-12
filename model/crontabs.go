package model

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func InitCronTabs() {
	updateAllEmployeesMoney()
	// updateSystemStatus()
}

func updateAllEmployeesMoney() {
	cronTask := cron.New()

	_, err := cronTask.AddFunc("0 0 1 1 *", UpdateAllAddMoney)
	if err != nil {
		fmt.Println(err)
	}

	cronTask.Start()
}

func updateSystemStatus() {
	cronTask := cron.New()

	_, err := cronTask.AddFunc("* * * * *", ChangeSystemSettlementToTrue)
	if err != nil {
		fmt.Println(err)
	}

	cronTask.Start()
}
