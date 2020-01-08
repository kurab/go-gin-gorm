package usecase

import domain "github.com/kurab/go-gin-gorm/domain/model"

type UserRepository interface {
    Store(user domain.Users) (err error)
    FindAll() (event []domain.Users, err error)
    FindById(id int) (event domain.Users, err error)
}
