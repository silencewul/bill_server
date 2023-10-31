package utils

import (
	"testing"
)

//func TestValidateEmail(t *testing.T) {
//	e1 := "luochuanyuewu@qq.com"
//	if !ValidateEmail(e1) {
//		t.Errorf("验证邮箱失败:%s", e1)
//	}
//
//	e2 := "@test"
//	if !ValidateEmail(e2) {
//		t.Errorf("验证邮箱失败:%s", e2)
//	}
//
//}
func TestValidateEmail(t *testing.T) {

	tests := []struct {
		name string
		args string
		want bool
	}{
		{"正常情况", "luochuanyuewu@qq.com", true},
		{"没有用户名的情况", "@test.com", false},
		{"没有域名的情况", "123@", false},
		{"只有@的情况", "@", false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmail(tt.args); got != tt.want {
				t.Errorf("ValidateEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateNickname(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"用户名是手机号的情况", "15330341148", false},
		{"小于4的情况", "123", false},
		{"用户名过长", "fdshdjfbgnsjdhweri1", false},
		{"正常情况", "你好_.·", true},
		{"用户名为空的情况", "", false},
		{"非法字符的情况", "luochuan_j", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateNickname(tt.args); got != tt.want {
				t.Errorf("ValidateNickname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateBankCard(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateBankCard(tt.args.value); got != tt.want {
				t.Errorf("ValidateBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateChineseAndAlphanumeric(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateChineseAndAlphanumeric(tt.args.value); got != tt.want {
				t.Errorf("ValidateChineseAndAlphanumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateDate(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateDate(tt.args.value); got != tt.want {
				t.Errorf("ValidateDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateDateTime(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateDateTime(tt.args.value); got != tt.want {
				t.Errorf("ValidateDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateEmail1(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmail(tt.args.value); got != tt.want {
				t.Errorf("ValidateEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateFloat(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateFloat(tt.args.value); got != tt.want {
				t.Errorf("ValidateFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateIdCard(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateIdCard(tt.args.value); got != tt.want {
				t.Errorf("ValidateIdCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateInteger(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateInteger(tt.args.value); got != tt.want {
				t.Errorf("ValidateInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateMobile(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateMobile(tt.args.value); got != tt.want {
				t.Errorf("ValidateMobile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateNickname1(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateNickname(tt.args.value); got != tt.want {
				t.Errorf("ValidateNickname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePassword(tt.args.value); got != tt.want {
				t.Errorf("ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatePhone(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePhone(tt.args.value); got != tt.want {
				t.Errorf("ValidatePhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateQQ(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateQQ(tt.args.value); got != tt.want {
				t.Errorf("ValidateQQ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateSite(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateSite(tt.args.value); got != tt.want {
				t.Errorf("ValidateSite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateWeiXin(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateWeiXin(tt.args.value); got != tt.want {
				t.Errorf("ValidateWeiXin() = %v, want %v", got, tt.want)
			}
		})
	}
}
