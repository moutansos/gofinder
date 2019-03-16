package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

//TODO: Way to register channel

func main() {
	fmt.Println("Starting Go Finder bot...")

	config := readConfig("config.json")

	discord, err := discordgo.New("Bot " + config.BotKey)
	if err != nil {
		log.Fatal("Unable to connect to discord")
	}
	user, err := discord.User("@me")
	if err != nil {
		log.Fatal("Cant retrieve the user for self:", err)
	}
	id := user.ID
	log.Printf("User id is %s", id)

	discord.AddHandler(initBot)
	discord.AddHandler(command)

	if err := discord.Open(); err != nil {
		log.Fatal("Error opening discord:", err)
	}
	defer discord.Close()

	listings := make(chan listing)

	go runFinder(listings)

	for listing := range listings {
		content := fmt.Sprintf("%s: %s", listing.Title, listing.Url)
		_, err = discord.ChannelMessageSend(config.ChannelId, content)
		if err != nil {
			log.Println("Error sending listing: ", err)
		}
	}

	// <-make(chan struct{})
}

func initBot(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(0, "I'm alive!")
	if err != nil {
		log.Fatal("Error setting my status: ", err)
	}
	servers := discord.State.Guilds
	log.Println(fmt.Sprintf("Now connected to %d servers.", len(servers)))
}

func command(discord *discordgo.Session, message *discordgo.MessageCreate) {
	botUser, err := discord.User("@me")
	if err != nil {
		log.Printf("Cannot get the id of self")
		return
	}

	messageUser := message.Author

	if messageUser.ID == botUser.ID || messageUser.Bot {
		return //This is either another bot or its this bot
	}

	log.Printf("From: %s Message: %s", message.Author.Username, message.Content)
	discord.ChannelMessageSend(message.ChannelID, message.Content)
}
