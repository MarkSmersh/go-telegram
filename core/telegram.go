package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

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
		me, err := t.GetMe()

		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

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

			// slog.Debug(u.InlineQuery.Query)

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

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", t.Token, method)

	reqBody, _ := json.Marshal(params)

	// slog.Debug(string(reqBody))

	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		slog.Error(err.Error())
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

func genericRequest[R any](data []byte, err error) (R, error) {
	var res R

	if err != nil {
		slog.Error(err.Error())
		return res, err
	}

	json.Unmarshal(data, &res)

	return res, nil
}

func (t *Telegram) GetMe() (User, error) {
	res, err := genericRequest[general.User](t.Request("getMe", nil))
	return t.NewUser(res), err
}

func (t *Telegram) SendMessage(params methods.SendMessage) (Message, error) {
	res, err := genericRequest[general.Message](t.Request("sendMessage", params))
	return t.NewMessage(res), err
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

func (t *Telegram) SendPhoto(params methods.SendPhoto) (Message, error) {
	result, err := t.Request("sendPhoto", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return t.NewMessage(data), err
}

func (t *Telegram) SendMediaGroup(params methods.SendMediaGroup) ([]Message, error) {
	result, err := t.Request("sendMediaGroup", params)
	data := []general.Message{}
	json.Unmarshal(result, &data)

	messages := []Message{}

	for _, m := range data {
		messages = append(messages, t.NewMessage(m))
	}

	return messages, err
}

func (t *Telegram) AnswerInlineQuery(params methods.AnswerInlineQuery) (bool, error) {
	return genericRequest[bool](t.Request("answerInlineQuery", params))
}
