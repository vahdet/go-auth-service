package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// structs
type (
	User struct {
		Id          int64     `json:"id"`
		Name        string    `json:"name" validate:"required"`
		Email       string    `json:"email" validate:"required"`
		Location    string    `json:"location"`
		Language    string    `json:"language"`
		Created     time.Time `json:"created"`
		LastChanged time.Time `json:"lastChanged"`
	}

	CustomClaims struct {
		UserType string `json:"userType"`
		jwt.StandardClaims
	}

	AuthToken struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
	}
)
