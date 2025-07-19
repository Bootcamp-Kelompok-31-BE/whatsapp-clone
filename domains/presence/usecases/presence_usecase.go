package usecases

import (
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/infrastructures/redis"
)

type PresenceUsecase interface {
	SetOnline(userID string) error
	SetOffline(userID string) error
	GetStatus(userID string) (string, error)
	GetLastSeen(userID string) (string, error)
}

type presenceUsecase struct{}

func NewPresenceUsecase() PresenceUsecase {
	return &presenceUsecase{}
}

func (u *presenceUsecase) SetOnline(userID string) error {
	return redis.SetUserOnline(userID)
}

func (u *presenceUsecase) SetOffline(userID string) error {
	return redis.SetUserOffline(userID)
}

func (u *presenceUsecase) GetStatus(userID string) (string, error) {
	return redis.GetUserStatus(userID)
}

func (u *presenceUsecase) GetLastSeen(userID string) (string, error) {
	return redis.GetLastSeen(userID)
}
