package infrastructure

import (
    "github.com/gin-gonic/gin"
    domain "github.com/kurab/go-gin-gorm/domain/model"
    "github.com/kurab/go-gin-gorm/interface/controller"
)

type Routing interface {
    SetRouting(ctrl controller.UsersController)
    Run(port string)
}

type routing struct {
    DB  *DB
    Gin *gin.Engine
}

func NewRouting(db *DB) Routing {
    return &routing{
        DB:  db,
        Gin: gin.Default(),
    }
}

func (r *routing) SetRouting(ctrl controller.UsersController) {
    u := r.Gin.Group("/api/v1/user")
    {
        // Create
        u.POST("/register", func(c *gin.Context) { ctrl.Add(c) })
        // Read
        u.GET("/get", func(c *gin.Context) { ctrl.Users(c) })
        u.GET("/get/:id", func(c *gin.Context) { ctrl.UserById(c) })
        // Update (skip)
        // Delete (skip)
    }
}

func (r *routing) Run(port string) {
    r.Gin.Run(port)

    defer r.DB.Connect.Close()

    r.DB.Connect.Set("gorm:table_options", "ENGIN=InnoDB")
    r.DB.Connect.AutoMigrate(&domain.Users{})
    r.DB.Connect.LogMode(true)
}
