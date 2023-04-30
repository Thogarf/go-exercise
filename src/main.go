package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const commandPrefix string = "!go"

func main() {
	godotenv.Load()
	TOKEN := os.Getenv("TOKEN")
	botSession, err := discordgo.New("Bot " + TOKEN)

	if err != nil {
		log.Fatal(err)
	}

	botSession.AddHandler(func(sess *discordgo.Session, msg *discordgo.MessageCreate) {
		if msg.Author.ID == sess.State.User.ID {
			return
		}

		//DM Logic placeholder

		//Server Logic
		args := strings.Split(msg.Content, " ")

		if args[0] != commandPrefix {
			return
		}

		//Enter command "!go hello" in Discord
		if args[1] == "hello" {
			HelloWorldHandler(sess, msg)
		}
	})

	//Sets Discord Intents
	botSession.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = botSession.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer botSession.Close()

	fmt.Println("The bot is now online!")

	//Listener
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func HelloWorldHandler(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	sess.ChannelMessageSend(msg.ChannelID, "world!")
}
