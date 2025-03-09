package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goAPI/src"
	"goAPI/src/adapters"
	"goAPI/src/db/handlers"
	"goAPI/src/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"reflect"
)

type NameGetter interface {
	GetName() string
}

type S1 struct {
	NameGetter
	name string
}

func (s *S1) GetName() string {
	return s.name
}

type S2 struct {
	*S1
}

func mainOld() {
	router := gin.Default()
	router.GET("/albums", src.GetAlbums)
	dsn := "host=localhost user=postgres password=password dbname=thingyDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dbConn, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dbConn.AutoMigrate(&models.User{})

	/*
		var email = "test"
		var user = models.User{Name: "Michael2", Username: "mkaiser2", Email: &email}
		result1 := dbConn.Create(&user)
		fmt.Println(result1)
	*/

	var users []models.User
	result := dbConn.Find(&users)
	fmt.Println(result.RowsAffected)
	//var userList []models.User = result.Scan(&models.User{})
	fmt.Println(users[0].Username)
	fmt.Println(users[0].Name)
	fmt.Println(users[0].ID)
	fmt.Println(users[0].CreatedAt)
	bytes, _ := json.Marshal(users)
	fmt.Fprint(os.Stdout, string(bytes))

	fmt.Println(reflect.TypeOf(handlers.UserHandlerObject.GetAll()))
	fmt.Println(adapters.UserAdapterObject.GetAll())
	fmt.Println(handlers.UserHandlerObject.Get(1))
	fmt.Println(adapters.UserAdapterObject.GetAll())
	//fmt.Println(handlers.GetAll(models.User{}, handlers.UserHandlerObj.Connection))

	//var x = handlers.UserHandlerObj
	//fmt.Println(x.TableName())
	//fmt.Println(x.GetAll())
	//router.Run("localhost:8080")
}

func main() {
	router := src.InitRouter()

	router.Run("localhost:8080")

}
