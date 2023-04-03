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
)

func (c *Client) OAuthLogin(ctx context.Context, code string) (*OAuthLoginResp, error) {
	output := OAuthLoginResp{}
	uri := fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code", "/sns/oauth2/access_token", c.appID, c.appSecret, code)
	if err := c.request(ctx, http.MethodGet, uri, nil, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
