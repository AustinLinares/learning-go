package main

import (
	"fmt"
	"regexp"
)

// PaymentType types and constants.
type PaymentType string

const (
	PickUpPaymentTypeCash PaymentType = "efectivo"
	PickUpPaymentTypeCard PaymentType = "tarjeta"
)

func paymentTypeChecker(paymentMethod string) (result PaymentType) {
	result = PickUpPaymentTypeCash
	cardWasUsed, _ := regexp.MatchString(`(?i)credit.card|(?i)debit.card|(?i)pos`, paymentMethod)
	if cardWasUsed {
		result = PickUpPaymentTypeCard
	}
	fmt.Println(paymentMethod, result)
	return result
}

func main() {
	paymentTypeChecker("credit_card")
	paymentTypeChecker("Credit_card")
	paymentTypeChecker("creDit_card")
	paymentTypeChecker("creDit card")
	paymentTypeChecker("  creDit card ")
	paymentTypeChecker("a  creDit card b")
	fmt.Println(" ")

	paymentTypeChecker("Debit_card")
	paymentTypeChecker("debit_card")
	paymentTypeChecker("debit_card")
	paymentTypeChecker("DebIt_cArd")
	fmt.Println(" ")

	paymentTypeChecker("cash")
	paymentTypeChecker("pos")
}
