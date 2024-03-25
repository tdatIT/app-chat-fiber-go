package user

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Expired int    `json:"expired"`
}
