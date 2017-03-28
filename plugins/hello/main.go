package hello

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/go-chat-bot/bot"
)

func gibberish(command *bot.Cmd) (msg string, err error) {

	fmt.Println("cmd: " + command.Command)
	fmt.Println("msg: " + command.Message)
	fmt.Println("arg: ", command.Args)
	msg = "asdf awrf asdf asf sf"

	cmd := exec.Command("ping", "-c 4", "127.0.0.1")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	msg = "# pinging \n" + "``` \n" + out.String() + "\n ```"

	return
}

func init() {

	fmt.Println("init")
	bot.RegisterCommand("hello", "testing bot.", "blabber", gibberish)
}
