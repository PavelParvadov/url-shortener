package di

import "url/internal/user"

type IStatRepository interface {
	AddClick(linkId uint)
}
type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	FindUserByEmail(email string) (*user.User, error)
}
