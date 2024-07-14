package application_test

import (
	"testing"

	"github.com/jrpolesi/fullcycle-hexagonal-arch/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	id := uuid.NewV4().String()
	product.ID = id

	require.Equal(t, id, product.GetID())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	name := "Hello"
	product.Name = name

	require.Equal(t, name, product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	status := application.ENABLED
	product.Status = status

	require.Equal(t, status, product.GetStatus())
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	price := 10.0
	product.Price = price

	require.Equal(t, price, product.GetPrice())
}
