package bot

import (
	"SmsForwarder_tgbot/internal/sms"
	"fmt"
	"log"
	"strconv"
	"strings"

	"gopkg.in/telebot.v3"
)

// RegisterSearchCall 注册搜索通话命令
func (h *Handler) RegisterSearchCall(bot *telebot.Bot) {
	log.Println("RegisterSearchCall")
	bot.Handle("/sc", func(c telebot.Context) error {
		log.Println(c.Message().Text)
		fullText := c.Message().Text

		parts := strings.Fields(fullText)

		if len(parts) == 0 {
			err := c.Send("/sc <手机号> [页码] [分页大小] [类型 1=呼入 2=呼出 3=未接]")
			if err != nil {
				return err
			}
		}

		phone := parts[1]
		pageNum := 1
		pageSize := 10
		tp := 1

		if len(parts) > 2 {
			fmt.Println(len(parts))
			pageNum, _ = strconv.Atoi(parts[2])
		}

		if len(parts) > 3 {
			pageSize, _ = strconv.Atoi(parts[3])
		}
		if len(parts) > 4 {
			tp, _ = strconv.Atoi(parts[4])
		}

		res, err := sms.Post("/call/query", sms.Data{
			Phone:    phone,
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
