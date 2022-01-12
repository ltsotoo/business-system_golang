package model

import (
	"math"
	"time"
)

func calculate(contract *Contract, payment *Payment) (theoreticalPushMoney float64, fine float64, pushMoney float64, businessMoney float64) {
	var task Task = contract.Tasks[0]

	if task.UID != "" {
		if contract.IsSpecial {
			theoreticalPushMoney, fine, pushMoney, businessMoney = special(contract.PayType, contract.EndDeliveryDate.Time, payment, &task)
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
	percent := 0.01
	if tempPrice == 0 || task.Price == tempPrice {
		theoreticalPushMoney = round(payment.Money*task.Product.Type.PushMoneyPercentages*percent, 3)
		businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentages*percent, 3)
	} else if task.Price > tempPrice {
		//产品任务标准总价
		paymentTotalStandardPrice := task.StandardPrice * float64(task.Number)

		//任务总回款小于标准价
		if task.PaymentTotalPrice < paymentTotalStandardPrice {
			difference := paymentTotalStandardPrice - task.PaymentTotalPrice
			if difference >= payment.Money {
				//该次回款没达到标准价格
				theoreticalPushMoney = round(payment.Money*task.Product.Type.PushMoneyPercentages*percent, 3)
				businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentages*percent, 3)
			} else {
				//该次回款达到标准价格计算两部分提成
				//标准提成
				theoreticalPushMoney1 := round(difference*task.Product.Type.PushMoneyPercentages*percent, 3)
				businessMoney1 := round(difference*task.Product.Type.BusinessMoneyPercentages*percent, 3)
				//高出部分提成
				theoreticalPushMoney2 := round((payment.Money-difference)*task.Product.Type.PushMoneyPercentagesUp*percent, 3)
				theoreticalPushMoney = theoreticalPushMoney1 + theoreticalPushMoney2
				businessMoney2 := round((payment.Money-difference)*task.Product.Type.BusinessMoneyPercentagesUp*percent, 3)
				businessMoney = businessMoney1 + businessMoney2
			}
		} else {
			//回款大于或等于标准价
			theoreticalPushMoney = round(payment.Money*task.Product.Type.PushMoneyPercentagesUp*percent, 3)
			businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentagesUp*percent, 3)
		}
	} else if task.Price < tempPrice {
		tempPushMoneyPercentages := task.Product.Type.PushMoneyPercentages - (tempPrice-task.Price)/tempPrice*100*task.Product.Type.PushMoneyPercentagesDown
		if tempPushMoneyPercentages < task.Product.Type.MinPushMoneyPercentages {
			tempPushMoneyPercentages = task.Product.Type.MinPushMoneyPercentages
		}
		theoreticalPushMoney = round(payment.Money*tempPushMoneyPercentages*percent, 3)
		businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentages*percent, 3)
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

	//实际提成
	realPushMoney = theoreticalPushMoney - fine

	return
}

func special(payType int, endDeliveryDate time.Time, payment *Payment, task *Task) (theoreticalPushMoney float64, fine float64, realPushMoney float64, businessMoney float64) {
	percent := 0.01
	theoreticalPushMoney = round(payment.Money*task.PushMoneyPercentages*percent, 3)
	businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentages*percent, 3)
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
