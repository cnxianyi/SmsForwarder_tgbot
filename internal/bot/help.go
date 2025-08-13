package bot

import (
	"log"

	"gopkg.in/telebot.v3"
)

// RegisterHelp 注册帮助命令
func (h *Handler) RegisterHelp(bot *telebot.Bot) {
	log.Println("RegisterHelp")
	bot.Handle("/help", func(c telebot.Context) error {
		log.Println(c.Message().Text)
		return c.Send(`
获取详情: /star
获取电量: /ba
发送短信: /send <电话号码> <卡槽数字 1|2> <短信内容>
关键词搜索短信: /ss <关键词> [页码] [分页大小] [类型 1=接收 2=发送]
手机号搜索通话: /sc <手机号> [页码] [分页大小] [类型 1=呼入 2=呼出 3=未接]
`)
	})
}
