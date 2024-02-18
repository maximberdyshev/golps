package consumer

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Consumer struct {
	Bot *tgbotapi.BotAPI
}

const (
	errBotSendTxt        = "ошибка отправки текстового сообщения"
	errBotSendTxtAndMrkp = "ошибка отправки текстового сообщения и inline-разметки"
	errCacthCallback     = "ошибка получения callback"
)

func (c *Consumer) SendNewTxt(userTgID int64, txt string) error {
	msg := tgbotapi.NewMessage(userTgID, txt)

	if _, err := c.Bot.Send(msg); err != nil {
		return fmt.Errorf("%s: %w", errBotSendTxt, err)
	}

	return nil
}

func (c *Consumer) SendNewTxtAndMrkup(userTgID int64, txt string, mrkp InlineMarkup) error {
	msg := tgbotapi.NewMessage(userTgID, txt)

	switch mrkp {
	case DefaultInlineMarkup:
		msg.ReplyMarkup = c.setDefaultInlineMarkup()
	case RequestInlineMarkup:
		msg.ReplyMarkup = c.setRequestInlineMarkup()
	case BackInlineMarkup:
		msg.ReplyMarkup = c.setBackInlineMarkup()
	}

	if _, err := c.Bot.Send(msg); err != nil {
		return fmt.Errorf("%s: %w", errBotSendTxtAndMrkp, err)
	}

	return nil
}

func (c *Consumer) SendEditTxtAndMrkp(chatTgID int64, msgID int, txt string, mrkp InlineMarkup) error {
	var choosenMarkup tgbotapi.InlineKeyboardMarkup

	switch mrkp {
	case DefaultInlineMarkup:
		choosenMarkup = c.setDefaultInlineMarkup()
	case RequestInlineMarkup:
		choosenMarkup = c.setRequestInlineMarkup()
	case BackInlineMarkup:
		choosenMarkup = c.setBackInlineMarkup()
	}

	msg := tgbotapi.NewEditMessageTextAndMarkup(chatTgID, msgID, txt, choosenMarkup)
	if _, err := c.Bot.Send(msg); err != nil {
		return fmt.Errorf("%s: %w", errBotSendTxtAndMrkp, err)
	}

	return nil
}

func (c *Consumer) CatchCallback(id, data string) error {
	callback := tgbotapi.NewCallback(id, data)
	if _, err := c.Bot.Request(callback); err != nil {
		return fmt.Errorf("%s: %w", errCacthCallback, err)
	}

	return nil
}
