package model

import (
	"gin-basic-framework/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" validate:'min=5,max=10'`
	Password string `json:"password" validate:'min=5,max=10'`
	Email    string `json:"email"`
}

func AddUser(data map[string]interface{}) (bool, error) {
	var user User
	ifUserExist := global.GLOBAL_DB.Where("username = ?", data["username"]).First(&user)
	if ifUserExist.Error == nil {
		// User already exists
		return false, errors.New("user already exists")
	}
	// Create a new user
	newUser := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Email:    data["email"].(string),
	}
	// Save the new user to the database
	if err := global.GLOBAL_DB.Create(&newUser).Error; err != nil {
		return false, err
	}

	return true, nil
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
