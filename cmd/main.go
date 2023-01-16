package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/shomali11/slacker"
	"strconv"
)

var slackToken string = "xapp-1-A04K6EJKHU4-4649936008115-d19e7ce560932e73a7b28ec0cd60589e69c0d02ceb794b11181f4862513260fe"

func main() {
	os.Setenv("SLACK_BOT_TOKEN","xoxb-4646931901317-4647210563189-qLjuElNZ2j9RMiym4KBwZ2kW")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A04K6EJKHU4-4649936008115-d19e7ce560932e73a7b28ec0cd60589e69c0d02ceb794b11181f4862513260fe")
	fmt.Println("Programa 06: Slack Bot para calcular edad")
	fmt.Println("Este programa sirve par aampliar mis conocimientos en GO")
	fmt.Println("Token asignado: " + slackToken)

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Examples: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx,cancel := context.WithCancel(context.Background())

	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}