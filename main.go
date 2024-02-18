package main

import (
	"fmt"
	"golearnpatternstate/consumer"
	"golearnpatternstate/states"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type User struct {
	userTgID       int64
	chatTgID       int64
	tgUserName     string
	tgCallback     string
	vendingMachine *states.VendingMachine
}

type application struct {
	UserMemo map[string]*User
	BotMemo  *consumer.Consumer
}

func main() {
	//
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%v", err)
	}

	// debug users
	user1, err1 := strconv.Atoi(os.Getenv("USER_1"))
	if err1 != nil {
		log.Fatalf("%v", err1)
	}
	user2, err2 := strconv.Atoi(os.Getenv("USER_2"))
	if err2 != nil {
		log.Fatalf("%v", err2)
	}

	//
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	//
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("%v", err)
	}

	//
	app := &application{
		UserMemo: make(map[string]*User),
		BotMemo:  &consumer.Consumer{Bot: bot},
	}

	//
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		// jData, err := json.Marshal(update)
		// if err != nil {
		// 	log.Println(err)
		// }
		// log.Println("jData:", string(jData))

		var (
			tgUserName, tgUserMSG, callbackData string
			tgUserID, tgChatID                  int64
		)

		switch {
		case update.Message != nil:
			tgUserMSG = update.Message.Text
			tgUserName = update.Message.From.UserName
			tgUserID = update.Message.From.ID
			tgChatID = update.Message.Chat.ID

			log.Printf("получен message: %s от пользователя: %s(%d)\n", tgUserMSG, tgUserName, tgUserID)

			// check access
			switch tgUserID {
			case int64(user1), int64(user2):
			default:
				continue
			}
			//

			uuid := app.findUser(tgUserID, tgChatID, tgUserName, callbackData)

			msgTxt := fmt.Sprintf("Осталось предметов: %v", app.UserMemo[uuid].vendingMachine.GetItemCount())
			err := app.BotMemo.SendNewTxtAndMrkup(tgUserID, msgTxt, consumer.DefaultInlineMarkup)
			if err != nil {
				log.Println(err)
			}

		case update.CallbackQuery != nil:
			tgUserName = update.CallbackQuery.From.UserName
			tgUserID = update.CallbackQuery.From.ID
			tgChatID = update.CallbackQuery.Message.Chat.ID
			callbackData = update.CallbackQuery.Data

			log.Printf("получен callback: %s от пользователя: %s(%d)\n", callbackData, tgUserName, tgUserID)

			// check access
			switch tgUserID {
			case int64(user1), int64(user2):
			default:
				continue
			}
			//

			if err := app.BotMemo.CatchCallback(update.CallbackQuery.ID, callbackData); err != nil {
				log.Println(err)
			}

			uuid := app.findUser(tgUserID, tgChatID, tgUserName, callbackData)

			app.checkCallbackRoute(uuid, update)

		default:
			continue
		}
	}
}

func (app *application) findUser(userTgID, chatTgID int64, tgUserName, tgCallback string) (uuid string) {
	uuid = strconv.FormatInt(userTgID, 10) + strconv.FormatInt(chatTgID, 10)

	if _, ok := app.UserMemo[uuid]; !ok {
		app.UserMemo[uuid] = &User{
			userTgID:       userTgID,
			chatTgID:       chatTgID,
			tgUserName:     tgUserName,
			tgCallback:     tgCallback,
			vendingMachine: states.NewVendingMachine(1, 10),
		}
	}

	return uuid
}

func (app *application) checkCallbackRoute(uuid string, update tgbotapi.Update) {
	switch update.CallbackQuery.Data {
	case "Запросить предмет":
		app.UserMemo[uuid].vendingMachine.RequestItem(app.BotMemo, update)
	case "Добавить предмет":
		app.UserMemo[uuid].vendingMachine.AddItem(1, app.BotMemo, update)
	case "Вставить монеты":
		app.UserMemo[uuid].vendingMachine.InsertMoney(10, app.BotMemo, update)
	case "В меню":
		app.UserMemo[uuid].vendingMachine.Exit(app.BotMemo, update)
	}
}
