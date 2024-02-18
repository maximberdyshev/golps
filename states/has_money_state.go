package states

import (
	"fmt"
	"golearnpatternstate/consumer"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem(bot *consumer.Consumer, update tgbotapi.Update) {
	msgTxt := "Выполняется другая операция"
	if err := bot.SendEditTxtAndMrkp(
		update.CallbackQuery.From.ID,
		update.CallbackQuery.Message.MessageID,
		msgTxt,
		consumer.BackInlineMarkup,
	); err != nil {
		log.Println(err)
	}
}

func (i *HasMoneyState) addItem(count int, bot *consumer.Consumer, update tgbotapi.Update) {
	msgTxt := "Выполняется другая операция"
	if err := bot.SendEditTxtAndMrkp(
		update.CallbackQuery.From.ID,
		update.CallbackQuery.Message.MessageID,
		msgTxt,
		consumer.BackInlineMarkup,
	); err != nil {
		log.Println(err)
	}
}

func (i *HasMoneyState) insertMoney(money int, bot *consumer.Consumer, update tgbotapi.Update) {
	msgTxt := "Выполняется другая операция"
	if err := bot.SendEditTxtAndMrkp(
		update.CallbackQuery.From.ID,
		update.CallbackQuery.Message.MessageID,
		msgTxt,
		consumer.BackInlineMarkup,
	); err != nil {
		log.Println(err)
	}
}

func (i *HasMoneyState) dispenseItem(bot *consumer.Consumer, update tgbotapi.Update) {
	i.vendingMachine.itemCount -= 1

	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem, "noItem")

		msgTxt := fmt.Sprintf("Предметов больше нет.\n\nОсталось предметов: %v", i.vendingMachine.GetItemCount())
		if err := bot.SendEditTxtAndMrkp(
			update.CallbackQuery.From.ID,
			update.CallbackQuery.Message.MessageID,
			msgTxt,
			consumer.DefaultInlineMarkup,
		); err != nil {
			log.Println(err)
		}
	} else {
		i.vendingMachine.SetState(i.vendingMachine.HasItem, "hasItem")

		msgTxt := fmt.Sprintf("Осталось предметов: %v", i.vendingMachine.GetItemCount())
		if err := bot.SendEditTxtAndMrkp(
			update.CallbackQuery.From.ID,
			update.CallbackQuery.Message.MessageID,
			msgTxt,
			consumer.DefaultInlineMarkup,
		); err != nil {
			log.Println(err)
		}
	}
}

func (i *HasMoneyState) exit(bot *consumer.Consumer, update tgbotapi.Update) {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem, "noItem")

		msgTxt := fmt.Sprintf("Предметов больше нет.\n\nОсталось предметов: %v", i.vendingMachine.GetItemCount())
		if err := bot.SendEditTxtAndMrkp(
			update.CallbackQuery.From.ID,
			update.CallbackQuery.Message.MessageID,
			msgTxt,
			consumer.DefaultInlineMarkup,
		); err != nil {
			log.Println(err)
		}
	} else {
		i.vendingMachine.SetState(i.vendingMachine.HasItem, "hasItem")

		msgTxt := fmt.Sprintf("Осталось предметов: %v", i.vendingMachine.GetItemCount())
		if err := bot.SendEditTxtAndMrkp(
			update.CallbackQuery.From.ID,
			update.CallbackQuery.Message.MessageID,
			msgTxt,
			consumer.DefaultInlineMarkup,
		); err != nil {
			log.Println(err)
		}
	}
}
