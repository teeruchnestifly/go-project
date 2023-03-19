package service_test

import (
	"testing"

	"go-rest-api/services"

	"github.com/stretchr/testify/assert"
)

func TestAmountMoreThanTenThousandExpectMethodIsCREDIT(t *testing.T) {

	paymentMethod := services.CalculatePaymentMethod(20000)

	assert.Equal(t, "CREDIT", paymentMethod)
}

func TestAmountLessThanTenThousandExpectMethodIsDEBIT(t *testing.T) {

	paymentMethod := services.CalculatePaymentMethod(2000)

	assert.Equal(t, "DEBIT", paymentMethod)
}
