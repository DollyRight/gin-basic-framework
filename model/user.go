package model

import (
	"fmt"
	"gin-basic-framework/global"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func AddUser(data map[string]interface{}) bool {
	fmt.Println(data["username"].(string))
	global.GLOBAL_DB.Create(&User{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Email:    data["email"].(string),
	})
	return true
}

func GetUserByID(id uint) (*User, error) {
	var user User
	result := global.GLOBAL_DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	result := global.GLOBAL_DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

//func EditUser(id int, data map[string]interface{}) bool {
//	db.Save(&User{ID: id, Username: data["username"], Password: data["password"], Email: data["email"]})
//}
