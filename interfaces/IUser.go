package interfaces

import "github.com/Mojtaba-Afshar/ekart/entities"

type IUser interface {
	Login(user *entities.User) string
	Register(user *entities.User) string
	GetProfile(UserID int) *entities.User
	SearchUser(serachQuery string)
	LogOut(userID int) string
}
