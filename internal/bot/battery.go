package bot

import (
	"SmsForwarder_tgbot/internal/sms"
	"fmt"
	"log"

	"gopkg.in/telebot.v3"
)

// RegisterBattery 获取电池信息
func (h *Handler) RegisterBattery(bot *telebot.Bot) {
	log.Println("RegisterBattery")
	bot.Handle("/ba", func(c telebot.Context) error {
		log.Println(c.Message().Text)
		res, err := sms.Post("/battery/query", sms.Data{})
		if err != nil {
			fmt.Println(err)
		}

		return c.Send(res, telebot.ModeMarkdownV2)
	})
}
