package bot

import (
	"log"
    "gopkg.in/telegram-bot-api.v4"
)


func (b *Bot) handleGroupCommand(message *tgbotapi.Message) {
    // Обработка команд из группового чата
    switch message.Command() {
    case "start":
        // Обработка команды /start из группового чата
        b.handleGroupStart(message.Chat.ID)
    case "help":
        // Обработка команды /help из группового чата
        b.handleGroupHelp(message.Chat.ID)
    // Добавьте обработку других команд, если необходимо
    default:
        // Обработка неизвестных команд
        b.handleUnknownCommand(message.Chat.ID)
    
    }
}

func (b *Bot) handleGroupMessage(message *tgbotapi.Message) {
    // Обработка текстовых сообщений из группового чата
    // Пример: Ответить на текстовое сообщение
    responseText := "Спасибо за ваше сообщение: " + message.Text
    responseMsg := tgbotapi.NewMessage(message.Chat.ID, responseText)
    _, err := b.api.Send(responseMsg)
    if err != nil {
        log.Println("Ошибка при отправке сообщения:", err)
    }
}

func (b *Bot) handleGroupStart(chatID int64) {
    // Обработка команды /start из группового чата
    // Пример: Отправить приветственное сообщение
    msg := tgbotapi.NewMessage(chatID, "Привет, это бот! Для получения справки используйте команду /help.")
    _, err := b.api.Send(msg)
    if err != nil {
        log.Println("Ошибка при отправке сообщения:", err)
    }
}
func (b *Bot) handleGroupHelp(chatID int64) {
    // Обработка команды /help из группового чата
    // Пример: Отправить справочное сообщение
    helpText := "Это справочное сообщение. Вы можете использовать команды:\n\n" +
        "/start - начать взаимодействие с ботом\n" +
        "/help - получить эту справку\n" +
        // Добавьте здесь другие команды и их описания
        "..."
    msg := tgbotapi.NewMessage(chatID, helpText)
    _, err := b.api.Send(msg)
    if err != nil {
        log.Println("Ошибка при отправке сообщения:", err)
    }
}

func (b *Bot) handleUnknownCommand(chatID int64) {
    // Обработка неизвестных команд из группового чата
    // Пример: Отправить сообщение о неизвестной команде
    msg := tgbotapi.NewMessage(chatID, "Неизвестная команда. Используйте /help для получения справки.")
    _, err := b.api.Send(msg)
    if err != nil {
        log.Println("Ошибка при отправке сообщения:", err)
    }
}