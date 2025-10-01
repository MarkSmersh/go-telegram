package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/MarkSmersh/go-telegram/core/components"
	"github.com/MarkSmersh/go-telegram/types/general"
	"github.com/MarkSmersh/go-telegram/types/methods"
)

type Telegram struct {
	Token   string
	Eventer components.Updater
}

func NewTelegram(token string) Telegram {
	return Telegram{
		Token:   token,
		Eventer: components.Updater{},
	}
}

func (t *Telegram) Init(callback func(e User)) {
	if callback != nil {
		me, _ := t.GetMe()
		callback(me)
	}
	t.Polling()
}

func (t *Telegram) Polling() {
	updateId := 0

	for {
		req := methods.GetUpdates{
			Offset: updateId,
		}

		updates, _ := t.GetUpdates(req)

		for i := range updates {
			u := updates[i]

			if u.Message != nil {
				e := u.Message

				if e.Text != "" && (e.Text)[0] == '/' {
					go t.Eventer.Commands.Invoke(*u.Message)
					break
				}

				go t.Eventer.Messages.Invoke(*u.Message)
				break
			}

			if u.InlineQuery != nil {
				go t.Eventer.InlineQuery.Invoke(*u.InlineQuery)
			}

			if u.CallbackQuery != nil {
				go t.Eventer.CallbackQuery.Invoke(*u.CallbackQuery)
			}
		}

		if len(updates) <= 0 {
			continue
		}

		updateId = updates[len(updates)-1].UpdateID + 1
	}
}

func (t *Telegram) Request(method string, params any) ([]byte, error) {
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

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s?%s", t.Token, method, paramsValues.Encode())

	res, err := http.Get(url)

	if err != err {
		log.Println(err)
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)

	if err != err {
		log.Println(err)
		return []byte{}, err
	}

	result := general.TelegramRes{}

	json.Unmarshal(body, &result)

	resultBytes, err := json.Marshal(result.Result)

	if !result.Ok {
		slog.Debug(url)
		slog.Error(
			fmt.Sprintf("Telegram error. Code: %d. Description: %s", result.ErrorCode, result.Description),
		)
		return resultBytes, errors.New(result.Description)
	}

	return resultBytes, err
}

func (t *Telegram) GetMe() (User, error) {
	result, err := t.Request("getMe", nil)
	data := general.User{}
	json.Unmarshal(result, &data)
	return t.NewUser(data), err
}

func (t *Telegram) SendMessage(params methods.SendMessage) (Message, error) {
	result, err := t.Request("sendMessage", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return t.NewMessage(data), err
}

func (t *Telegram) ForwardMessage(params methods.ForwardMessage) (Message, error) {
	result, err := t.Request("forwardMessage", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return t.NewMessage(data), err
}

func (t *Telegram) CopyMessage(params methods.CopyMessage) (Message, error) {
	result, err := t.Request("copyMessage", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return t.NewMessage(data), err
}

func (t *Telegram) GetUpdates(params methods.GetUpdates) ([]general.Update, error) {
	result, err := t.Request("getUpdates", params)
	data := []general.Update{}
	json.Unmarshal(result, &data)
	return data, err
}

func (t *Telegram) EditMessageText(params methods.EditMessageText) (Message, error) {
	result, err := t.Request("editMessageText", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return t.NewMessage(data), err
}

func (t *Telegram) EditMessageReplyMarkup(params methods.EditMessageReplyMarkup) (Message, error) {
	result, err := t.Request("editMessageReplyMarkup", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return t.NewMessage(data), err
}

func (t *Telegram) AnswerCallbackQuery(params methods.AnswerCallbackQuery) error {
	_, err := t.Request("answerCallbackQuery", params)
	return err
}

func (t *Telegram) DeleteMessage(params methods.DeleteMessage) (bool, error) {
	result, err := t.Request("deleteMessage", params)
	data := false
	json.Unmarshal(result, &data)
	return data, err
}

func (t *Telegram) DeleteMessages(params methods.DeleteMessages) (bool, error) {
	result, err := t.Request("deleteMessages", params)
	data := false
	json.Unmarshal(result, &data)
	return data, err
}

func (t *Telegram) GetChatMember(params methods.GetChatMember) (ChatMember, error) {
	result, err := t.Request("getChatMember", params)
	data := general.ChatMember{}
	json.Unmarshal(result, &data)
	return t.NewChatMember(data, nil), err
}

func (t *Telegram) BanChatMember(params methods.BanChatMember) (bool, error) {
	result, err := t.Request("banChatMember", params)
	data := false
	json.Unmarshal(result, &data)
	return data, err
}

func (t *Telegram) UnbanChatMember(params methods.UnbanChatMember) (bool, error) {
	result, err := t.Request("unbanChatMember", params)
	data := false
	json.Unmarshal(result, &data)
	return data, err
}

func (t *Telegram) GetChat(params methods.GetChat) (general.ChatFullInfo, error) {
	result, err := t.Request("getChat", params)
	data := general.ChatFullInfo{}
	json.Unmarshal(result, &data)
	return data, err
}
