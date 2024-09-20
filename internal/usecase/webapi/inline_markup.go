package webapi

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	BackInlineMarkup    string = "Back"
	RequestInlineMarkup string = "Request"
	DefaultInlineMarkup string = "Default"
)

func (w *WebAPI) inlineMarkup(m string) tgbotapi.InlineKeyboardMarkup {
	switch m {
	case BackInlineMarkup:
		return w.backInlineMarkup()
	case RequestInlineMarkup:
		return w.requestInlineMarkup()
	default:
		return w.defaultInlineMarkup()
	}
}

func (w *WebAPI) backInlineMarkup() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В меню", "В меню"),
		),
	)
}

func (w *WebAPI) requestInlineMarkup() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Вставить монеты", "Вставить монеты"),
			tgbotapi.NewInlineKeyboardButtonData("В меню", "В меню"),
		),
	)
}

func (w *WebAPI) defaultInlineMarkup() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Запросить предмет", "Запросить предмет"),
			tgbotapi.NewInlineKeyboardButtonData("Добавить предмет", "Добавить предмет"),
		),
	)
}
