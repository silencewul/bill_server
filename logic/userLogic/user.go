package userLogic

import (
	"bill/models"
	"bill/modules/constant"
	"bill/modules/log"
	"bill/modules/utils"
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
