package database

import (
    domain "github.com/kurab/go-gin-gorm/domain/model"
    "errors"
)

type UserRepository interface {
    Store(data domain.Users) (err error)
    FindAll() (users []domain.Users, err error)
    FindById(id int) (user domain.Users, err error)
}

type userRepository struct {
    DB DB
}

func NewUserRepository(db DB) UserRepository {
    return &userRepository{db}
}

func (repo *userRepository) Store(data domain.Users) (err error) {
    repo.DB.NewRecord(data)
    repo.DB.Create(&data)
    return err
}

func (repo *userRepository) FindAll() (users []domain.Users, err error) {
    users = []domain.Users{}
    repo.DB.Find(&users)
    return users, err
}

func (repo *userRepository) FindById(id int) (user domain.Users, err error) {
    user = domain.Users{}
    repo.DB.Where("ID = ?", id).First(&user)
    if user.ID <= 0 {
        return domain.Users{}, errors.New("User, not found")
    }
    return user, nil
}
