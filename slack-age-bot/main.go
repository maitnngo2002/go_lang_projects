package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

// function to print command events to terminal
func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command events: ")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	// set up the environment
	os.Setenv("SLACK_BOT_TOKEN", SLACK_BOT_TOKEN_HERE)
	os.Setenv("SLACK_APP_TOKEN", SLACK_BOT_TOKEN_HERE)

	// create the bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: 	"yob calculator",
		Example: 		"my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year") // get the year from the user
			yob, err := strconv.Atoi(year) // convert to integer
			if err != nil {
				println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	context, cancel := context.WithCancel(context.Background())

	defer cancel() 
	/* If you fail to cancel the context, the goroutine that WithCancel or WithTimeout created will be retained in memory indefinitely 
	 (until the program shuts down), causing a memory leak. If you do this a lot, your memory will balloon significantly. 
	It's best practice to use a defer cancel() immediately after calling WithCancel() or WithTimeout() */

	err := bot.Listen(context)
	if err != nil {
		log.Fatal(err)
	}
}
