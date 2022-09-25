package noti

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aca/x/log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	telegram_token  string
	telegram_chatid int64
	bot             *tgbotapi.BotAPI
)

func init() {
	telegram_token = os.Getenv("NOTI_TELEGRAM_TOKEN")
	telegram_chatid, _ = strconv.ParseInt(os.Getenv("NOTI_TELEGRAM_CHATID"), 10, 0)
	var err error
	bot, err = tgbotapi.NewBotAPI(telegram_token)
	if err != nil {
		log.Errorf("noti: failed to init telegram: %v", err)
		return
	}
	log.Debug("noti: initialized telegram bot")
}

func Send(v ...interface{}) error {
	msg := tgbotapi.NewMessage(820947043, fmt.Sprint(v...))
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
