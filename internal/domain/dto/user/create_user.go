package user

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Avatar    string `json:"avatar"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname"`
}

type CreateUserResponse struct {
	UserDetailResponse
}
