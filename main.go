package main

import (
	"github.com/Luca-11/BotDiscord.git/bot"
	"github.com/Luca-11/BotDiscord.git/config"
)



func main(){
	err := config.LoadConfig()
	if err != nil {
		return
	}
	
	bot.Start()

	<- make(chan struct{})

	return
}