package moneystruct

type CreateMoney struct {
	Amount      float64 `validate:"required"`
	Reason      string  `validate:"required"`
	Description *string `validate:"omitempty"`
}

type GetMoneyByDate struct {
	Year  uint `validate:"required,min=1970,max=2038"`
	Month uint `validate:"required,min=1,max=12"`
	Day   uint `validate:"required,min=1,max=31"`
}

type GetMoneyByDateRange struct {
	InitYear  uint `validate:"required,min=1970,max=2038"`
	InitMonth uint `validate:"required,min=1,max=12"`
	InitDay   uint `validate:"required,min=1,max=31"`
	EndYear   uint `validate:"required,min=1970,max=2038"`
	EndMonth  uint `validate:"required,min=1,max=12"`
	EndDay    uint `validate:"required,min=1,max=31"`
}
