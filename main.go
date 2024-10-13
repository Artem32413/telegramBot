package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()
}

func mustToken() string{
	token := flag.String("token-bot", "", "token-bot-pass")
	flag.Parse()
	if *token == ""{
		log.Fatal("token-bot = nil")
	}
}

