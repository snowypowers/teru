package main

import (
	"gopkg.in/telegram-bot-api.v4"
)

var regions = map[string][]string{
	"North":[]string{
		"Lim Chu Kang",
		"Mandai",
		"Seletar",
		"Sembawang",
		"Sungei Kadut",
		"Woodlands",
		"Yishun"},
	"South":[]string{
		"Bukit Merah",
		"Bukit Timah",
		"City",
		"Geylang",
		"Jalan Bahar",
		"Kallang",
		"Marine Parade",
		"Queenstown",
		"Sentosa",
		"Southern Islands",
		"Tanglin"},
	"East":[]string{
		"Bedok",
		"Changi",
		"Hougang",
		"Pasir Ris",
		"Paya Lebar",
		"Pulau Tekong",
		"Pulau Ubin",
		"Punggol",
		"Sengkang",
		"Tampines"},
	"West":[]string{
		"Boon Lay",
		"Bukit Batok",
		"Bukit Panjang",
		"Choa Chu Kang",
		"Jurong East",
		"Jurong Island",
		"Jurong West",
		"Pioneer",
		"Tengah",
		"Tuas",
		"Western Islands",
		"Western Water Catchment"},
	"Central":[]string{
		"Ang Mo Kio",
		"Bishan",
		"Central Water Catchment",
		"Novena",
		"Serangoon",
		"Toa Payoh"}}

var wf2KbRow1 = tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/wf2 all"))
var wf2KbRow2 = tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/wf2 north"))
var wf2KbRow3 = tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/wf2 west"), tgbotapi.NewKeyboardButton("/wf2 central"), tgbotapi.NewKeyboardButton("/wf2 west"))
var wf2KbRow4 = tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/wf2 south"))
var wf2FullKb = tgbotapi.ReplyKeyboardMarkup{
	[][]tgbotapi.KeyboardButton{wf2KbRow1, wf2KbRow2, wf2KbRow3, wf2KbRow4},
	true,
	true,
	false}
