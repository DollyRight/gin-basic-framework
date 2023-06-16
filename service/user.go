package service

import (
	"fmt"
	"gin-basic-framework/global"
	"gin-basic-framework/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUser 创建用户接口
// @Summary 创建新用户
// @Description 创建一个新用户
// @Accept json
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param email formData string true "电子邮件"
// @Success 200 {object} gin.H
// @Router /api/v1/user [post]
func CreateUser(c *gin.Context) {
	newUserInfo := make(map[string]interface{})
	newUserInfo["username"] = c.PostForm("username")
	newUserInfo["password"] = c.PostForm("password")
	newUserInfo["email"] = c.PostForm("email")

	status := model.AddUser(newUserInfo)
	if status {
		global.GLOBAL_LOGGER.Info("User has been created successfully")
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "successfully create a new user",
		})
	}

}

// GetUser 获取用户接口
// @Summary 获取用户信息
// @Description 根据用户ID获取用户信息
// @Accept json
// @Produce json
// @Param id path uint true "用户ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /api/v1/user/{id} [get]
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	fmt.Println(userID)
	// 将字符串类型的用户ID转换为uint类型
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Invalid user ID",
		})
		return
	}

	// 调用GetUserByID方法从数据库中获取用户信息
	user, err := model.GetUserByID(uint(id))
	if err != nil {
		global.GLOBAL_LOGGER.Info("User not found")
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "User not found",
		})
		return
	}

	// 返回用户信息给客户端
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}

func DeleteUser(c *gin.Context) {

}
