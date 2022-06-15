package model

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	AccessToken           string `json:"accessToken"`
	AccessTokenExpiredAt  int64  `json:"AccessTokenExpiredAt"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiredAt int64  `json:"RefreshTokenExpiredAt"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
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
