package bot

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func init() {
	var err error
	Bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
}

func TelegramWebhookHandler(w http.ResponseWriter, r *http.Request) {
	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if update.Message == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	if update.Message.IsCommand() && update.Message.Command() == "start" {

		// Кнопка со ссылкой
		btn := tgbotapi.NewInlineKeyboardButtonURL(
			"🎮 Open Casino",
			"https://casinoenginebot.ru",
		)

		// Клавиатура
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(btn),
		)

		// Сообщение
		msg := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"🎁 Gift cases with the highest chances!",
		)

		msg.ReplyMarkup = keyboard

		_, _ = Bot.Send(msg)
	}

	w.WriteHeader(http.StatusOK)
}
