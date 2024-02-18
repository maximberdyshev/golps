package consumer

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type InlineMarkup string

const (
	DefaultInlineMarkup InlineMarkup = "Default"
	RequestInlineMarkup InlineMarkup = "Request"
	BackInlineMarkup    InlineMarkup = "Back"
)

func (c *Consumer) setDefaultInlineMarkup() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Запросить предмет", "Запросить предмет"),
			tgbotapi.NewInlineKeyboardButtonData("Добавить предмет", "Добавить предмет"),
		),
	)
}

func (c *Consumer) setRequestInlineMarkup() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Вставить монеты", "Вставить монеты"),
			tgbotapi.NewInlineKeyboardButtonData("В меню", "В меню"),
		),
	)
}

func (c *Consumer) setBackInlineMarkup() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В меню", "В меню"),
		),
	)
}
