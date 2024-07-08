package userstruct

type CreateUser struct {
	Name     string `validate:"required,min=3"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type UpdateUser struct {
	ID       int    `validate:"required"`
	Name     string `validate:"required,min=3"`
	Email    string `validate:"required,email"`
	Password string `validate:"omitempty,min=8"`
}
