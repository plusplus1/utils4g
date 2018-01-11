package str

import (
	"fmt"
	"net/url"
)

func (u *Util) UrlEncode(data interface{}) (result string, err error) {
	if strData, ok := data.(string); ok {
		result = url.QueryEscape(strData)
		return
	}

	if m, ok := data.(map[string]string); ok {
		result = u.mssUrlEncode(m)
		return
	}

	if m, ok := data.(map[string]interface{}); ok {
		result = u.msiUrlEncode(m)
		return
	}

	if m, ok := data.(map[interface{}]interface{}); ok {
		result = u.miiUrlEncode(m)
		return
	}

	err = fmt.Errorf("UrlEncode fail, data type not supported")
	return

}

func (u *Util) mssUrlEncode(data map[string]string) string {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, v)
	}
	return values.Encode()
}

func (u *Util) msiUrlEncode(data map[string]interface{}) string {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, fmt.Sprintf("%v", v))
	}
	return values.Encode()
}

func (u *Util) miiUrlEncode(data map[interface{}]interface{}) string {
	values := url.Values{}
	for k, v := range data {
		values.Set(fmt.Sprintf("%v", k), fmt.Sprintf("%v", v))
	}
	return values.Encode()
}
