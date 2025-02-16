package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/lassejlv/trombo/commands"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var ENV_WARNING = "No token found in .env file (mabye this is production, then you don't need to worry about this)"

func main() {
	// Enable the pretty logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	err := godotenv.Load()

	if err != nil {
		log.Warn().Msg(ENV_WARNING)
	}

	// get token
	token := os.Getenv("DISCORD_TOKEN")

	if token == "" {
		log.Warn().Msg(ENV_WARNING)
	}

	// create new client
	bot, err := discordgo.New("Bot " + token)

	if err != nil {
		slog.Error(err.Error())
	}

	bot.AddHandler(messageCreate)

	bot.Identify.Intents = discordgo.IntentGuilds | discordgo.IntentsGuildMessages | discordgo.IntentGuildMessageReactions | discordgo.IntentsGuildMessageTyping

	err = bot.Open()

	if err != nil {
		slog.Error(err.Error())
	}

	log.Info().Msgf("%s is now running. Press CTRL-C to exit.", fmt.Sprintf("%s#%s", bot.State.User.Username, bot.State.User.Discriminator))

	// Keep the bot alive lol
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// close client
	bot.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.Bot {
		return
	}

	prefix := "+"

	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	commandName := strings.Split(" "+m.Content, " ")[1]

	for _, command := range commands.Commands {
		if command.Name == strings.Replace(commandName, prefix, "", 1) {
			command.Handler(s, m)
		}
	}
}
