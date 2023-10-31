package setting

import (
	"testing"
)

func TestAll(t *testing.T) {
	if got := GetSiteSetting(); got == nil {
		t.Error("获取SiteSetting配置失败")
	}
	if got := GetConfig(); got == nil {
		t.Error("获取Config配置失败")
	}
}
