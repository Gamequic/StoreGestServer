package foodstruct

type CreateFood struct {
	Name   string `validate:"required,min=3"`
	Amount uint   `validate:"required,gt=0"`
	Photo  string `validate:"omitempty,url"`
	IsKg   bool   `validate:"omitempty"`
}

type UpdateFood struct {
	ID     int    `validate:"required"`
	Name   string `validate:"omitempty,min=3"`
	Amount uint   `validate:"omitempty,gt=0"`
	Photo  string `validate:"omitempty,url"`
	IsKg   bool   `validate:"omitempty"`
}
