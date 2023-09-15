package hooks

import (
	"log"

	"github.com/gtuk/discordwebhook"
)

func SendDiscordMessage(username, content, url string) bool {
	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	if err := discordwebhook.SendMessage(url, message); err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
