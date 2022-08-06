package service

import (
	"errors"

	"github.com/isther/easyTrader/model"
	"github.com/isther/easyTrader/pkg/jwt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(user *model.User) error {
	if model.UsernameExist(user.Username) {
		return errors.New("username exist")
	}
	return model.NewUser(user)
}

func (u *UserService) Login(user *model.User) (string, error) {
	tmpUser, err := model.FindUserByUsername(user.Username)
	if err != nil {
		return "", err
	}

	if user.Password != tmpUser.Password {
		return "", errors.New("error password")
	}

	token, err := jwt.BuildMapClaimsJwt(user.Username, user.Password)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserService) SetPassword(user *model.User, password string) error {
	return user.SetPassword(password)
}

func (u *UserService) SetBinanceApiKey(user *model.User, apiKey string) error {
	return user.SetBinanceApiKey(apiKey)
}

func (u *UserService) SetBinanceSecretKey(user *model.User, secretKey string) error {
	return user.SetBinanceSecretKey(secretKey)
}

func (u *UserService) SetDingDingTalkApiKey(user *model.User, apiKey string) error {
	return user.SetDingDingTalkApiKey(apiKey)
}

func (u *UserService) SetDingDingTalkSecretKey(user *model.User, secretKey string) error {
	return user.SetDingDingTalkSecretKey(secretKey)
}
