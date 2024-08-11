package ordersstruct

import (
	foodservice "storegestserver/pkg/features/food/service"
)

type CreateOrders struct {
	FoodList   []uint    `validate:"required,min=1,dive,gt=0"`
	FoodAmount []float64 `validate:"required,min=1,dive,gt=0"`
	Amount     float64   `validate:"required,gt=0"`
}

type GetOrdersByDate struct {
	Year  uint `validate:"required,min=1970,max=2038"`
	Month uint `validate:"required,min=1,max=12"`
	Day   uint `validate:"required,min=1,max=31"`
}

type GetOrdersByDateRange struct {
	InitYear  uint `validate:"required,min=1970,max=2038"`
	InitMonth uint `validate:"required,min=1,max=12"`
	InitDay   uint `validate:"required,min=1,max=31"`
	EndYear   uint `validate:"required,min=1970,max=2038"`
	EndMonth  uint `validate:"required,min=1,max=12"`
	EndDay    uint `validate:"required,min=1,max=31"`
}

type Statistics struct {
	Products     []foodservice.Food
	Average      float64
	OrdersNumber int
}
