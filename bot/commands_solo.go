package bot

import (
	"log"
    "gopkg.in/telegram-bot-api.v4"
	"strings"
	"strconv"
	"fmt"
)

func (b *Bot) handleStart(chatID int64, userState *UserState) {
	if userState.State == StateStart {
	msg := tgbotapi.NewMessage(chatID, "Доброго времени суток! Вы не зарегистрированы, давайте знакомиться. Пожалуйста, отправьте свое ФИО.")
	_, err := b.api.Send(msg)
	if err != nil {
		log.Println(err)
	}
	userState.State = StateFullName
}
}

func (b *Bot) handleRegistration(message *tgbotapi.Message, userState *UserState) {
chatID := message.Chat.ID

switch userState.State {
case StateFullName:
	parts := strings.Fields(message.Text)
	if len(parts) == 3 {
		userState.FirstName = parts[0]
		userState.LastName = parts[1]
		userState.MiddleName = parts[2]
		userState.State = StateRegion

		keyboard := tgbotapi.ReplyKeyboardMarkup{
			Keyboard: [][]tgbotapi.KeyboardButton{
				{
					tgbotapi.NewKeyboardButton("Азербайджан"),
					tgbotapi.NewKeyboardButton("Армения"),
					tgbotapi.NewKeyboardButton("Грузия"),
				},
				{
					tgbotapi.NewKeyboardButton("Казахстан"),
					tgbotapi.NewKeyboardButton("Киргизия"),
					tgbotapi.NewKeyboardButton("Монголия"),
				},
				{
					tgbotapi.NewKeyboardButton("Таджикистан"),
					tgbotapi.NewKeyboardButton("Туркменистан"),
					tgbotapi.NewKeyboardButton("Узбекистан"),
				},
			},
			OneTimeKeyboard: true,
		}

		msg := tgbotapi.NewMessage(chatID, "Выберите ваш регион:")
		msg.ReplyMarkup = keyboard

		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
		}
	} else {
		msg := tgbotapi.NewMessage(chatID, "Пожалуйста, отправьте ФИО в правильном формате (Ф И О).")
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
case StateRegion:
	userState.Region = message.Text
	userState.State = StatePhoneNumber

	msg := tgbotapi.NewMessage(chatID, "Отправьте ваш номер телефона.")
	_, err := b.api.Send(msg)
	if err != nil {
		log.Println(err)
	}
case StatePhoneNumber:
	userState.PhoneNumber = message.Text
	userState.State = StateEmail

	msg := tgbotapi.NewMessage(chatID, "Отправьте вашу почту.")
	_, err := b.api.Send(msg)
	if err != nil {
		log.Println(err)
	}
case StateEmail:
	userState.Email = message.Text
	userState.State = StateComplete

	msg := tgbotapi.NewMessage(chatID, "Регистрация завершена. Спасибо!")
	fmt.Println(userState)
	userIDStr := strconv.Itoa(message.From.ID)

	err := CreateUserAccount(userIDStr, userState.FirstName, userState.LastName, userState.MiddleName, userState.PhoneNumber, userState.Region, userState.Email)
	if err != nil {
		log.Println(err)
	}

	_, sendErr := b.api.Send(msg)
		if sendErr != nil {
			log.Println(sendErr)
		}

}
}


func (b *Bot) handleSoloCommand(message *tgbotapi.Message) {
    // Обработка команд из группового чата
    switch message.Command() {
    case "help":
        // Обработка команды /start из группового чата
        b.handleGroupStart(message.Chat.ID)

    default:
        // Обработка неизвестных команд
        b.handleUnknownCommand(message.Chat.ID)
    
    }
}