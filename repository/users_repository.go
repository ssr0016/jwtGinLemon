package repository

import "gins/model"

type UsersRepository interface {
	Save(users model.Users)
	Update(users model.Users)
	Delete(usersId int)
	FindById(usersId int) (model.Users, error)
	FindAll() []modelUsers
	FindByUsername(username string) (model.Users, error)
}
