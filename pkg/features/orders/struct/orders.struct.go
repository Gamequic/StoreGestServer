package ordersstruct

type CreateOrders struct {
	FoodList   []uint `validate:"required,min=1,dive,gt=0"`
	FoodAmount []uint `validate:"required,min=1,dive,gt=0"`
	Amount     uint   `validate:"required,gt=0"`
}

type GetOrdersByDate struct {
	Year  uint `validate:"required,min=1970,max=2038"`
	Month uint `validate:"required,min=1,max=12"`
	Day   uint `validate:"required,min=1,max=31"`
}
