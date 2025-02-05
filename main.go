package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"arono/commands"
	"arono/db"
	"arono/util"
)

const prefix string = "~"

var challengeMap util.ChallengeMap = util.ChallengeMap{Map: make(map[string][]string)}
var duelMap util.DuelMap = util.DuelMap{Map: make(map[string]util.GameState)}

var dbConn db.DBConn = db.NewDBConnection("arono.db")

func main() {
	godotenv.Load(".env")

	session, _ := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	session.AddHandler(messageCreate)
	session.Identify.Intents = discordgo.IntentsGuildMessages

	session.Open()

	fmt.Println("Bot is running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	session.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore if the message is from itself or doesn't start with the prefix
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, prefix) {
		return
	}

	// extracting the command and the args after it
	contentSplit := strings.Split(m.Content, " ")
	command, rawArgs := contentSplit[0][1:], contentSplit[1:]

	// splitting and appending to the array
	var args []string
	for _, a := range rawArgs {
		if a != "" {
			args = append(args, a)
		}
	}

	// switch case for the commands
	// (this has to be done because go doesn't allow dynamic function calls)
	switch command {
	case "ping":
		commands.Ping(s, m, args)
	case "help":
		commands.Help(s, m, args)
	case "register":
		commands.Register(s, m, args, &dbConn)
	case "handle":
		commands.Handle(s, m, args, &dbConn)
	case "challenge":
		commands.Challenge(s, m, args, &duelMap, &challengeMap, &dbConn)
	case "accept":
		commands.Accept(s, m, args, &duelMap, &challengeMap)
	case "end":
		commands.End(s, m, args, &duelMap, &challengeMap)
	case "update":
		commands.Update(s, m, args, &duelMap, &challengeMap)
	}
}
