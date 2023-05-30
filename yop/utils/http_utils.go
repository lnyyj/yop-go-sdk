// Package utils
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/20 4:15 PM
package utils

import (
	"github.com/lnyyj/yop-go-sdk/yop/request"
	"net/url"
	"sort"
	"strings"
)

func NormalizePath(path string) string {

	return strings.ReplaceAll(Normalize(path), "%2F", "/")
}

func Normalize(value string) string {
	var firstEncodeStr = url.QueryEscape(value)
	return encodeSpecialChar(firstEncodeStr)
}

func encodeSpecialChar(str string) string {
	// 空格
	str = strings.ReplaceAll(str, "+", "%20")
	return str
}

func EncodeParameters(params map[string][]string) string {
	if 0 == len(params) {
		return ""
	}
	var encodedNameValuePair []string
	for k, v := range params {
		for i := range v {
			encodedNameValuePair = append(encodedNameValuePair, toNameValuePair(k, v[i]))
		}
	}
	return strings.Join(encodedNameValuePair, "&")
}

func toNameValuePair(paramName string, paramValue string) string {
	return Normalize(paramName) + "=" + Normalize(paramValue)
}

func GetCanonicalQueryString(params map[string][]string) string {
	if 0 == len(params) {
		return ""
	}

	var parameterStrings []string

	for k, v := range params {
		if nil == v || 0 == len(v) {
			parameterStrings = append(parameterStrings, Normalize(k)+"=")
		} else {
			for i := range v {
				parameterStrings = append(parameterStrings, Normalize(k)+"="+Normalize(v[i]))
			}
		}

	}
	sort.Strings(parameterStrings)
	return strings.Join(parameterStrings, "&")
}

func UsePayloadForQueryParameters(yopRequest request.YopRequest) bool {
	var requestIsPOST = 0 == strings.Compare("POST", yopRequest.HttpMethod)
	var requestHasNoPayload = 0 == len(yopRequest.Content)
	return requestIsPOST && requestHasNoPayload
}
