package main

import (
	"SmsForwarder_tgbot/internal/bot"
	"SmsForwarder_tgbot/internal/sms"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func main() {
	// 尝试从 .env 文件加载环境变量
	// 如果文件不存在，则会忽略错误
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading configuration from environment variables.")
	}

	// 从环境变量中获取机器人 API Token
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("Telegram API Token not found. Please set the TELEGRAM_BOT_TOKEN environment variable.")
	}

	s := os.Getenv("SIGN")
	if s == "" {
		log.Fatal("SIGN not found. Please set the SIGN environment variable.")
	}
	sms.Sign = s

	sms.GetSign()

	bu := os.Getenv("BASE_URL")
	if bu == "" {
		log.Fatal("BASE_URL not found. Please set the BASE_URL environment variable.")
	}
	sms.BaseUrl = bu

	// 初始化机器人设置
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	// 创建机器人实例
	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	botHandler := bot.NewHandler()

	botHandler.RegisterStar(b)
	botHandler.RegisterSend(b)
	botHandler.RegisterHelp(b)
	botHandler.RegisterBattery(b)
	botHandler.RegisterSearchSms(b)
	botHandler.RegisterSearchCall(b)

	log.Println("Bot is running...")
	b.Start()
}
