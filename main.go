package main

import (
	"log"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

const (
	botToken string = "492097173:AAHF7fNIvFe_1OFgtUy3vXXKuXuGGQ04TvM"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		command, param := procMessage(update.Message.Text)

		log.Printf("[%s] command = %s, param = %s", update.Message.From.UserName, command, param)

		response := route(command, param)
		log.Printf("[%s] response = %s", update.Message.From.UserName, response)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func procMessage(originMessage string) (command string, param string) {

	command = ""
	param = ""

	message := strings.TrimSpace(originMessage)

	if strings.HasPrefix(message, "/") {
		//是一个命令
		fields := strings.Fields(message)
		command = strings.ToLower(fields[0])
		if len(fields) > 1 {
			param = fields[1]
		}
	} else {
		//不是一个命令
		param = message
	}

	return

}

func route(command string, param string) (response string) {
	switch command {
	case "/echo":
		response = param
	case "/start":
		response = "欢迎使用OriginMatrix机器人"
	default:
		response = param

	}
	return
}
