package controller

import (
    domain "github.com/kurab/go-gin-gorm/domain/model"
    "github.com/kurab/go-gin-gorm/usecase"
    "strconv"
    "time"
)

type UsersController interface {
    Add(c Context)
    Users(c Context)
    UserById(c Context)
}

type usersController struct {
    Interactor usecase.UserInteractor
}

func NewUsersController(ui usecase.UserInteractor) UsersController {
    return &usersController{ui}
}

func (ctrl *usersController) Add(c Context) {
    data := domain.Users{}
    data.CreatedAt = time.Now()
    if err := c.BindJSON(&data); err != nil {
        c.String(500, "Bad request: "+err.Error())
    }

    if err := ctrl.Interactor.Add(data); err != nil {
        c.JSON(409, err)
    }
    c.JSON(201, data)
}

func (ctrl *usersController) Users(c Context) {
    users, err := ctrl.Interactor.Users()
    if err != nil {
        c.JSON(404, err)
        return
    }
    c.JSON(200, users)
}

func (ctrl *usersController) UserById(c Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    user, err := ctrl.Interactor.UserById(id)
    if err != nil {
        c.JSON(404, err)
        return
    }
    c.JSON(200, user)
}
