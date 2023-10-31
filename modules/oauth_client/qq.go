package oauth_client

//
//import (
//	"bill/global"
//	"bill/models"
//	"context"
//	"golang.org/x/oauth2"
//)
//
//const qqOpenIDUrl = "https://graph.qq.com/oauth2.0/me"
//
//type qqOauthClient struct {
//}
//
//func (client qqOauthClient) GetConfig() *oauth2.Config {
//	return qqConf
//}
//
//func (client qqOauthClient) GetAuthorizeUrl(callbackUrl string) string {
//	cfg := client.GetConfig()
//	cfg.RedirectURL = callbackUrl
//	return cfg.AuthCodeURL("state")
//}
//
//func (client qqOauthClient) Upload(c context.Context, code string) (*oauth2.Token, error) {
//	panic("implement me")
//}
//
//func (client qqOauthClient) GetUserInfo(c context.Context, token *oauth2.Token) (*models.OAuthUser, error) {
//	panic("implement me")
//}
//
//var qqConf *oauth2.Config
//
//func init() {
//
//	qqConf = &oauth2.Config{
//		ClientID:     global.ConfigFile.GetString("third_user.github.client_id"),
//		ClientSecret: global.ConfigFile.GetString("third_user.github.client_secret"),
//		Scopes:       []string{"user:email"},
//		//RedirectURL:  global.ConfigFile.GetString("third_user.github.redirect"),
//		Endpoint: oauth2.Endpoint{
//			AuthURL:  "https://graph.qq.com/oauth2.0/authorize",
//			TokenURL: "https://graph.qq.com/oauth2.0/token",
//		},
//	}
//}
//
//func convertQQUser(qqUser *models.QQUser) *models.OAuthUser {
//	oUser := &models.OAuthUser{
//		Id:     qqUser.Login,
//		Name:   qqUser.Name,
//		Avatar: qqUser.AvatarUrl,
//	}
//
//	return oUser
//}
