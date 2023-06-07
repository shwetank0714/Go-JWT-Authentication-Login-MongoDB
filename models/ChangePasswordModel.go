package model

type ChangePasswordModel struct {
	Email           string `json:"email"`
	CurrentPassword string `json:"current_password"`
	NewPassWord     string `json:"new_password"`
}
