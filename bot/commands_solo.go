//commands_solo.go
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
	if message.IsCommand() {
		parts := strings.Fields(message.Text)
		if len(parts) >= 2 && parts[0] == "/approve" {
			// Получите значение "id" из parts[1] и обработайте его
			id := parts[1]
			b.handleApproveCommand(message.Chat.ID, id)
			//////////////////////////// вызов функции для отпрвки данному юзеру смс с ссылкой на чат
			
		} else if parts[0] == "/help" {
			b.sendActionsKeyboard(message.Chat.ID)

			// Отправляем сообщение с командами
			helpMessage := "Список доступных команд:\n" +
				"/approve [ID] - одобрить пользователя\n" +
				"/help - отобразить этот список команд"
			msg := tgbotapi.NewMessage(message.Chat.ID, helpMessage)
			_, err := b.api.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}else if {
			
		} else {
			b.handleUnknownCommand(message.Chat.ID)
		}
	} else {
		// Обработка обычных текстовых сообщений
		b.handleRegularMessage(message)
	}
}

func (b *Bot) handleApproveCommand(chatID int64, id string) {
    userID, err := strconv.Atoi(id)
    if err != nil {
        // Обработка ошибки, если id не может быть преобразовано в int
        errorMessage := "Некорректный ID пользователя."
        msg := tgbotapi.NewMessage(chatID, errorMessage)
        _, err := b.api.Send(msg)
        if err != nil {
            log.Println(err)
        }
        return
    }

    // Вызываем функцию для обновления статуса
    err = updateUserStatusToApproved(userID)
    if err != nil {
        // Обработка ошибки, например, отправка сообщения об ошибке
        errorMessage := "Произошла ошибка при обновлении статуса пользователя."
        msg := tgbotapi.NewMessage(chatID, errorMessage)
        _, err := b.api.Send(msg)
        if err != nil {
            log.Println(err)
        }
        return
    }

    // Отправляем подтверждение пользователю
    confirmationMessage := "Статус пользователя с ID " + id + " был успешно изменен на одобренный."
    msg := tgbotapi.NewMessage(chatID, confirmationMessage)
    _, err = b.api.Send(msg)
    if err != nil {
        log.Println(err)
    }
}

func (b *Bot) handleRegularMessage(message *tgbotapi.Message) {
	chatID := message.Chat.ID

	// Обработка обычных текстовых сообщений
	switch message.Text {
	case "Одобрить нового пользователя":
		users, err := getUsersWithStatusZero()
    if err != nil {
        // Обработка ошибки, например, отправка сообщения об ошибке
        errorMessage := "Произошла ошибка при запросе данных из базы данных."
        msg := tgbotapi.NewMessage(chatID, errorMessage)
        _, err := b.api.Send(msg)
        if err != nil {
            log.Println(err)
        }
        return
    }

    if len(users) == 0 {
        msg := tgbotapi.NewMessage(chatID, "Нет новых пользователей для одобрения.")
        _, err := b.api.Send(msg)
        if err != nil {
            log.Println(err)
        }
        return
    }

    // Отправляем список пользователей боту
    userListMessage := "Список пользователей для одобрения:\n"
    for _, user := range users {
        userListMessage += user + "\n"
    }
    msg := tgbotapi.NewMessage(chatID, userListMessage)
    _, err = b.api.Send(msg)
    if err != nil {
        log.Println(err)
    }

	case "добавить пост":
		msg := tgbotapi.NewMessage(chatID, "До свидания! Если у вас возникнут вопросы, не стесняйтесь спрашивать.")
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
		}
	default:
		// Обработка других обычных текстовых сообщений
		msg := tgbotapi.NewMessage(chatID, "Я не понимаю вашего сообщения.")
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func (b *Bot) sendActionsKeyboard(chatID int64) {
	keyboard := createActionsKeyboard()

	msg := tgbotapi.NewMessage(chatID, "Выберите действие:")
	msg.ReplyMarkup = keyboard

	_, err := b.api.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b *Bot) handleUnknownCommand(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Неизвестная команда. Используйте /help для получения помощи.")
	_, err := b.api.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func createActionsKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Одобрить нового пользователя"),
			tgbotapi.NewKeyboardButton("добавить пост"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("проверить"),
			tgbotapi.NewKeyboardButton("Действие 4"),
		),
	)

	return keyboard
}