package database

import (
    "github.com/jinzhu/gorm"
)

type DB interface {
    Create(value interface{}) *gorm.DB
    NewRecord(value interface{}) bool
    Find(out interface{}, where ...interface{}) *gorm.DB
    First(out interface{}, where ...interface{}) *gorm.DB
    Where(query interface{}, args ...interface{}) *gorm.DB
}
