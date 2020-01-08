package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	User    string
	Pass    string
	Host    string
	Port    string
	DBName  string
	Option  string
	Connect *gorm.DB
}

func NewMySQL() *DB {
	c := NewConfig()
	return mySQLConnect(&DB{
		User:   c.DB.DBMySQL.User,
		Pass:   c.DB.DBMySQL.Pass,
		Host:   c.DB.DBMySQL.Host,
		Port:   c.DB.DBMySQL.Port,
		DBName: c.DB.DBMySQL.DBName,
	})
}

func mySQLConnect(d *DB) *DB {
	// connect MySQL
	DBMS := "mysql"
	OPTION := "?charset=utf8&parseTime=True"
	CONNECT := d.User + ":" + d.Pass + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.DBName + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	d.Connect = db
	return d
}

func NewPostgreSQL() *DB {
	c := NewConfig()
	return postgreSQLConnect(&DB{
		User:   c.DB.DBPostgreSQL.User,
		Pass:   c.DB.DBPostgreSQL.Pass,
		Host:   c.DB.DBPostgreSQL.Host,
		Port:   c.DB.DBPostgreSQL.Port,
		DBName: c.DB.DBPostgreSQL.DBName,
	})
}

func postgreSQLConnect(d *DB) *DB {
	DBMS := "postgres"
	CONNECT := "host=" + d.Host + " port=" + d.Port + " user=" + d.User + " dbname=" + d.DBName + " password=" + d.Pass + " sslmode=disable"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	d.Connect = db
	return d
}

// CRUD
// Create
func (db *DB) Create(value interface{}) *gorm.DB {
	return db.Connect.Create(value)
}

func (db *DB) NewRecord(value interface{}) bool {
	return db.Connect.NewRecord(value)
}

// Read
func (db *DB) Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.Find(out, where...)
}

func (db *DB) First(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.First(out, where...)
}

func (db *DB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.Connect.Where(query, args...)
}

//Update/Delete(skip)
