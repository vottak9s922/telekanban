package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found - using system vars")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is missing")
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	sugar.Infof("Authorized on account %s", bot.Self.UserName)

	// use polling
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// handler loop

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		msgText := update.Message.Text

		sugar.Infof("Received from %d: %s", chatID, msgText)

		var reply string

		switch msgText {
		case "/start":
			reply = "hey, this is the kanban bot!"
		case "/help":
			reply = "/start, /help"
		default:
			reply = "i only understand /start and /help for now"
		}

		msg := tgbotapi.NewMessage(chatID, reply)
		msg.ParseMode = tgbotapi.ModeMarkdown

		if _, err := bot.Send(msg); err != nil {
			sugar.Error(err)
		}
	}

}
