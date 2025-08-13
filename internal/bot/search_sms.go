package bot

import (
	"SmsForwarder_tgbot/internal/sms"
	"fmt"
	"log"
	"strconv"
	"strings"

	"gopkg.in/telebot.v3"
)

// RegisterSearchSms 注册搜索短信命令
func (h *Handler) RegisterSearchSms(bot *telebot.Bot) {
	log.Println("RegisterSearchSms")
	bot.Handle("/ss", func(c telebot.Context) error {
		log.Println(c.Message().Text)
		fullText := c.Message().Text

		parts := strings.Fields(fullText)

		if len(parts) == 0 {
			err := c.Send("/ss <关键词> [页码] [分页大小] [类型 1=接收 2=发送]")
			if err != nil {
				return err
			}
		}

		keyword := parts[1]
		pageNum := 1
		pageSize := 10
		tp := 1

		if len(parts) > 2 {
			pageNum, _ = strconv.Atoi(parts[2])
		}

		if len(parts) > 3 {
			pageSize, _ = strconv.Atoi(parts[3])
		}
		if len(parts) > 4 {
			tp, _ = strconv.Atoi(parts[4])
		}

		res, err := sms.Post("/sms/query", sms.Data{
			Keyword:  keyword,
			PageNum:  pageNum,
			PageSize: pageSize,
			Type:     tp,
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
