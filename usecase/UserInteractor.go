package usecase

import domain "github.com/kurab/go-gin-gorm/domain/model"

type UserInteractor interface {
    Add(data domain.Users) (err error)
    Users() (users []domain.Users, err error)
    UserById(id int) (user domain.Users, err error)
}

type userInteractor struct {
    User UserRepository
}

func NewUserInteractor(repo UserRepository) UserInteractor {
    return &userInteractor{User: repo}
}

func (interactor *userInteractor) Add(data domain.Users) (err error) {
    err = interactor.User.Store(data)
    if err != nil {
        return err
    }
    return nil
}

func (interactor *userInteractor) Users() (users []domain.Users, err error) {
    users, err = interactor.User.FindAll()
    if err != nil {
        return []domain.Users{}, err
    }
    return users, nil
}

func (interactor *userInteractor) UserById(id int) (user domain.Users, err error) {
    user, err = interactor.User.FindById(id)
    if err != nil {
        return domain.Users{}, err
    }
    return user, nil
}
