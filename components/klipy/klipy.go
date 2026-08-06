package klipy

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/MarkSmersh/go-telegram/components/klipy/types/general"
	"github.com/MarkSmersh/go-telegram/components/klipy/types/methods"
)

type Klipy struct {
	token string
}

func NewKlipy(token string) Klipy {
	return Klipy{
		token: token,
	}
}

// func genericRequest[R, P any, M string](f func(M, P) ([]byte, error), method M, params P) (R, error) {
// 	var res R
//
// 	data, err := f(method, params)
//
// 	if err != nil {
// 		return res, err
// 	}
//
// 	json.Unmarshal(data, &res)
//
// 	return res, nil
// }

func genericRequest[R any](data []byte, err error) (R, error) {
	var res R

	if err != nil {
		return res, err
	}

	json.Unmarshal(data, &res)

	return res, nil
}

func (k Klipy) Search(params methods.Search) ([]general.Item, error) {
	return genericRequest[[]general.Item](k.Request("search", params))
}

func (k Klipy) Request(method string, params any) ([]byte, error) {
	// https://api.klipy.com/api/v1/{app_key}/gifs/search?page={page}&per_page={per_page}&q={q}&customer_id={customer_id}&locale={country_code}&content_filter={content_filter}

	paramsValues := url.Values{}

	if params != nil {
		var paramsMap map[string]any

		tmp, _ := json.Marshal(params)

		d := json.NewDecoder(strings.NewReader(string(tmp[:])))

		d.UseNumber()

		d.Decode(&paramsMap)

		for k, v := range paramsMap {
			paramsValues.Add(k, fmt.Sprintf("%v", v))
		}
	}

	url := fmt.Sprintf(
		"https://api.klipy.com/api/v1/%s/gifs/search?%s",

		k.token,
		paramsValues.Encode(),
	)

	slog.Debug(url)

	res, err := http.Get(url)

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		slog.Error(err.Error())
	}

	var r general.Response

	json.Unmarshal(data, &r)

	data = r.Data.Data

	return data, nil
}
