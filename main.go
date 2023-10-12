// main.go
package main

import (
	"Bot_Evgenia/bot"
	// "Bot_Evgenia/web"
	"log"
)

func main() {
	token := "5435221086:AAFtVvXSL4ZJ427yHpCQXkqVnjUz0eHI3C8"
	bot, err := bot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Start()
}