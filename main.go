package main

import (
	"encoding/json"
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
	dis, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("unable to create a bot, bad token ", err)
	}
	discord = dis
	user, err := dis.User("@me")
	if err != nil {
		log.Fatal("Unable to create user for bot ", user, err)
	}
	err = dis.Open()
	if err != nil {
		log.Fatal("Unable to connect to dicsord ", err)
	}
}

func messageHadnler(w http.ResponseWriter, r *http.Request) {
	input := map[string]any{}
	json.NewDecoder(r.Body).Decode(&input)
	message := ""
	for k, v := range input {
		message = message + k + " - " + fmt.Sprintln(v) + "\n"
	}
	err := sendDiscordMessage(message)
	if err != nil {
		w.WriteHeader(200)
		return
	}
	fmt.Println(err)
}

func sendDiscordMessage(message string) error {
	_, err := discord.ChannelMessageSend(channel_id, message)
	return err
}

func main() {
	fmt.Println("Bot is running...")
	http.HandleFunc("/", messageHadnler)
	log.Fatal(http.ListenAndServe(":8092", nil))
}
