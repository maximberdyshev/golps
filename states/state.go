package states

import (
	"golearnpatternstate/consumer"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// интерфейс State описывает возможные actions
type State interface {
	addItem(int, *consumer.Consumer, tgbotapi.Update)
	requestItem(*consumer.Consumer, tgbotapi.Update)
	insertMoney(int, *consumer.Consumer, tgbotapi.Update)
	dispenseItem(*consumer.Consumer, tgbotapi.Update)
	exit(*consumer.Consumer, tgbotapi.Update)
}
