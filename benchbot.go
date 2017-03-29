package main

import (
	"os"

	"github.com/go-chat-bot/bot/slack"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/sanjaybv/net-buddy/plugins/hello"
)

func main() {

	// do not forget to supply the slack token
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
