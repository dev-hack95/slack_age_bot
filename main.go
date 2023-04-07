package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}

}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4298311150771-5069975416550-JB52bGPa7fDVjFZxJK7wb0ex")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A0528FVCPRQ-5062015764695-7b8831c5517a8fe146d8e05b07136b63ed29a776f737fc834674735ee105cf1c")

	// Create Bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
}
