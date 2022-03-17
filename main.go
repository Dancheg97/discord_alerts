package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/ghodss/yaml"
	"github.com/joho/godotenv"
)

var discord *discordgo.Session
var channel_id string
var port string

func init() {
	godotenv.Load()
	envport, found := os.LookupEnv("PORT")
	if !found {
		log.Fatal("problem loading token from env")
	}
	port = envport
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

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	inpyaml, err := yaml.JSONToYAML([]byte(bodyBytes))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}
	highlightedYaml := "```yaml\n" + string(inpyaml) + "\n```"
	_, err = discord.ChannelMessageSend(channel_id, highlightedYaml)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}
	fmt.Println(err)
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	_, err = discord.ChannelMessageSend(channel_id, string(bodyBytes))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}
	fmt.Println(err)
}

func main() {
	fmt.Println("Bot is running...")
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/text", jsonHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
