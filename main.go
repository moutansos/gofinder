package main

import (
	"log"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Println("Starting Go Finder bot...")

	discord, err := discordgo.New("Bot " + "<<API KEY>>")
	if err != nil {
		log.Fatal("Unable to connect to discord")
	}
	user, err := discord.User("@me")
	if err != nil {
		log.Fatal("Cant retireve the user for self:", err)
	}
	id := user.ID
	log.Printf("User id is %s", id)

	discord.AddHandler(initBot)

	if err := discord.Open(); err != nil {
		log.Fatal("Error opening discord:", err)
	}
	defer discord.Close()
	
	<-make(chan struct{})
}

func initBot(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(0, "I'm alive!")
	if err != nil {
		log.Fatal("Error setting my status: ", err)
	}
	servers := discord.State.Guilds
	log.Println(fmt.Sprintf("Now connected to %d servers.", len(servers)))
}