package cli

import (
	"fmt"
	"github.com/santaniello/full-cycle-arquitetura-hexagonal/application"
)

const (
	CREATE  = "create"
	ENABLE  = "enable"
	DISABLE = "disable"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	switch action {
	case CREATE:
		return create(service, productName, price)
	case ENABLE:
		return enable(service, productId)
	case DISABLE:
		return disable(service, productId)
	default:
		return execDefault(service, productId)
	}
}

func create(service application.ProductServiceInterface, productName string, price float64) (string, error) {
	var result = ""
	product, err := service.Create(productName, price)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	return result, nil
}

func enable(service application.ProductServiceInterface, productId string) (string, error) {
	var result = ""
	product, err := service.Get(productId)
	if err != nil {
		return result, err
	}
	productEnabled, err := service.Enable(product)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("Product %s has been enabled.", productEnabled.GetName())
	return result, nil
}

func disable(service application.ProductServiceInterface, productId string) (string, error) {
	var result = ""
	product, err := service.Get(productId)
	if err != nil {
		return result, err
	}
	productDisabled, err := service.Disable(product)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("Product %s has been disabled.", productDisabled.GetName())
	return result, nil
}

func execDefault(service application.ProductServiceInterface, productId string) (string, error) {
	var result = ""
	res, err := service.Get(productId)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	return result, nil
}
