package handlers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IBaseHandler interface {
	GetAll() any
	Get(id uint) any
	DeleteAll() *gorm.DB
	Delete(id uint) *gorm.DB
}

type BaseHandler struct {
	IBaseHandler
	objects    any
	Connection *gorm.DB
}

func (b BaseHandler) Get(id uint) any {
	_ = b.Connection.First(&b.objects, id)
	return b.objects
}

func (b BaseHandler) GetAll() any {
	_ = b.Connection.Find(&b.objects)
	return b.objects
}

/*
DeleteAll Function to delete all entries from a single table
*/
func (b BaseHandler) DeleteAll() *gorm.DB {
	results := b.Connection.Delete(&b.objects)
	return results
}

/*
Delete Function to delete a single instance out of the database
*/
func (b BaseHandler) Delete(id uint) *gorm.DB {
	results := b.Connection.Delete(&b.objects, id)
	return results
}

/*
NewBaseHandler Function to create a new BaseHandler with all the necessary parameters
*/
func NewBaseHandler[T any](objects []T) BaseHandler {
	dsn := "host=localhost user=postgres password=password dbname=thingyDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dbConn, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return BaseHandler{Connection: dbConn, objects: objects}
}
