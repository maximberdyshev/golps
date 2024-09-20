package webapi

// TODO: move to usecase ?

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// интерфейс State описывает возможные actions
type State interface {
	addItem(int, WebAPI, tgbotapi.Update)
	requestItem(WebAPI, tgbotapi.Update)
	insertMoney(int, WebAPI, tgbotapi.Update)
	dispenseItem(WebAPI, tgbotapi.Update)
	exit(WebAPI, tgbotapi.Update)
}
