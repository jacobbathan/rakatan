package discord

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/jacobbathan/rakatan/mandalore/config"
	"github.com/jacobbathan/rakatan/mandalore/handler"
	"github.com/bwmarrin/discordgo"
)

var log = config.NewLogger("discord")

func Init() {
	// Create a new Discord Session
	token := config.GetEnv().DISCORD_BOT_TOKEN
	// commandsChannelId := config.GetEnv().COMMANDS_CHANNEL_ID

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Errorf("error creating discord session: %v", err)
		return
	}

	// Open Websocket Connection
	err = dg.Open()
	if err != nil {
		log.Errorf("error opening connection with discord: %v", err)
		return
	}

	// Register all Handlers and Actions
	registerCommandHandlers(dg, handler.CommandHandlers)

	if err != nil {
		log.Errorf("error creating slash commands: %v", err)
		return
	}

	log.Infoln("Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	log.Warnln("B4-D4 is exiting. Graceful shutdown in action...")

	dg.Close()
}

func registerCommandHandlers(s *discordgo.Session, commandHandlers []handler.CommandHandler) {
	log.Infof("registering %d command handlers", len(commandHandlers))
	// Register all Commands
	for _, handler := range commandHandlers {
		cmd := handler.Command()
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			log.Errorf("error creating command %s: %v", cmd.Name, err)
			continue
		}
	}
	// Register all Handlers
	s.AddHandler(func(dg *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			for _, handler := range commandHandlers {
				if handler.Command().Name == i.ApplicationCommandData().Name {
					handler.Handler(dg, i)
					return
				}
			}
		}
	})

}

