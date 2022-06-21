package main

import (
	"os"
	"reflect"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func StartBot() {
	tkn := os.Getenv("BtcBot")
	bot, err := tgbotapi.NewBotAPI(tkn)
	if err != nil {
		panic(err)
	}
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		//checking type of message to be text
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				//greeting message
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a bot for alerting BTC rate\n\n Type \"\\alert <x>\" to get notified when rate will be lower than <x> in USD")
				bot.Send(msg)
			default:
				// item := update.Message.Text
				// //searching for recipes
				// checkresult := SearchRecipes(item)
				// if len(checkresult) == 0 {
				// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No matches was found")
				// 	bot.Send(msg)
				// }
				// //sending all matches
				// for _, elem := range checkresult {
				// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, elem)
				// 	bot.Send(msg)
				//}
			}
		} else {
			//asking to type something
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please set some alert")
			bot.Send(msg)
		}
	}
}
