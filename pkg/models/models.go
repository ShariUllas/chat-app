package models

type (
	SignUpRequest struct {
		Name     string `json:"name" binding:"omitempty"`
		Email    string `json:"email" binding:"omitempty"`
		Password string `json:"password" binding:"omitempty"`
	}
	LoginRequest struct {
		Email    string `json:"email" binding:"omitempty"`
		Password string `json:"password" binding:"omitempty"`
	}
)
