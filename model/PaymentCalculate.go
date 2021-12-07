package model

import (
	"math"
	"time"
)

func calculate(contract *Contract, payment *Payment) (theoreticalPushMoney float64, fine float64, pushMoney float64) {
	if contract.PayType == 1 {
		//人民币合同
		var task Task
		for i := range contract.Tasks {
			if contract.Tasks[i].UID == payment.TaskUID {
				task = contract.Tasks[i]
				break
			}
		}

		if task.UID != "" {
			if contract.IsSpecial {
				theoreticalPushMoney, fine, pushMoney = calculate_cny_sp(contract.EndDeliveryDate.Time, payment.PaymentDate.Time, &task)
			} else {
				theoreticalPushMoney, fine, pushMoney = calculate_cny(contract.EndDeliveryDate.Time, payment.PaymentDate.Time, &task)
			}
		}
	} else if contract.PayType == 2 {
		//外币合同
		if contract.IsSpecial {
			theoreticalPushMoney, fine, pushMoney = calculate_usd_sp(contract.EndDeliveryDate.Time, payment.PaymentDate.Time, &contract.Tasks, contract.TotalAmount)
		} else {
			theoreticalPushMoney, fine, pushMoney = calculate_usd(contract.EndDeliveryDate.Time, payment.PaymentDate.Time, &contract.Tasks, contract.TotalAmount)
		}
	}
	return
}

func calculate_cny(endDeliveryDate time.Time, paymentDate time.Time, task *Task) (theoreticalPushMoney float64, fine float64, realPushMoney float64) {

	var rate float64
	//理论提成
	rate = (task.Price - task.StandardPrice) / task.StandardPrice
	if rate >= 0 {
		rate = task.Product.Type.PushMoneyPercentages + rate*100*task.Product.Type.PushMoneyPercentagesUp
	} else {
		rate = task.Product.Type.PushMoneyPercentages + rate*100*task.Product.Type.PushMoneyPercentagesUp
	}
	theoreticalPushMoney = round(task.TotalPrice*rate, 3)
	//回款延迟扣除
	tdoa := calculateTDOA(endDeliveryDate, paymentDate)
	if tdoa > 60 {
		rate = float64(tdoa-60) * 3 / 1000
		if rate > 0.5 {
			rate = 0.5
		}
		fine = round(theoreticalPushMoney*rate, 3)
	}
	realPushMoney = round(theoreticalPushMoney-fine, 3)
	return
}

func calculate_cny_sp(endDeliveryDate time.Time, paymentDate time.Time, task *Task) (theoreticalPushMoney float64, fine float64, realPushMoney float64) {

	var rate float64
	//理论提成
	theoreticalPushMoney = round(task.TotalPrice*task.PushMoneyPercentages, 3)
	//回款延迟扣除
	tdoa := calculateTDOA(endDeliveryDate, paymentDate)
	if tdoa > 60 {
		rate = float64(tdoa-60) * 3 / 1000
		if rate > 0.5 {
			rate = 0.5
		}
		fine = round(theoreticalPushMoney*rate, 3)
	}
	realPushMoney = round(theoreticalPushMoney-fine, 3)
	return
}

func calculate_usd(endDeliveryDate time.Time, paymentDate time.Time, tasks *[]Task, totalAmount float64) (theoreticalPushMoney float64, fine float64, realPushMoney float64) {
	//理论提成
	scales := make([]float64, len(*tasks))

	for i, task := range *tasks {
		if i == (len(scales) - 1) {
			scales[i] = 1
			for j := 0; j < (len(scales) - 1); j++ {
				scales[i] = scales[i] - scales[j]
			}
			scales[i] = round(scales[i], 3)
		} else {
			scales[i] = round(task.TotalPrice/totalAmount, 3)
		}

		theoreticalPushMoney = theoreticalPushMoney + task.TotalPrice*task.PushMoneyPercentages*scales[i]
	}
	theoreticalPushMoney = round(theoreticalPushMoney, 3)
	//回款延迟扣除
	var rate float64
	tdoa := calculateTDOA(endDeliveryDate, paymentDate)
	if tdoa > 60 {
		rate = float64(tdoa-60) * 3 / 1000
		if rate > 0.5 {
			rate = 0.5
		}
		fine = round(theoreticalPushMoney*rate, 3)
	}
	realPushMoney = round(theoreticalPushMoney-fine, 3)
	return
}
func calculate_usd_sp(endDeliveryDate time.Time, paymentDate time.Time, tasks *[]Task, totalAmount float64) (theoreticalPushMoney float64, fine float64, realPushMoney float64) {
	//理论提成
	scales := make([]float64, len(*tasks))

	for i, task := range *tasks {
		if i == (len(scales) - 1) {
			scales[i] = 1
			for j := 0; j < (len(scales) - 1); j++ {
				scales[i] = scales[i] - scales[j]
			}
			scales[i] = round(scales[i], 3)
		} else {
			scales[i] = round(task.TotalPrice/totalAmount, 3)
		}

		var rate float64
		rate = (task.Price - task.StandardPriceUSD) / task.StandardPriceUSD
		if rate >= 0 {
			rate = task.Product.Type.PushMoneyPercentages + rate*100*task.Product.Type.PushMoneyPercentagesUp
		} else {
			rate = task.Product.Type.PushMoneyPercentages + rate*100*task.Product.Type.PushMoneyPercentagesUp
		}

		theoreticalPushMoney = theoreticalPushMoney + task.TotalPrice*rate*scales[i]
	}
	theoreticalPushMoney = round(theoreticalPushMoney, 3)
	//回款延迟扣除
	var rate float64
	tdoa := calculateTDOA(endDeliveryDate, paymentDate)
	if tdoa > 60 {
		rate = float64(tdoa-60) * 3 / 1000
		if rate > 0.5 {
			rate = 0.5
		}
		fine = round(theoreticalPushMoney*rate, 3)
	}
	realPushMoney = round(theoreticalPushMoney-fine, 3)
	return
}

func calculateTDOA(EndDeliveryDateString time.Time, paymentDate time.Time) (tdoa int) {

	if EndDeliveryDateString.IsZero() {
		return 0
	}

	t1x := EndDeliveryDateString.Unix()
	t2x := paymentDate.Unix()

	if t1x >= t2x {
		return 0
	}

	days := (t2x - t1x) / 86400
	daysX := (t2x - t1x) % 86400

	if daysX > 0 {
		tdoa = int(days) + 1
	} else {
		tdoa = int(days)
	}

	return
}

func round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
