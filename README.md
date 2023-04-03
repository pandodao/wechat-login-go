# wechat-login-go

A Simple Go package for [WeChat OAuth](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html)

一个简单的 Go 包，用于微信网页授权。

## Installation

```bash
go get github.com/pandodao/wechat-login-go
```

## Preparation

- Please read the official documentation first: [WeChat OAuth](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html)
- Sign in to [WeChat 测试号](https://mp.weixin.qq.com/debug/cgi-bin/sandboxinfo?action=showinfo&t=sandbox/index), complete all the basic information, including "JS接口安全域名" and "授权回调页面域名"（at "网页授权获取用户基本信息" section）.
- And then you can get the `AppID` and `AppSecret` of the test account.

## Usage

### 1. Init wechat client

```go
wechatClient := wechat.New(WECHAT_APP_ID, WECHAT_APP_SECRET)
```

### 2. Ask user to visit the following URL to grant access

```
https://open.weixin.qq.com/connect/oauth2/authorize?appid=APPID&redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE&state=STATE#wechat_redirect
```

- `APPID` is the `AppID` of your test account
- `REDIRECT_URI` is a URI, which will be called by WeChat after user grant access to your app. Make sure the URI's domain is in the "授权回调页面域名" list.

Wechat will redirect user to the `REDIRECT_URI` with a `code` parameter, which is used to get the access token.

```
REDIRECT_URI?code=CODE&state=STATE
```

You must read the `code` parameter from the URL, and then use it to get the access token.

### 3. Exchange code for access token

```go
code := r.URL.Query().Get("code")
resp, err := wechatClient.OAuthLogin(ctx, code)
if err != nil {
  return nil, err
}

if resp.AccessToken == "" || resp.OpenID == "" {
  return nil, errors.New("invalid access token")
}

// save the open ID, union ID and access token for later use
```