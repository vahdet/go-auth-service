package services

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/vahdet/go-auth-service/app/utils"
	"github.com/vahdet/go-auth-service/models"
	reftknstrpb "github.com/vahdet/go-refresh-token-store-redis/proto"
	usrstrpb "github.com/vahdet/go-user-store-redis/proto"
	"golang.org/x/net/context"
	"gopkg.in/go-playground/validator.v9"
)

const (
	JWT_TOKEN_SIGNING_KEY = "AllYourBase"
)

var validate *validator.Validate

type AuthService struct {
	usrStrCli    usrstrpb.UserServiceClient
	refTknStrCli reftknstrpb.TokenServiceClient
}

func NewAuthService(usrStrCli usrstrpb.UserServiceClient, refTknStrCli reftknstrpb.TokenServiceClient) *AuthService {
	return &AuthService{usrStrCli, refTknStrCli}
}

func (s *AuthService) CreateUser(user *models.User) (*models.User, error) {
	// validation of the input
	if err := validate.Struct(user); err != nil {
		valErr := err.(validator.ValidationErrors)
		log.WithFields(log.Fields{
			"username": user.Name,
			"email":    user.Email,
		}).Error(fmt.Sprintf("validation failed: '%#v'", valErr))
		return nil, err
	}

	// grpc call
	createdUserId, err := s.usrStrCli.Create(context.Background(), &usrstrpb.CreateRequest{
		Name:     user.Name,
		Email:    user.Email,
		Language: user.Language,
		Location: user.Location,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"username": user.Name,
			"email":    user.Email,
		}).Error(fmt.Sprintf("creation failed: '%#v'", err))
		return nil, err
	}

	createdUser, err := s.usrStrCli.Get(context.Background(), createdUserId)
	if err != nil {
		log.WithFields(log.Fields{
			"username": user.Name,
			"email":    user.Email,
		}).Error(fmt.Sprintf("getting created user failed: '%#v'", err))
		return nil, err
	}

	return utils.ConvertProtoToModel(createdUser)
}

func (s *AuthService) GetTokens(userId int64) (*models.AuthToken, error) {
	// access token
	claims := models.CustomClaims{
		"member", // userType
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "go-auth-service",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(JWT_TOKEN_SIGNING_KEY))
	if err != nil {
		log.WithFields(log.Fields{
			"userid": userId,
		}).Error(fmt.Sprintf("creating or signing token failed: '%#v'", err))
		return nil, err
	}

	// refresh token
	refrTknProto, err := s.refTknStrCli.Get(context.Background(), &reftknstrpb.UserId{
		Value: userId,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"userid": userId,
		}).Error(fmt.Sprintf("creating refresh token failed: '%#v'", err))
		return nil, err
	}

	return &models.AuthToken{
		Token:        signedToken,
		RefreshToken: refrTknProto.Token,
	}, nil
}

func (s *AuthService) PurgeRefreshToken(userId int64) (int64, error) {
	return 0, nil
}
