package services

import (
	"github.com/Mojtaba-Afshar/ekart/entities"
	"github.com/Mojtaba-Afshar/ekart/interfaces"
)

type UserService struct {
}

func InitUserServic() interfaces.IUser {
	return &UserService{}
}

func (u *UserService) Login(user *entities.User) string {
	return "user Loggedin"
}
func (u *UserService) Register(user *entities.User) string {
	return "user registred"
}
func (u *UserService) GetProfile(UserID int) *entities.User {
	return &entities.User{
		Name:     "MJ",
		userID:   123,
		Password: "Mj123",
		Email:    "mj@gmail.com"}
}
func (u *UserService) SearchUser(query string) {

}
func (u *UserService) LogOut(UserID int) string {
	return "user logged out"
}
