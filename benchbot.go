package main

import (
	"os"

	"github.com/go-chat-bot/bot/slack"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/sanjaybv/random/plugins/bench"
)

func main() {
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
