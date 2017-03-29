package hello

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-chat-bot/bot"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

// TODO use flags instead of envs
var config *ssh.ClientConfig

func hello(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("Hello, %s!", command.User.RealName)
	return
}

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
	msg += "\n# pinging \n" + "``` \n" + out.String() + "\n ```"

	return
}

// reach command ssh into the local ssh server and executes whoami
func reach(command *bot.Cmd) (msg string, err error) {

	client, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Println("Failed to create session: ", err)
		return "", errors.Wrap(err, "Failed to create session")
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	return b.String(), nil
}

func initSSH() {

	config = &ssh.ClientConfig{
		User: "sanjay",
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("PASSWORD")),
		},
	}
}

func init() {

	fmt.Println("init")

	// invoke this command in Slack using "!hello"
	bot.RegisterCommand(
		"hello",
		"Sends a hello message with the username.",
		"",
		hello)

	bot.RegisterCommand(
		"gibber",
		"testing bot.",
		"blabber",
		gibberish)

	bot.RegisterCommand(
		"reach",
		"SSH into localhost",
		"",
		reach)

	initSSH()
}
