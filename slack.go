package main

import (
	"log"
	"os"

	"github.com/slack-go/slack"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	OAUTH_TOKEN := os.Getenv("OAUTH")
	CHANNEL_ID := os.Getenv("CHANNEL")

	api := slack.New(OAUTH_TOKEN)
	attachment := slack.Attachment{
		Pretext: "Mensagem:",
		Text:    "Olá, eu sou seu bot Jarvis",
	}

	channelId, timestamp, err := api.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionText("Iniciando Bot", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Printf("Enviando informacoes para o canal: %s em: %s\n", channelId, timestamp)
}
