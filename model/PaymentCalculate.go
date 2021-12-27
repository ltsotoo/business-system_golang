package model

import (
	"math"
	"time"
)

func calculate(contract *Contract, payment *Payment) (theoreticalPushMoney float64, fine float64, pushMoney float64, businessMoney float64) {
	var task Task
	for i := range contract.Tasks {
		if contract.Tasks[i].UID == payment.TaskUID {
			task = contract.Tasks[i]
			break
		}
	}

	if task.UID != "" {
		if contract.IsSpecial {
			theoreticalPushMoney, fine, pushMoney = special(contract.PayType, contract.EndDeliveryDate.Time, payment, &task)
		} else {
			theoreticalPushMoney, fine, pushMoney, businessMoney = simple(contract.PayType, contract.EndDeliveryDate.Time, payment, &task)
		}
	}
	return
}

func simple(payType int, endDeliveryDate time.Time, payment *Payment, task *Task) (theoreticalPushMoney float64, fine float64, realPushMoney float64, businessMoney float64) {
	var tempPrice float64
	if payType == 1 {
		tempPrice = task.StandardPrice
	} else {
		tempPrice = task.StandardPriceUSD
	}
	if task.Price >= tempPrice {
		//产品任务标准价总额
		paymentTotalStandardPrice := task.StandardPrice * float64(task.Number)

		if task.PaymentTotalPrice < paymentTotalStandardPrice {
			difference := paymentTotalStandardPrice - task.PaymentTotalPrice
			if difference >= payment.Money {
				//回款没达到标准价格
				theoreticalPushMoney = round(payment.Money*task.Product.Type.PushMoneyPercentages, 3)
			} else {
				//回款达到标准价格计算两部分提成
				//标准提出
				theoreticalPushMoney1 := round(difference*task.Product.Type.PushMoneyPercentages, 3)
				//高出部分提成
				theoreticalPushMoney2 := round((payment.Money-difference)*task.Product.Type.PushMoneyPercentagesUp, 3)
				theoreticalPushMoney = theoreticalPushMoney1 + theoreticalPushMoney2
				businessMoney = round((payment.Money-difference)*task.Product.Type.BusinessMoneyPercentagesUp, 3)
			}
		} else {
			theoreticalPushMoney = round(payment.Money*task.Product.Type.PushMoneyPercentagesUp, 3)
			businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentagesUp, 3)
		}
	} else {
		theoreticalPushMoney = round(payment.Money*task.Product.Type.PushMoneyPercentages, 3)
	}

	//回款延迟扣除
	tdoa := calculateTDOA(endDeliveryDate, payment.PaymentDate.Time)
	if tdoa > 60 {
		rate := float64(tdoa-60) * 2 / 1000
		if rate > 0.4 {
			rate = 0.4
		}
		fine = round(theoreticalPushMoney*rate, 3)
	}

	//标准业务费
	var tempBusinessMoney float64
	if payType == 1 {
		tempBusinessMoney = round(payment.Money/task.TotalPrice*task.Product.Type.BusinessMoneyPercentages, 3)
	} else {
		tempBusinessMoney = round(payment.MoneyUSD/task.TotalPrice*task.Product.Type.BusinessMoneyPercentages, 3)
	}
	//业务费合并
	businessMoney = businessMoney + tempBusinessMoney

	//实际提成
	realPushMoney = theoreticalPushMoney - fine

	return
}

func special(payType int, endDeliveryDate time.Time, payment *Payment, task *Task) (theoreticalPushMoney float64, fine float64, realPushMoney float64) {
	theoreticalPushMoney = round(payment.Money*task.PushMoneyPercentages, 3)

	//回款延迟扣除
	tdoa := calculateTDOA(endDeliveryDate, payment.PaymentDate.Time)
	if tdoa > 60 {
		rate := float64(tdoa-60) * 2 / 1000
		if rate > 0.4 {
			rate = 0.4
		}
		fine = round(theoreticalPushMoney*rate, 3)
	}

	realPushMoney = theoreticalPushMoney - fine

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
