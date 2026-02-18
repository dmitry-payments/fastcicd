package bot

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

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

		payload := map[string]string{
			"chat_id": strconv.FormatInt(update.Message.Chat.ID, 10),
			"text":    "🎁 Gift cases with the highest chances!",
			"reply_markup": `{
				"inline_keyboard":[[
					{
						"text":"🎮 Open Casino",
						"web_app":{"url":"https://casinoenginebot.ru"}
					}
				]]
			}`,
		}

		resp, err := Bot.MakeRequest("sendMessage", payload)
		if err != nil {
			log.Println(err)
			return
		}
		if !resp.Ok {
			log.Println(resp.Description)
		}
	}

	w.WriteHeader(http.StatusOK)
}
