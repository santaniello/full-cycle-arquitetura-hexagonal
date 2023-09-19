package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/santaniello/full-cycle-arquitetura-hexagonal/adapters/cli"
	"github.com/santaniello/full-cycle-arquitetura-hexagonal/application"
	mock_application "github.com/santaniello/full-cycle-arquitetura-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRunCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := createFakeProduct()

	productMock := mockProduct(ctrl, product)

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(product.GetName(), product.GetPrice()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	result, err := cli.Run(service, cli.CREATE, "", product.GetName(), product.GetPrice())

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRunEnableProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := createFakeProduct()

	productMock := mockProduct(ctrl, product)

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Get(product.GetID()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product %s has been enabled.", product.GetName())
	result, err := cli.Run(service, cli.ENABLE, product.GetID(), "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRunDisableProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := createFakeProduct()

	productMock := mockProduct(ctrl, product)

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Get(product.GetID()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product %s has been disabled.", product.GetName())
	result, err := cli.Run(service, cli.DISABLE, product.GetID(), "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func TestRunExecDefault(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := createFakeProduct()

	productMock := mockProduct(ctrl, product)

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Get(product.GetID()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	result, err := cli.Run(service, "get", product.GetID(), "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}

func createFakeProduct() application.ProductInterface {
	product := application.Product{}
	product.ID = "abc"
	product.Name = "Product Test"
	product.Status = application.ENABLED
	product.Price = 25.99
	return &product
}

func mockProduct(ctrl *gomock.Controller, product application.ProductInterface) *mock_application.MockProductInterface {
	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(product.GetID()).AnyTimes()
	productMock.EXPECT().GetStatus().Return(product.GetStatus()).AnyTimes()
	productMock.EXPECT().GetPrice().Return(product.GetPrice()).AnyTimes()
	productMock.EXPECT().GetName().Return(product.GetName()).AnyTimes()
	return productMock
}
