package bot

import (
	"SmsForwarder_tgbot/internal/sms"
	"fmt"
	"log"
	"strconv"
	"strings"

	"gopkg.in/telebot.v3"
)

// RegisterSend 注册发送短信命令
func (h *Handler) RegisterSend(bot *telebot.Bot) {
	log.Println("RegisterSend")
	bot.Handle("/send", func(c telebot.Context) error {
		log.Println(c.Message().Text)
		fullText := c.Message().Text

		parts := strings.Fields(fullText)

		if len(parts) < 4 {
			return c.Send("使用方法: /send <电话号码> <卡槽数字 1|2> <短信内容>")
		}

		phoneNumber := parts[1] // 切片索引1是电话号码
		simSlot, _ := strconv.Atoi(parts[2])
		messageContent := parts[3] // 切片索引2是短信内容

		if len(parts) > 4 {
			messageContent = strings.Join(parts[3:], " ")
		}

		err := c.Send(fmt.Sprintf("正在使用卡槽 %d 发送短信到 %s，内容: '%s'", simSlot, phoneNumber, messageContent))
		if err != nil {
			fmt.Println(err)
		}

		res, err := sms.Post("/sms/send", sms.Data{
			SimSlot:     simSlot,
			PhoneNumber: phoneNumber,
			MsgContent:  messageContent,
		})

		if err != nil {
			fmt.Println(err)
			return err
		}

		err = c.Send(res, telebot.ModeMarkdownV2)
		if err != nil {
			return err
		}

		return nil
	})
}
