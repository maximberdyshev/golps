package webapi

// TODO: move to usecase ?

import (
	"fmt"
	"log"

	"golps/internal/usecase/entity"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem(bot WebAPI, update tgbotapi.Update) {
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

func (i *ItemRequestedState) addItem(count int, bot WebAPI, update tgbotapi.Update) {
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

func (i *ItemRequestedState) insertMoney(money int, bot WebAPI, update tgbotapi.Update) {
	if money < i.vendingMachine.itemPrice {
		msgTxt := "Прежде необходимо вставить моменты"
		markup := BackInlineMarkup
		if err := bot.EditTextAndMarkup(entity.Message{
			ChatID:    update.CallbackQuery.From.ID,
			MessageID: &update.CallbackQuery.Message.MessageID,
			Text:      msgTxt,
			Markup:    &markup,
		}); err != nil {
			log.Println(err)
		}
	} else {
		i.vendingMachine.SetState(i.vendingMachine.hasMoney, "hasMoney")
		i.vendingMachine.itemCount--

		msgTxt := "Выдан предмет"
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
}

func (i *ItemRequestedState) dispenseItem(bot WebAPI, update tgbotapi.Update) {
	msgTxt := "Прежде необходимо вставить моменты"
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

func (i *ItemRequestedState) exit(bot WebAPI, update tgbotapi.Update) {
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
