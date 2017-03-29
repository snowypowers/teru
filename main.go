package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/telegram-bot-api.v4"

	"github.com/snowypowers/sgweatherbot/poller"
	"github.com/snowypowers/sgweatherbot/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")
	webLink := os.Getenv("WEBSITE_LINK")
	apiKey := os.Getenv("DATA_API_KEY")

	//Store
	s := poller.Poller()
	wf2 := poller.Subscription{
		"wf2",
		"https://api.data.gov.sg/v1/environment/2-hour-weather-forecast",
		apiKey,
		300}
	wf2Close := s.Listen(wf2)

	//Bot
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
		//Command Parsing
		cmd := update.Message.Command()
		args := update.Message.CommandArguments()
		out := ""
		switch cmd {
			case "wf2":
				out = getWF2(s.Value(wf2), args)
			case "start":
				out = "Hello! Welcome! This bot is under construction!"
			default:
				out = "hi"
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, out)
		bot.Send(msg)
	}

	//Teardown
	close(wf2Close)
}

func getWF2(data []byte, args string) string {
	d := store.ParseWf2(data)
	a := store.ParseArea(args)

	if a == "All" {
		return printAllWf2(d)
	}
	if a != "" {
		f := d.Forecasts[a]
		return "The weather at " + a + " is " + f
	}
	return "Sorry, did you spell the area wrongly?"
}

func printAllWf2(w store.WF2Update) string {
	s := "Weather for " + w.Timestamp.Format("Mon Jan 2 03:04 PM ") + " \n"
	for k,v:= range w.Forecasts {
		s += k + ": " + v + "\n"
	}
	return s
}
