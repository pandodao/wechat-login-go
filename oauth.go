package wechat

import (
	"context"
	"fmt"
	"net/http"
)

type (
	OAuthLoginResp struct {
		AccessToken    string `json:"access_token"`
		ExpiresIn      int    `json:"expires_in"`
		RefreshToken   string `json:"refreshToken"`
		OpenID         string `json:"openid"`
		Scope          string `json:"scope"`
		IsSnapshotUser int    `json:"is_snapshotuser"`
		UnionID        string `json:"unionid"`
	}

	UserInfoResp struct {
		OpenID     string   `json:"openid"`
		NickName   string   `json:"nickname"`
		Sex        int      `json:"sex"`
		Province   string   `json:"province"`
		City       string   `json:"city"`
		Country    string   `json:"country"`
		Headimgurl string   `json:"headimgurl"`
		Privilege  []string `json:"privilege"`
		UnionID    string   `json:"unionid"`
	}
)

func (c *Client) OAuthLogin(ctx context.Context, code string) (*OAuthLoginResp, error) {
	output := OAuthLoginResp{}
	uri := fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code", "/sns/oauth2/access_token", c.appID, c.appSecret, code)
	if err := c.request(ctx, http.MethodGet, uri, nil, &output); err != nil {
		return nil, err
	}

	return &output, nil
}

func (c *Client) QueryUserinfo(ctx context.Context, accessToken, openID, lang string) (*UserInfoResp, error) {
	output := UserInfoResp{}
	uri := fmt.Sprintf("%s?access_token=%s&openid=%s&lang=%s", "/sns/userinfo", accessToken, openID, lang)
	if err := c.request(ctx, http.MethodGet, uri, nil, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
