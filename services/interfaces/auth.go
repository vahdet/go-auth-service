package interfaces

import "github.com/vahdet/go-auth-service/models"

type (
	AuthService interface {
		// User
		// GetUserByUserName(userName string) (*models.User, error)
		// GetUserByEmail(email string) (*models.User, error)
		// GetUser(userId int64) (*models.User, error)
		CreateUser(user *models.User) (*models.User, error)
		// UpdateUser(userId int64, user *models.User) (*models.User, error)
		// DeleteUser(userId int64) (*models.User, error)

		// Token
		GetTokens(userId int64) (*models.AuthToken, error)
		PurgeRefreshToken(userId int64) (int64, error)
	}
)
