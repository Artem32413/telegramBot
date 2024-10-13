package main

import (
	"flag"
	"lesson/clients/telegram"
	"log"
)
const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(mustToken())
}

func mustToken() string{
	token := flag.String("token-bot", "", "token-bot-pass")
	flag.Parse()
	if *token == ""{
		log.Fatal("token-bot = nil")
	}
	return *token 
}

