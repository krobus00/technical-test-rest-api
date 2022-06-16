package model

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,min=5,max=32"`
	Password string `json:"password" validate:"required,min=5,max=32"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,min=5,max=32"`
	Password string `json:"password" validate:"required,min=5,max=32"`
}

type TokenResponse struct {
	AccessToken           string `json:"accessToken"`
	AccessTokenExpiredAt  int64  `json:"AccessTokenExpiredAt"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiredAt int64  `json:"RefreshTokenExpiredAt"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	DateColumn
}
