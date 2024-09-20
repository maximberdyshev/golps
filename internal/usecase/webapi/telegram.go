package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"golps/config"
	"golps/internal/usecase"
	"golps/internal/usecase/entity"
	"golps/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

const (
	errSendText          = "error: sending text message"
	errSendTextAndMarkup = "error: sending text message and inline markup"
	errCacthCallback     = "error: catching callback"
	errToken             = "error: missing or invalid token"
)

type (
	User struct {
		userID         int64
		chatID         int64
		userName       string
		callback       string
		vendingMachine *VendingMachine
	}

	WebAPI struct {
		*tgbotapi.BotAPI
		Users map[string]*User
	}
)

func New(ctx context.Context) (*WebAPI, error) {
	token := config.FromContext(ctx).Telegram.Token

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errToken, err)
	}
	return &WebAPI{bot, make(map[string]*User)}, nil
}

func (w *WebAPI) Run(uc *usecase.UseCase) {
	logger := logger.FromContext(uc.Ctx)

	// DEBUG:
	user1, err := strconv.Atoi(os.Getenv("USER_1"))
	if err != nil {
		logger.Fatal("user_1", zap.Error(err))
	}
	user2, err := strconv.Atoi(os.Getenv("USER_2"))
	if err != nil {
		logger.Fatal("user_2", zap.Error(err))
	}
	// --

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := w.GetUpdatesChan(updateConfig)
	for update := range updates {
		jData, err := json.Marshal(update)
		if err != nil {
			logger.Error("json", zap.Error(err))
		}
		logger.Info("recieved update", zap.String("update", string(jData)))

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

			logger.Info("recivied message", zap.String("message", tgUserMSG),
				zap.String("user_name", tgUserName), zap.Int64("user_id", tgUserID),
				zap.Int64("chat_id", tgChatID))

			switch tgUserID {
			case int64(user1), int64(user2):
			default:
				continue
			}

			uuid := w.searchUser(tgUserID, tgChatID, tgUserName, tgUserMSG)

			msgTxt := fmt.Sprintf("Осталось предметов: %v", w.Users[uuid].vendingMachine.GetItemCount())
			markup := DefaultInlineMarkup
			err := w.NewText(entity.Message{
				ChatID: tgChatID,
				Text:   msgTxt,
				Markup: &markup},
				true)
			if err != nil {
				log.Println(err)
			}

		case update.CallbackQuery != nil:
			tgUserName = update.CallbackQuery.From.UserName
			tgUserID = update.CallbackQuery.From.ID
			tgChatID = update.CallbackQuery.Message.Chat.ID
			callbackData = update.CallbackQuery.Data

			logger.Info("recivied callback", zap.String("callback", callbackData),
				zap.String("user_name", tgUserName), zap.Int64("user_id", tgUserID),
				zap.Int64("chat_id", tgChatID))

			switch tgUserID {
			case int64(user1), int64(user2):
			default:
				continue
			}

			if err := w.catchCallback(update.CallbackQuery.ID, callbackData); err != nil {
				logger.Warn("can't catch callback", zap.Error(err))
			}

			uuid := w.searchUser(tgUserID, tgChatID, tgUserName, callbackData)

			w.callbackRouting(uuid, update)

		default:
			continue
		}
	}
}

func (w *WebAPI) NewText(m entity.Message, withMarkup bool) error {
	msg := tgbotapi.NewMessage(m.ChatID, m.Text)
	if withMarkup {
		msg.ReplyMarkup = w.inlineMarkup(*m.Markup)
	}

	if _, err := w.Send(msg); err != nil {
		if withMarkup {
			return fmt.Errorf("%s: %w", errSendTextAndMarkup, err)
		}
		return fmt.Errorf("%s: %w", errSendText, err)
	}
	return nil
}

func (w *WebAPI) EditTextAndMarkup(m entity.Message) error {
	msg := tgbotapi.NewEditMessageTextAndMarkup(m.ChatID, *m.MessageID, m.Text, w.inlineMarkup(*m.Markup))
	if _, err := w.Send(msg); err != nil {
		return fmt.Errorf("%s: %w", errSendTextAndMarkup, err)
	}
	return nil
}

func (w *WebAPI) catchCallback(id, text string) error {
	cb := tgbotapi.NewCallback(id, text)
	if _, err := w.Request(cb); err != nil {
		return fmt.Errorf("%s: %w", errCacthCallback, err)
	}
	return nil
}

func (w *WebAPI) callbackRouting(uuid string, update tgbotapi.Update) {
	switch update.CallbackQuery.Data {
	case "Запросить предмет":
		w.Users[uuid].vendingMachine.RequestItem(*w, update)
	case "Добавить предмет":
		w.Users[uuid].vendingMachine.AddItem(1, *w, update)
	case "Вставить монеты":
		w.Users[uuid].vendingMachine.InsertMoney(10, *w, update)
	case "В меню":
		w.Users[uuid].vendingMachine.Exit(*w, update)
	}
}

// TODO: move to repo ?
func (w *WebAPI) searchUser(userID, chatID int64, tgUserName, tgCallback string) (uuid string) {
	uuid = strconv.FormatInt(userID, 10) + strconv.FormatInt(chatID, 10)

	if _, ok := w.Users[uuid]; !ok {
		w.Users[uuid] = &User{
			userID:         userID,
			chatID:         chatID,
			userName:       tgUserName,
			callback:       tgCallback,
			vendingMachine: NewVendingMachine(1, 10),
		}
	}
	return uuid
}
