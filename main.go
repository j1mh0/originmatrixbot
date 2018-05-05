package main

import (
	"log"
	"strings"

	"github.com/google/uuid"
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
		//处理非inline模式的消息
		if update.Message != nil {
			command, param := procMessage(update.Message)
			log.Printf("[%s] command = %s, param = %s", update.Message.From.UserName, command, param)

			response := messageRoute(command, param)
			log.Printf("[%s] response = %s", update.Message.From.UserName, response)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}

		//处理inline模式的消息
		if update.InlineQuery != nil {
			log.Printf("[%s] query = %s", update.InlineQuery.From, update.InlineQuery.Query)
			var inlineConfig tgbotapi.InlineConfig
			inlineConfig.InlineQueryID = update.InlineQuery.ID
			results := make([]interface{}, 0)
			query := strings.ToLower(update.InlineQuery.Query)
			if result := queryRoute(query); result != nil {
				results = append(results, result)
				inlineConfig.Results = results
				if _, err := bot.AnswerInlineQuery(inlineConfig); err != nil {
					log.Printf("query response error = %s", err)
				}
			}
		}

	}
}

func procMessage(originMessage *tgbotapi.Message) (command string, param string) {

	if originMessage.IsCommand() {
		command = strings.ToLower(originMessage.Command())
		param = originMessage.CommandArguments()
	} else {
		param = originMessage.Text
	}

	return

}

func messageRoute(command string, param string) (response string) {
	switch command {
	case "echo":
		response = param
	case "start":
		response = "Hello,I am OriginMatrix offical bot"
	default:
		response = param

	}
	return
}

func queryRoute(query string) (result interface{}) {
	switch query {
	case "test":
		result = tgbotapi.NewInlineQueryResultArticle(uuid.New().String(), "Test Title", "Test Message")
	default:
		result = nil
	}
	return
}
