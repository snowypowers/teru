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
		var msg tgbotapi.MessageConfig
		switch cmd {
			case "wf2":
				msg = processWf2(update, s.Value(wf2), args)
			case "start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! Welcome! This bot is under construction!\nType \\help for instructions!")
			case "help":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Use command \\wf2 to get the weather forecast for the next 2 hours")
			case "about":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Data retrieved from NEA. Made in Go.")
			default:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "hi")
		}
		bot.Send(msg)
	}

	//Teardown
	close(wf2Close)
}

func processWf2(update tgbotapi.Update, data []byte, args string) tgbotapi.MessageConfig {
	d := store.ParseWf2(data)
	a := store.ParseArea(args)
	log.Printf("args: %s", a)
	var msg tgbotapi.MessageConfig
	switch(a) {
		case "All":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, printAllWf2(d))
		case "North", "South", "East", "West", "Central":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, printAreaWf2(d, a))
		case "":
			if args == "" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Choose a region or type in the area name after \\wf2. \nFor example, \\wf2 Bedok.")
				msg.ReplyMarkup = wf2FullKb
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, did you spell the area wrongly?")
			}
		default:
			f := d.Forecasts[a]
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "The weather at " + a + " is " + f)
	}
	return msg
}

func printAllWf2(w store.WF2Update) string {
	s := "Weather for " + w.Timestamp.Format("Mon Jan 2 03:04 PM ") + " \n"
	for k,v:= range w.Forecasts {
		s += k + ": " + v + "\n"
	}
	return s
}

func printAreaWf2(w store.WF2Update, area string) string {
	s := "The 2 hour Nowcast for " + area + " at " + w.Timestamp.Format("Mon Jan 2 03:04 PM ") + " \n"
	for _,v := range regions[area] {
		s += v + ": " + w.Forecasts[v]
	}
	return s
}
