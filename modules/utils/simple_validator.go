package utils

import (
	"gopkg.in/go-playground/validator.v8"
	"regexp"
	"strconv"
	"unicode/utf8"
)

// 参考 15个常用的js正则表达式 https://www.jb51.net/article/115170.htm
// 想学习正则表达式基础,就看这里: https://www.runoob.com/regexp/regexp-syntax.html

// ValidatePositiveNumber 检查是否是一个有效的正整数
func ValidatePositiveNum(str string) bool {
	if ok, err := regexp.MatchString(`[1-9]\d*`, str); ok || err != nil {
		return false
	}
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	return true
}

// ValidateEmail 检查是否是一个有效的邮箱字符串
func ValidateEmail(value string) bool {

	matched, err := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", value)
	if !matched || err != nil {
		return false
	}
	return true
}

func Validate6Code(value string) bool {
	matched, err := regexp.MatchString("^\\d{6}$", value)
	if !matched || err != nil {
		return false
	}
	return true
}

// ValidateNickname 验证用户名/真名
func ValidateNickname(value string) bool {

	length := (len(value) + utf8.RuneCountInString(value)) / 2

	if length > 18 || length < 4 {
		return false
	}

	//用户名不能是手机号
	if ValidateMobile(value) {
		return false
	}
	/* 中文、英文、数字包括下划线
	 */
	ok, err := regexp.MatchString(`^[\x{4e00}-\x{9fa5}a-zA-z0-9_]+$`, value)
	if !ok || err != nil {
		return false
	}

	return true
}

func ValidatePassword(value string) bool {
	ok, err := regexp.MatchString(`^[\S]{5,20}$`, value)
	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateIdCard 验证身份证
func ValidateIdCard(value string) bool {
	ok, err := regexp.MatchString(`^\d{17}[0-9xX]$`, value)
	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateBankCard 验证银行卡
func ValidateBankCard(value string) bool {
	ok, err := regexp.MatchString(`^(\d{16,19})$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateMobile 验证手机号
func ValidateMobile(value string) bool {

	if len(value) != 11 {
		return false
	}

	ok, err := regexp.MatchString(`^1\d{10}$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

// ValidatePhone 验证座机号
func ValidatePhone(value string) bool {
	ok, err := regexp.MatchString(`^(\d{4}-|\d{3}-)?(\d{8}|\d{7})$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateDate 验证时间
func ValidateDate(value string) bool {
	ok, err := regexp.MatchString(`^(\d{4}|\d{2})-((0?([1-9]))|(1[0-2]))-((0?[1-9])|([12]([0-9]))|(3[0|1]))$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateDateTime 验证时间
func ValidateDateTime(value string) bool {
	ok, err := regexp.MatchString(`^(?:(?!0000)[0-9]{4}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1[0-9]|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[0-9]{2}(?:0[48]|[2468][048]|[13579][26])|(?:0[48]|[2468][048]|[13579][26])00)-02-29)$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateQQ 验证QQ
func ValidateQQ(value string) bool {
	ok, err := regexp.MatchString(`^[1-9]\d{4,}$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateWeiXin 验证微信号
func ValidateWeiXin(value string) bool {
	ok, err := regexp.MatchString(`^[a-zA-Z]{1}[-_a-zA-Z0-9]{5,19}$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateInteger 验证整数
func ValidateInteger(value string) bool {
	ok, err := regexp.MatchString(`^[+-]?\d{1,9}$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateFloat 验证浮点数
func ValidateFloat(value string) bool {
	ok, err := regexp.MatchString(`(?i)^(([+-]?[1-9]{1}\d*)|([+-]?[0]{1}))(\.(\d){1,2})?$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateSite 验证网址
func ValidateSite(value string) bool {
	ok, err := regexp.MatchString(`^(http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

// ValidateChineseAndAlphanumeric 验证中文和alpha数字
func ValidateChineseAndAlphanumeric(value string) bool {
	ok, err := regexp.MatchString(`^([\x{4e00}-\x{9fa5}]|[a-zA-Z0-9_.·])*$`, value)

	if !ok || err != nil {
		return false
	}

	return true
}

func newValidator() *validator.Validate {
	c := &validator.Config{
		TagName: "validate",
	}

	v := validator.New(c)

	return v
}
