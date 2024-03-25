package data

import (
	"BlogServ/config"
	"BlogServ/models"
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	c := config.FetchConfig()
	//fmt.Println(c)
	db := InitDB(c)
	//db.Create(&models.User{Username: "test", Password: "test"})
	var user models.User
	db.First(&user, 1)
	fmt.Println(user.Username)
}
