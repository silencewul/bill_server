package userLogic

import (
	"bill/models"
	"bill/modules/constant"
	"bill/modules/log"
	"bill/modules/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/* 登录时用到,根据keyword类型(邮箱,用户名,手机号)来获取用户
 */
func GetUserByLoginField(keyword string) (*models.User, error) {
	user := new(models.User)
	/*	if utils.ValidateEmail(keyword) {
			user.Email = keyword
		} else if utils.ValidateMobile(keyword) {
			user.PhoneNum = keyword
		} else {
			user.Nickname = keyword
		}*/
	if utils.ValidateMobile(keyword) {
		user.PhoneNum = keyword
	}

	if err := user.Get(); err != nil {
		return nil, constant.ErrUserNotFound
	}

	return user, nil
}

// Login 登录成功返回当前登录用户信息以及登录是否成功
func Login(request *models.UserLoginPayload) (*models.User, error) {

	user, err := GetUserByLoginField(request.Input)

	if err != nil {
		log.GetSugar().Errorf("用户登陆失败,请求:%v", request.Input)
		return nil, err
	}

	ok := VerifyPassword(user.Password, request.Pwd)

	if !ok {
		return nil, constant.ErrUserPwdMismatch
	}

	return user, nil
}

/*
校验密码是否正确.
*/
func VerifyPassword(originPassword, password string) bool {
	// 登录验证密码
	err := bcrypt.CompareHashAndPassword([]byte(originPassword), []byte(password))

	return err == nil

}

/* GetMe 获取当前登录的用户信息
* 注意,由于用户是转成json缓存到redis里的,user表里某些字段在json转换时会被忽略掉。
* 所以更新角色时,建议从redis取出用户的id,并通过id从数据库拿到用户后再操作.
 */
func GetMe(c *gin.Context) *models.User {
	val, exists := c.Get("user")
	if !exists {
		return nil
	}

	me, ok := val.(*models.User)

	if !ok {
		return nil
	}

	return me
}