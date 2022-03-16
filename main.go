package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var token string

func init() {
	godotenv.Load()
	envtoken, found := os.LookupEnv("TOKEN")
	if !found {
		log.Fatal("problem loading token from env")
	}
	token = envtoken
}

func main() {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("unable to create a bot, bad token ", err)
	}
	user, err := discord.User("@me")
	if err != nil {
		log.Fatal("Unable to create user for bot ", err)
	}
	err = discord.Open()
	if err != nil {
		log.Fatal("Unable to connect to dicsord ", err)
	}
	fmt.Println(user, " bot is running")
	<-make(chan struct{})
}
