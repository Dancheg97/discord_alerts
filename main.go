package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var discord *discordgo.Session
var channel_id string

func init() {
	godotenv.Load()
	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Fatal("problem loading token from env")
	}
	channel, found := os.LookupEnv("CHANNEL_ID")
	if !found {
		log.Fatal("problem loading channel id from env")
	}
	channel_id = channel
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("unable to create a bot, bad token ", err)
	}
	user, err := discord.User("@me")
	if err != nil {
		log.Fatal("Unable to create user for bot ", user, err)
	}
	err = discord.Open()
	if err != nil {
		log.Fatal("Unable to connect to dicsord ", err)
	}
}

func messageHadnler(w http.ResponseWriter, r *http.Request) {
	
}

func sendDiscordMessage(message string) error {
	_, err := discord.ChannelMessageSend(channel_id, message)
	return err
}

func main() {
	fmt.Println(" - bot is running...")
	http.HandleFunc("/", messageHadnler)
	
}
