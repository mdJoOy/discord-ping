package main

import (
	"fmt"
	"github.com/mdjOoy/discord-ping/bot"
	"github.com/mdjOoy/discord-ping/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()

	<-make(chan struct{})
	return

}
