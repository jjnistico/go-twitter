package gotwit

import (
	"net/url"
	"strings"
)

type getOption func() (key string, val string)

func With(key string, vals ...string) getOption {
	return func() (string, string) {
		return key, strings.Join(vals, ",")
	}
}

func buildQueryParamsFromOptions(options []getOption) url.Values {
	urlVals := url.Values{}
	for _, opt := range options {
		key, val := opt()
		urlVals.Set(key, val)
	}
	return urlVals
}
