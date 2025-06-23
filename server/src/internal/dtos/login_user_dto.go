package dtos

type LoginUserDTO struct {
	Email    string `json:"email" validated:"required,email"`
	Password string `json:"password" validated:"required"`
}
