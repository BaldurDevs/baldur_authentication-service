package dto

type UserAuthData struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
