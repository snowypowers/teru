# Teru - SG Weather Bot

A Telegram bot for SG Weather written in Go. Conversations powered by API.AI. Data taken from NEA through data.gov.sg.

A simple chatbot on Telegram with NLP integrations. Retrieve 2 hour Nowcasts through the command `/weather`. Done as a project to pick up Golang.

## Try
The bot is hosted on Telegram under the tag @teru_bot

## Requirements

Golang (v1.8)

## Installation

```go
go get github.com/snowypowers/teru
```

Create a `.env` file in your working directory (eg. root) with the following lines:

```env
DATA_API_KEY = (API Key to access data.gov.sg API)
BOT_TOKEN = (Bot Token provided by Telegram)
NLP_TOKEN = (NLP Client Access Token provided by API.ai)
WEBSITE_LINK = (Your base website address)
```

From your working directory, run
```sh
$ ./go/bin/teru
```
## Links

[Go](https://golang.org/) - Programming language by Google
[Telegram](https://telegram.org/) - Messenging Service, supports games and bots!
[API.ai](https://api.ai/) - Natural Language Interactions for Bots
[Data.gov.sg Developers](https://developers.data.gov.sg/) - API Platform by SG Govt
