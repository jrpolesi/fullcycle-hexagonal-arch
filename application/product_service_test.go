package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jrpolesi/fullcycle-hexagonal-arch/application"
	mock_application "github.com/jrpolesi/fullcycle-hexagonal-arch/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{persistence}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{persistence}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil).AnyTimes()
	product.EXPECT().Enable().Return(nil).AnyTimes()

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(product).Return(product, nil).AnyTimes()

	service := application.ProductService{persistence}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
