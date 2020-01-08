package main

import (
    "github.com/kurab/go-gin-gorm/infrastructure"
    "github.com/kurab/go-gin-gorm/registry"
)

func main() {
    //    db := infrastructure.NewMySQL()
    db := infrastructure.NewPostgreSQL()

    registry := registry.NewInteractor(db)
    ctrl := registry.NewApp()

    r := infrastructure.NewRouting(db)
    r.SetRouting(ctrl)
    r.Run(":8080")
}
