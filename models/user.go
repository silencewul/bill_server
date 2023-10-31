package models

import (
	"bill/modules/constant"
	"bill/modules/log"
	"bill/modules/setting"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRegistPayload struct {
	PhoneNum string `json:"phone"  binding:"required"`
	Pwd      string `json:"password"  binding:"required"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nick_name"`
	Gender   int    `json:"gender"`
	Birthday string `json:"birthday"`
}

type UserLoginPayload struct {
	Input string `json:"input"  binding:"required"`
	Pwd   string `json:"password"  binding:"required"`
}

type UserChangeEmailPayload struct {
	OriginPwd string `json:"originPwd" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

type User struct {
	Id       int    `json:"id" xorm:"not null pk autoincr comment('用户ID') INT(11)"`
	PhoneNum string `json:"phone_num" xorm:"not null comment('用户邮箱') unique VARCHAR(32)"`
	//Email     string `json:"email,omitempty" xorm:"not null comment('用户邮箱') unique VARCHAR(128)"`
	Password string `json:"password,omitempty" xorm:"not null comment('用户密码') VARCHAR(64)"`
	Nickname string `json:"nickname" xorm:"not null comment('昵称') unique VARCHAR(64)"`
	Avatar   string `json:"avatar,omitempty" xorm:"not null default '' comment('头像') VARCHAR(255)"`
	Gender   int    `json:"gender" xorm:"not null comment('性别') INT(11)"`
	//LoginTime int    `json:"login_time,omitempty" xorm:"not null default 0 comment('最后登录时间') INT(11)"`
	//LoginIp   string `json:"login_ip,omitempty" xorm:"not null default '' comment('最后登录IP') VARCHAR(64)"`
	//CreatedIp string `json:"created_ip,omitempty" xorm:"not null default '' comment('注册IP') VARCHAR(64)"`
	CreatedAt time.Time `json:"created_at,omitempty" xorm:"created not null default 0 comment('注册时间') INT(10)"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xorm:"updated not null default 0 comment('最后更新时间') index INT(10)"`
}

func (user *User) HashPassword() {
	var hash []byte
	hash, _ = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	return
}

func (user *User) Get() error {
	ok, err := MasterDB.Get(user)
	if err != nil {
		log.GetSugar().Errorf("获取User %d 出错,sql错误:%s", user.Id, err.Error())
		return constant.ErrServerInternalError
	}

	if !ok {
		return constant.ErrUserNotFound
	}

	return nil
}

func (user *User) Delete() error {
	affected, err := MasterDB.ID(user.Id).Delete(user)
	if err != nil {
		log.GetSugar().Errorf("删除用户: %d 失败,sql错误:%s", err.Error())
		return constant.ErrServerInternalError
	}
	if affected == 0 {
		return constant.ErrDeleteFail
	}
	return nil
}

func (user *User) ExistsInRedis() bool {
	_, err := MasterRedis.Exists(fmt.Sprintf("%s%d", constant.LoginUser, user.Id)).Result()
	return err == nil
}

func (user *User) StoreToRedis() error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		log.GetSugar().Errorf("用户:%d Json编码失败,错误:%s", user.Id, err.Error())
		return err
	}
	loginUserKey := fmt.Sprintf("%s%d", constant.LoginUser, user.Id)

	expire := time.Second * time.Duration(setting.GetConfig().Server.TokenMaxAge)

	if err := MasterRedis.Set(loginUserKey, userBytes, expire).Err(); err != nil {
		log.GetSugar().Errorf("用户:%d 存入Redis失败,错误%s ", user.Id, err.Error())
		return err
	}
	return nil
}

// UserFromRedis 从redis中取出用户信息
func UserFromRedis(userID int) (*User, error) {
	key := fmt.Sprintf("%s%d", constant.LoginUser, userID)

	userBytes, err := MasterRedis.Get(key).Bytes()

	if err != nil {
		return nil, err
	}
	user := &User{}
	bytesErr := json.Unmarshal(userBytes, user)
	if bytesErr != nil {
		return user, err
	}
	return user, nil
}

func UserExistsInRedis(userId interface{}) bool {
	_, err := MasterRedis.Exists(fmt.Sprintf("%s%d", constant.LoginUser, userId)).Result()
	return err == nil
}

// UserToRedis 将用户信息存到redis
func UserToRedis(user *User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		log.GetSugar().Error(err)
		return err
	}
	loginUserKey := fmt.Sprintf("%s%d", constant.LoginUser, user.Id)

	expire := time.Second * time.Duration(setting.GetConfig().Server.TokenMaxAge)

	if err := MasterRedis.Set(loginUserKey, userBytes, expire).Err(); err != nil {
		log.GetSugar().Errorf("Redis存用户失败:%s\n ", err.Error())
		return err
	}
	return nil
}

// 删除缓存在Redis中的用户
func UserDeleteFromRedis(userId int) {
	MasterRedis.Del(fmt.Sprintf("%s%d", constant.LoginUser, userId))
}
