package oauth_client

import (
	"bill/modules/log"
	"context"
	"golang.org/x/oauth2"
)

const (
	Github = "github"
	QQ     = "qq"
	WeiXin = "weixin"
)

type OAuthClient interface {
	GetConfig() *oauth2.Config
	GetAuthorizeUrl(callbackUrl string) string
	GetToken(c context.Context, code string) (*oauth2.Token, error)
	//GetUserInfo(c context.Context, token *oauth2.Token) (*models.OAuthUser, error)
}

var clients = []string{Github, QQ, WeiXin}

func Create(clientType string) OAuthClient {
	var client OAuthClient
	if !clientExists(clientType) {
		log.GetSugar().Error("Oauth客户端不存在")
	}

	return client

}

func clientExists(clientType string) bool {
	exists := false

	for _, v := range clients {
		if clientType == v {
			exists = true
		}
	}

	return exists

}
