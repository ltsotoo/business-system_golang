package pushMoney

import (
	"business-system_golang/model"
	"math"
	"time"
)

func Calculate(contract *model.Contract) (contractPushMoneyQuery model.ContractPushMoneyQuery) {

	if contract.PayType == 1 && !contract.IsSpecial {
		contractPushMoneyQuery.ContractUID = contract.UID
		contractPushMoneyQuery.TaskTotalMoney = task(&contract.Tasks)
		contractPushMoneyQuery.Tasks = contract.Tasks
		d1s := contract.EndDeliveryDate.Format("2006-01-02")
		d2s := contract.EndPaymentDate
		d1, err1 := time.Parse("2006-01-02 15:04:05", d1s)
		d2, err2 := time.Parse("2006-01-02 15:04:05", d2s)
		if err1 == nil && err2 == nil {
			contractPushMoneyQuery.PaymentDays = payment(d1, d2)
		} else {
			contractPushMoneyQuery.PaymentDays = 0
		}
		if contractPushMoneyQuery.PaymentDays > 60 {
			temp := (float64(contractPushMoneyQuery.PaymentDays - 60)) * 3 / 1000
			if temp > 0.5 {
				temp = 0.5
			}
			contractPushMoneyQuery.PaymentMoneys = contractPushMoneyQuery.TaskTotalMoney * temp
			contractPushMoneyQuery.PaymentMoneys = round(contractPushMoneyQuery.PaymentMoneys, 3)
		}
		contractPushMoneyQuery.TotalMoney = contractPushMoneyQuery.TaskTotalMoney - contractPushMoneyQuery.PaymentMoneys
	} else {
		d1s := contract.EndDeliveryDate.Format("2006-01-02")
		d2s := contract.EndPaymentDate
		d1, err1 := time.Parse("2006-01-02 15:04:05", d1s)
		d2, err2 := time.Parse("2006-01-02 15:04:05", d2s)
		if err1 == nil && err2 == nil {
			contractPushMoneyQuery.PaymentDays = payment(d1, d2)
		} else {
			contractPushMoneyQuery.PaymentDays = 0
		}
		if contractPushMoneyQuery.PaymentDays > 60 {
			temp := (float64(contractPushMoneyQuery.PaymentDays - 60)) * 3 / 1000
			if temp > 0.5 {
				temp = 0.5
			}
			contractPushMoneyQuery.PaymentMoneys = contract.PaymentTotalAmount * temp
			contractPushMoneyQuery.PaymentMoneys = round(contractPushMoneyQuery.PaymentMoneys, 3)
		}
	}
	return
}

func base(totalMoney float64) (basePushMoney float64) {
	//给予办事处销售额的0.5%做为招待费
	bpm1 := totalMoney * 0.005
	//照回款额的1%给予办事处业务推广会
	bpm2 := totalMoney * 0.01
	//按照销售额的0.5%给予办事处办事处做为会展费
	bpm3 := totalMoney * 0.005

	basePushMoney = bpm1 + bpm2 + bpm3
	return
}

func task(tasks *[]model.Task) (taskTotalMoney float64) {

	for i, _ := range *tasks {
		temp := ((*tasks)[i].Price - (*tasks)[i].Product.StandardPrice) / (*tasks)[i].Product.StandardPrice
		if temp > 0 {
			temp = (*tasks)[i].Product.PushMoneyPercentages + temp*(*tasks)[i].Product.PushMoneyPercentagesUp*100
			(*tasks)[i].PushMoney = (*tasks)[i].TotalPrice * temp
			(*tasks)[i].PushMoney = round((*tasks)[i].PushMoney, 3)
		} else if temp < 0 {
			temp = (*tasks)[i].Product.PushMoneyPercentages + temp*(*tasks)[i].Product.PushMoneyPercentagesDown*100
			(*tasks)[i].PushMoney = (*tasks)[i].TotalPrice * temp
			(*tasks)[i].PushMoney = round((*tasks)[i].PushMoney, 3)
		} else {
			(*tasks)[i].PushMoney = (*tasks)[i].TotalPrice * (*tasks)[i].Product.PushMoneyPercentages
			(*tasks)[i].PushMoney = round((*tasks)[i].PushMoney, 3)
		}
		taskTotalMoney += (*tasks)[i].PushMoney
	}

	return
}

func payment(EndDeliveryDateString time.Time, EndPaymentDate time.Time) (paymentDays int) {

	t1s := EndDeliveryDateString.Unix()
	t2s := EndPaymentDate.Unix()
	if t1s >= t2s {
		return 0
	}

	days := (t2s - t1s) / 86400
	daysX := (t2s - t1s) % 86400
	if daysX > 0 {
		paymentDays = int(days) + 1
	} else {
		paymentDays = int(days)
	}
	return
}

func round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
