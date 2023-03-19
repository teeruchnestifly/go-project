package services

func CalculatePaymentMethod(amount float64) string {

	if amount >= 10000 {
		return "CREDIT"
	} else {
		return "DEBIT"
	}

}
