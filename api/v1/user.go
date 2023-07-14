package v1

import (
	"fmt"
	"gin-basic-framework/global"
	"gin-basic-framework/model"
	"gin-basic-framework/model/common/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserApi struct {
}

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
func (userApi *UserApi) CreateUser(c *gin.Context) {
	newUserInfo := make(map[string]interface{})
	newUserInfo["username"] = c.PostForm("username")
	newUserInfo["password"] = c.PostForm("password")
	newUserInfo["email"] = c.PostForm("email")

	_, err := model.AddUser(newUserInfo)
	if err != nil {
		strErr := fmt.Sprint(err)
		global.GLOBAL_LOGGER.Error(strErr)
		response.FailWithMessage(strErr, c)
	} else {
		global.GLOBAL_LOGGER.Info("User has been created successfully")
		response.OkWithMessage("User has been created successfully", c)
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
func (userApi *UserApi) GetUser(c *gin.Context) {
	userID := c.Param("id")
	fmt.Println(userID)
	// 将字符串类型的用户ID转换为uint类型
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		response.FailWithMessage("Invalid user ID", c)

	}
	// 调用GetUserByID方法从数据库中获取用户信息
	user, err := model.GetUserByID(uint(id))
	if err != nil {
		global.GLOBAL_LOGGER.Info("User not found")
		response.FailWithMessage("User not found", c)
	}
	// 返回用户信息给客户端
	response.OkWithDetailed(user, "User has been get", c)
}

func (userApi *UserApi) DeleteUser(c *gin.Context) {

}
