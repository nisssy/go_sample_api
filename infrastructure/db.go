package infrastructure

import (
	"go_sample_api/domain"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 今回はsqliteを使用するのでHost, Username, Passwordは省略可
// posgreに変更する際に使用
type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName:   c.DB.Production.DBName,
	})
}

func NewTestDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Test.Host,
		Username: c.DB.Test.Username,
		Password: c.DB.Test.Password,
		DBName:   c.DB.Test.DBName,
	})
}

func newDB(d *DB) *DB {
	db, err := gorm.Open(sqlite.Open(d.DBName+".db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	// テーブル作成
	db.AutoMigrate(&domain.User{})
	d.Connection = db
	return d
}

func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
