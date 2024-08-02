package user

type UserCreateDtos struct {
	Username string `json:"username" validate:"required,min=4,max=10"`
	Password string `json:"password" validate:"required,min=8,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,max=100"`
	IsActive bool   `json:"is_active"`
	Roles    int    `json:"roles"`
}
