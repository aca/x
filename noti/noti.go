package noti

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/aca/x/log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	telegram_token  string
	telegram_chatid int64
	bot             *tgbotapi.BotAPI

	botInit sync.Once
)

func Send(v ...interface{}) (err error) {
	botInit.Do(func() {
		telegram_token = os.Getenv("NOTI_TELEGRAM_TOKEN")
		telegram_chatid, _ = strconv.ParseInt(os.Getenv("NOTI_TELEGRAM_CHATID"), 10, 0)

		if telegram_chatid == 0 || telegram_token == "" {
			err = fmt.Errorf("noti: invalid telegram_token: %v, telegram_chatid: %v", telegram_token, telegram_chatid)
            return
		} else {
			bot, err = tgbotapi.NewBotAPI(telegram_token)
			if err != nil {
				err = fmt.Errorf("noti: failed to init telegram: %w", err)
				return
			}
			log.Debugf("noti: initialized telegram bot")
		}
	})

    if err != nil {
        return err
    }

	msg := tgbotapi.NewMessage(820947043, fmt.Sprint(v...))
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
