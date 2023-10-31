package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

type Config struct {
	Server      ServerConfig            `yaml:"server"`
	Mysql       MysqlConfig             `yaml:"mysql"`
	Redis       RedisConfig             `yaml:"redis"`
	Xorm        XormConfig              `yaml:"xorm"`
	Jwt         JwtConfig               `yaml:"jwt"`
	Smtp        SMTPConfig              `yaml:"smtp"`
	Sms         SMSConfig               `yaml:"sms"`
	Vod         VodConfig               `yaml:"vod"`
	Oauth2      map[string]OAuth2Config `yaml:"oauth2"`
	Oss         OssConfig               `yaml:"oss"`
	Permissions map[string][]string     `yaml:"permissions"`
}

type ServerConfig struct {
	Host        string `yaml:"server"`
	Port        string `yaml:"port"`
	TokenMaxAge int    `yaml:"tokenMaxAge"`
	Production  bool   `yaml:"production"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

type RedisConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Password  string `yaml:"password"`
	Prefix    string `yaml:"prefix"`
	MaxIdle   string `yaml:"maxIdle"`
	MaxActive string `yaml:"maxActive"`
}

type XormConfig struct {
	ShowSql  bool   `yaml:"showSql"`
	LogLevel int    `yaml:"logLevel"`
	Timezone string `yaml:"timezone"`
}

type JwtConfig struct {
	Secret  string `yaml:"secret"`
	Timeout int    `yaml:"timeout"`
}

type SMTPConfig struct {
	Enable     bool   `yaml:"enable"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	Sender     string `yaml:"sender"`
	SenderName string `yaml:"senderName"`
}

type SMSConfig struct {
	Enable       bool   `yaml:"enable"`
	RegionId     string `yaml:"regionId"`
	TemplateCode string `yaml:"TemplateCode"`
	Key          string `yaml:"key"`
	Secret       string `yaml:"secret"`
}

type VodConfig struct {
	Enabled         bool   `yaml:"enable"`
	Key             string `yaml:"key"`
	Secret          string `yaml:"secret"`
	RegionId        string `yaml:"regionId"`
	TemplateGroupId string `yaml:"templateGroupId"`
	PrivateKey      string `yaml:"privateKey"`
	CallbackUrl     string `yaml:"callbackUrl"`
	DevCallbackUrl  string `yaml:"devCallbackUrl"`
}

type OAuth2Config struct {
	Enable       bool   `yaml:"enable"`
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
}

type OssConfig struct {
	Endpoint        string `yaml:"endpoint"`
	CdnDomain       string `yaml:"cdnDomain"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
	CallbackUrl     string `yaml:"callbackUrl"`
	DevCallbackUrl  string `yaml:"devCallbackUrl"`
	ExpireTime      int    `yaml:"expireTime"`
}

type SiteSetting struct {
	Version     string `json:"version" yaml:"version" binding:"required"`
	Title       string `json:"title" yaml:"title" binding:"required"`
	SubTitle    string `json:"subTitle" yaml:"subTitle" binding:"required"`
	Domain      string `json:"domain" yaml:"domain" binding:"required"`
	DevDomain   string `json:"devDomain" yaml:"devDomain" binding:"required"`
	Logo        string `json:"logo" yaml:"logo" binding:"required"`
	Favicon     string `json:"favicon" yaml:"favicon" binding:"required"`
	SeoKeywords string `json:"seoKeywords" yaml:"seoKeywords" binding:"required"`
	SeoDesc     string `json:"seoDesc" yaml:"seoDesc" binding:"required"`
	Copyright   string `json:"copyright" yaml:"copyright" binding:"required"`
	Beian       string `json:"beian" yaml:"beian" binding:"required"`
	BeianLink   string `json:"beianLink" yaml:"beianLink" binding:"required"`
	BaiduStat   string `json:"baiduStat" yaml:"baiduStat" binding:"required"`
	PointName   string `json:"pointName" yaml:"pointName" binding:"required"`
	CoinName    string `json:"coinName" yaml:"coinName" binding:"required"`

	GlobalNav []map[string]string `json:"globalNav" yaml:"globalNav" binding:"required"`
}

type Setting struct {
	Site SiteSetting `json:"site" yaml:"site" binding:"required"`
}

var vConfig *viper.Viper
var vSetting *viper.Viper

var siteSetting = &SiteSetting{}
var config = &Config{}

func init() {
	initConfig()
	initSetting()
}

func initConfig() {
	v := viper.New()
	configPath := GetProjectConfigPath()
	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件错误:%s \n", err))
	}

	vConfig = v
	v.Unmarshal(config)
	fmt.Println("配置加载完毕!")
}

func initSetting() {
	v := viper.New()
	configPath := GetProjectConfigPath()
	v.AddConfigPath(configPath)
	v.SetConfigName("setting")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取设置文件错误:%s \n", err))
	}
	vSetting = v
	v.Sub("site").Unmarshal(siteSetting)
	fmt.Println("设置加载完毕!")
}

func GetSiteSetting() *SiteSetting {
	//if siteSetting == nil {
	//	site := new(SiteSetting)
	//
	//	vSetting.Sub("site").Unmarshal(site)
	//	siteSetting = site
	//	return site
	//}

	return siteSetting
}

func WriteSiteSetting(Site *SiteSetting) error {
	vSetting.Set("site", *Site)
	if err := vSetting.WriteConfig(); err != nil {
		return err
	}
	// 刷新内存中的设置
	siteSetting = Site
	return nil
}

func GetConfig() *Config {
	return config
}

func GetAppDomain() string {
	production := GetConfig().Server.Production
	if production {
		return GetSiteSetting().Domain
	} else {
		return GetSiteSetting().DevDomain
	}
}

func GetVodCallbackUrl() string {
	production := GetConfig().Server.Production
	if production {
		return GetConfig().Vod.CallbackUrl
	} else {
		return GetConfig().Vod.DevCallbackUrl
	}
}

func GetOssCallbackUrl() string {
	production := GetConfig().Server.Production
	if production {
		return GetConfig().Oss.CallbackUrl
	} else {
		return GetConfig().Oss.DevCallbackUrl
	}
}
func GetOssCdnDomain(withSlash bool) string {
	domain := GetConfig().Oss.CdnDomain

	if withSlash {
		domain = domain + "/"
	}
	return domain
}

const GUMOLA_HOME = "GUMOLA_HOME"

func GetProjectPath() string {
	return "./"
}

func GetProjectConfigPath() string {
	return filepath.Join(GetProjectPath(), "config")
}
