package hooks

import (
	"fmt"
	"time"

	"github.com/gtuk/discordwebhook"
)

func SendDiscordMessage(username, content, url string) bool {
	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}
	fmt.Printf("TimeNow:: %s\n", time.Now())
	if err := discordwebhook.SendMessage(url, message); err != nil {
		fmt.Printf("ERROR DISCORD:: %s\n", err)
		return false
	}

	return true
}
