package moneystruct

type CreateMoney struct {
	Amount      int     `validate:"required"`
	Reason      string  `validate:"required"`
	Description *string `validate:"omitempty"`
}

type GetMoneyByDate struct {
	Year  uint `validate:"required,min=1970,max=2038"`
	Month uint `validate:"required,min=1,max=12"`
	Day   uint `validate:"required,min=1,max=31"`
}
