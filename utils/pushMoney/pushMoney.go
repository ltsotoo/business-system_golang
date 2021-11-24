package pushMoney

import (
	"business-system_golang/model"
	"math"
	"time"
)

func Calculate(contract *model.Contract) (contractPushMoney model.ContractPushMoney) {

	if contract.PayType == 1 && !contract.IsSpecial {
		contractPushMoney.ContractUID = contract.UID
		contractPushMoney.TaskTotalMoney = task(&contract.Tasks)
		contractPushMoney.Tasks = contract.Tasks
		d1s := contract.EndDeliveryDate.Format("2006-01-02")
		d2s := contract.EndPaymentDate
		d1, err1 := time.Parse("2006-01-02", d1s)
		d2, err2 := time.Parse("2006-01-02", d2s)
		if err1 == nil && err2 == nil {
			contractPushMoney.PaymentDays = payment(d1, d2)
		} else {
			contractPushMoney.PaymentDays = 0
		}
		if contractPushMoney.PaymentDays > 60 {
			temp := (float64(contractPushMoney.PaymentDays - 60)) * 3 / 1000
			if temp > 0.5 {
				temp = 0.5
			}
			contractPushMoney.PaymentMoneys = contractPushMoney.TaskTotalMoney * temp
			contractPushMoney.PaymentMoneys = round(contractPushMoney.PaymentMoneys, 3)
		}
		contractPushMoney.TotalMoney = contractPushMoney.TaskTotalMoney - contractPushMoney.PaymentMoneys
	} else {
		d1s := contract.EndDeliveryDate.Format("2006-01-02")
		d2s := contract.EndPaymentDate
		d1, err1 := time.Parse("2006-01-02", d1s)
		d2, err2 := time.Parse("2006-01-02", d2s)
		if err1 == nil && err2 == nil {
			contractPushMoney.PaymentDays = payment(d1, d2)
		} else {
			contractPushMoney.PaymentDays = 0
		}
		if contractPushMoney.PaymentDays > 60 {
			temp := (float64(contractPushMoney.PaymentDays - 60)) * 3 / 1000
			if temp > 0.5 {
				temp = 0.5
			}
			contractPushMoney.PaymentMoneys = contract.PaymentTotalAmount * temp
			contractPushMoney.PaymentMoneys = round(contractPushMoney.PaymentMoneys, 3)
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
		temp := ((*tasks)[i].Price - (*tasks)[i].StandardPrice) / (*tasks)[i].StandardPrice
		if temp > 0 {
			println(temp)
			temp = (*tasks)[i].Product.Type.PushMoneyPercentages + temp*100*(*tasks)[i].Product.Type.PushMoneyPercentagesUp
			println(temp)
			(*tasks)[i].PushMoney = (*tasks)[i].TotalPrice * temp
			println((*tasks)[i].PushMoney)
			(*tasks)[i].PushMoney = round((*tasks)[i].PushMoney, 3)
		} else if temp < 0 {
			temp = (*tasks)[i].Product.Type.PushMoneyPercentages + temp*(*tasks)[i].Product.Type.PushMoneyPercentagesDown*100
			(*tasks)[i].PushMoney = (*tasks)[i].TotalPrice * temp
			(*tasks)[i].PushMoney = round((*tasks)[i].PushMoney, 3)
		} else {
			(*tasks)[i].PushMoney = (*tasks)[i].TotalPrice * (*tasks)[i].Product.Type.PushMoneyPercentages
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
