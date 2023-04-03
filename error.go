package wechat

import "fmt"

type ErrorWrapper struct {
	error `json:"-"`
	// {"errcode": 1234,"errmsg":"..."}
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (e *ErrorWrapper) Error() string {
	return fmt.Sprintf("wechat error: %d, %s", e.Errcode, e.Errmsg)
}
