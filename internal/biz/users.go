package biz

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"team-manager/ent"
	"team-manager/internal/data"

	"github.com/go-kratos/kratos/v2/log"
)

type UsersUsecase struct {
	log       *log.Helper
	usersRepo data.UsersRepo
}

func NewUsersUsecase(
	logger log.Logger,
	usersRepo data.UsersRepo,
) *UsersUsecase {
	return &UsersUsecase{
		log:       log.NewHelper(log.With(logger, "module", "biz/users")),
		usersRepo: usersRepo,
	}
}

func (uc *UsersUsecase) CreateUser(ctx context.Context, tgUser *tgbotapi.User) (*ent.User, error) {
	usersDto := data.UsersDto{
		TgId:         tgUser.ID,
		Username:     tgUser.UserName,
		Name:         tgUser.FirstName,
		Surname:      tgUser.LastName,
		LanguageCode: tgUser.LanguageCode,
		Email:        "",
		Phone:        "",
	}

	return uc.usersRepo.CreateUser(ctx, usersDto)
}
