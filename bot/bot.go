// dbot.go

package bot

import (
	"fmt"
	"log"


	"gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	api        *tgbotapi.BotAPI
	userStates map[int64]*UserState
}

type UserState struct {
	State       RegistrationState
	
	FirstName   string
	LastName    string
	MiddleName  string
	PhoneNumber string
	Email       string
	Region      string
}

type RegistrationState int

const (
	StateStart RegistrationState = iota
	StateFullName
	StateRegion
	StatePhoneNumber
	StateEmail
	StateComplete
)

func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:        api,
		userStates: make(map[int64]*UserState),
	}, nil
}

func (b *Bot) Start() {
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := b.api.GetUpdatesChan(u)
    if err != nil {
        log.Fatal(err)
    }

    for update := range updates {
        if update.Message == nil {
            continue
        }

        chatID := update.Message.Chat.ID
        userState := b.userStates[chatID] // Изменено

        tgID := update.Message.From.ID

        if update.Message.Chat.IsGroup() {
            
            // Обработка команд и текстовых сообщений из группового чата
            if update.Message.IsCommand() {
                // Обработка команд
                b.handleGroupCommand(update.Message )
            } else {
                // Обработка текстовых сообщений
                b.handleGroupMessage(update.Message, )
            }
        } else {
            // Это личное сообщение
            if userState == nil {
                userState = &UserState{State: StateStart}
                b.userStates[chatID] = userState
            }

            if checkForUserInSystem(fmt.Sprintf("%d", tgID)) == 0 {
                if userState.State == StateStart {
                    b.handleStart(chatID, userState)
                } else {
                    b.handleRegistration(update.Message, userState)
                }
            } else {
                isAdmin, err := checkAdminStatus(tgID)
                if err != nil {
                    log.Println("Ошибка при проверке admin_status:", err)
                } else if isAdmin  {
                    // Пользователь - администратор
                    fmt.Println("Пользователь админ")
                    b.handleSoloCommand(update.Message)

                } else {
                    // Пользователь - не администратор
					b.handleSoloCommand(update.Message)
                    fmt.Println("Пользователь не админ")
                }
            }
        }
    }
}





