package pushMoney

// func Calculate(contract *model.Contract) (contractPushMoney model.ContractPushMoney, code int) {

// }

// func base(totalMoney float64) (basePushMoney float64) {
// 	//给予办事处销售额的0.5%做为招待费
// 	bpm1 := totalMoney * 0.005
// 	//照回款额的1%给予办事处业务推广会
// 	bpm2 := totalMoney * 0.01
// 	//按照销售额的0.5%给予办事处办事处做为会展费
// 	bpm3 := totalMoney * 0.005

// 	basePushMoney = bpm1 + bpm2 + bpm3
// 	return
// }

// func task(tasks *[]model.Task) (taskPushMoney float64) {

// 	for _, task := range *tasks {
// 		temp := (task.Price - task.Product.StandardPrice) / task.Price
// 		if temp > 0 {
// 			temp = task.Product.PushMoneyPercentages + temp*task.Product.PushMoneyPercentagesUp
// 			task.PushMoney = task.TotalPrice * temp
// 		} else if temp < 0 {
// 			temp = -temp
// 			temp = task.Product.PushMoneyPercentages - temp*task.Product.PushMoneyPercentagesDown
// 			task.PushMoney = task.TotalPrice * temp
// 		} else {
// 			task.PushMoney = task.TotalPrice * task.Product.PushMoneyPercentages
// 		}
// 		taskPushMoney += task.PushMoney
// 	}

// 	return
// }
