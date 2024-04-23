package data

import (
	"context"
	"team-manager/ent/user"

	"team-manager/ent"
)

type UsersDto struct {
	TgId         int64
	Username     string
	Name         string
	Surname      string
	LanguageCode string
	Email        string
	Phone        string
}

type UsersRepo interface {
	GetUserByTgId(ctx context.Context, tgUserId int64) (*ent.User, error)
	CreateUser(ctx context.Context, usersDto UsersDto) (*ent.User, error)
}

type usersRepo struct {
	db *ent.Client
}

func NewUsersRepo(d *Data) UsersRepo {
	return &usersRepo{
		db: d.db,
	}
}

func (r *usersRepo) GetUserByTgId(ctx context.Context, tgId int64) (*ent.User, error) {
	return r.db.User.Query().Where(user.TgIDEQ(tgId)).First(ctx)
}

func (r *usersRepo) CreateUser(ctx context.Context, usersDto UsersDto) (*ent.User, error) {
	return r.db.User.Create().
		SetTgID(usersDto.TgId).
		SetUsername(usersDto.Username).
		SetName(usersDto.Name).
		SetSurname(usersDto.Surname).
		SetEmail(usersDto.Email).
		SetPhone(usersDto.Phone).
		Save(ctx)
}
