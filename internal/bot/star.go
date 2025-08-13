package bot

import (
	"SmsForwarder_tgbot/internal/sms"
	"fmt"
	"log"

	"gopkg.in/telebot.v3"
)

// RegisterStar 注册获取配置命令
func (h *Handler) RegisterStar(bot *telebot.Bot) {
	log.Println("RegisterStar")
	bot.Handle("/star", func(c telebot.Context) error {
		log.Println(c.Message().Text)
		res, err := sms.Post("/config/query", sms.Data{})
		if err != nil {
			fmt.Println(err)
		}

		return c.Send(res, telebot.ModeMarkdownV2)
	})
}
