package rule34

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/MarkSmersh/go-telegram/components/rule34/types/general"
	"github.com/MarkSmersh/go-telegram/components/rule34/types/methods"
)

type Rule34 struct {
	token  string
	userId int
}

func NewRule34(token string, userId int) Rule34 {
	return Rule34{
		token:  token,
		userId: userId,
	}
}

func (r Rule34) Request(method string, params any) ([]byte, error) {
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
		"https://api.rule34.xxx/index.php?page=dapi&s=%s&q=index&api_key=%s&user_id=%d&json=%d&%s",
		method,
		r.token,
		r.userId,
		1,
		paramsValues.Encode(),
	)

	slog.Debug(url)

	res, err := http.Get(url)

	if err != err {
		slog.Error(err.Error())
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)

	if err != err {
		slog.Error(err.Error())
		return []byte{}, err
	}

	if len(body) <= 0 {
		return body, errors.New("Nothing is found by your prompt.")
	}

	return body, err
}

func (r Rule34) Posts(params methods.Post) ([]general.Post, error) {
	result, err := r.Request("post", params)
	var data []general.Post
	json.Unmarshal(result, &data)
	return data, err
}
