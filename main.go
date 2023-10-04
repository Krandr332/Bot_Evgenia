package main

import (
	"Bot_Evgenia/bot"
	"Bot_Evgenia/web"
)


func main() {
	// Запуск Telegram бота в горутине
	go bot.StartBot()

	// Запуск веб-сайта
	web.StartWebServer()
}
