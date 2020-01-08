package registry

import (
    "github.com/kurab/go-gin-gorm/interface/controller"
    "github.com/kurab/go-gin-gorm/interface/database"
    "github.com/kurab/go-gin-gorm/usecase"
)

type interactor struct {
    db database.DB
}

type Interactor interface {
    NewApp() controller.App
}

func NewInteractor(db database.DB) Interactor {
    return &interactor{db}
}

func (i *interactor) NewApp() controller.App {
    return i.NewUsersController()
}

func (i *interactor) NewUsersController() controller.UsersController {
    return controller.NewUsersController(i.NewUserInteractor())
}

func (i *interactor) NewUserInteractor() usecase.UserInteractor {
    return usecase.NewUserInteractor(i.NewUsersRepository())
}

func (i *interactor) NewUsersRepository() database.UserRepository {
    return database.NewUserRepository(i.db)
}
