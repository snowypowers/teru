package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")
	webLink := os.Getenv("WEBSITE_LINK")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert(webLink+bot.Token, "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)

	//Website
	http.Handle("/", http.FileServer(http.Dir("./website")))

	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
