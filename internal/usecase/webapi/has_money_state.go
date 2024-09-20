package webapi

// TODO: move to usecase ?

import (
	"fmt"
	"log"

	"golps/internal/usecase/entity"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem(bot WebAPI, update tgbotapi.Update) {
	msgTxt := "Выполняется другая операция"
	markup := BackInlineMarkup
	if err := bot.EditTextAndMarkup(entity.Message{
		ChatID:    update.CallbackQuery.From.ID,
		MessageID: &update.CallbackQuery.Message.MessageID,
		Text:      msgTxt,
		Markup:    &markup,
	}); err != nil {
		log.Println(err)
	}
}

func (i *HasMoneyState) addItem(count int, bot WebAPI, update tgbotapi.Update) {
	msgTxt := "Выполняется другая операция"
	markup := BackInlineMarkup
	if err := bot.EditTextAndMarkup(entity.Message{
		ChatID:    update.CallbackQuery.From.ID,
		MessageID: &update.CallbackQuery.Message.MessageID,
		Text:      msgTxt,
		Markup:    &markup,
	}); err != nil {
		log.Println(err)
	}
}

func (i *HasMoneyState) insertMoney(money int, bot WebAPI, update tgbotapi.Update) {
	msgTxt := "Выполняется другая операция"
	markup := BackInlineMarkup
	if err := bot.EditTextAndMarkup(entity.Message{
		ChatID:    update.CallbackQuery.From.ID,
		MessageID: &update.CallbackQuery.Message.MessageID,
		Text:      msgTxt,
		Markup:    &markup,
	}); err != nil {
		log.Println(err)
	}
}

func (i *HasMoneyState) dispenseItem(bot WebAPI, update tgbotapi.Update) {
	i.vendingMachine.itemCount -= 1

	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem, "noItem")

		msgTxt := fmt.Sprintf("Предметов больше нет.\n\nОсталось предметов: %v", i.vendingMachine.GetItemCount())
		markup := DefaultInlineMarkup
		if err := bot.EditTextAndMarkup(entity.Message{
			ChatID:    update.CallbackQuery.From.ID,
			MessageID: &update.CallbackQuery.Message.MessageID,
			Text:      msgTxt,
			Markup:    &markup,
		}); err != nil {
			log.Println(err)
		}
	} else {
		i.vendingMachine.SetState(i.vendingMachine.HasItem, "hasItem")

		msgTxt := fmt.Sprintf("Осталось предметов: %v", i.vendingMachine.GetItemCount())
		markup := DefaultInlineMarkup
		if err := bot.EditTextAndMarkup(entity.Message{
			ChatID:    update.CallbackQuery.From.ID,
			MessageID: &update.CallbackQuery.Message.MessageID,
			Text:      msgTxt,
			Markup:    &markup,
		}); err != nil {
			log.Println(err)
		}
	}
}

func (i *HasMoneyState) exit(bot WebAPI, update tgbotapi.Update) {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem, "noItem")

		msgTxt := fmt.Sprintf("Предметов больше нет.\n\nОсталось предметов: %v", i.vendingMachine.GetItemCount())
		markup := DefaultInlineMarkup
		if err := bot.EditTextAndMarkup(entity.Message{
			ChatID:    update.CallbackQuery.From.ID,
			MessageID: &update.CallbackQuery.Message.MessageID,
			Text:      msgTxt,
			Markup:    &markup,
		}); err != nil {
			log.Println(err)
		}
	} else {
		i.vendingMachine.SetState(i.vendingMachine.HasItem, "hasItem")

		msgTxt := fmt.Sprintf("Осталось предметов: %v", i.vendingMachine.GetItemCount())
		markup := DefaultInlineMarkup
		if err := bot.EditTextAndMarkup(entity.Message{
			ChatID:    update.CallbackQuery.From.ID,
			MessageID: &update.CallbackQuery.Message.MessageID,
			Text:      msgTxt,
			Markup:    &markup,
		}); err != nil {
			log.Println(err)
		}
	}
}
