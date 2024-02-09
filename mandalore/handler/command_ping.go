package handler 

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type PingHandler struct{}

func (h *PingHandler) Command() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name: "ping",
		Description: "Tests bot command to see if it is alive at all",
	}
}

func (h *PingHandler) Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Ignore messages created by discord bot
	if i.Member.User.ID == s.State.User.ID {
		return
	}

	if i.Type == discordgo.InteractionApplicationCommand {
		// Check if the command name is "ping"
		if i.ApplicationCommandData().Name == "ping" {
			// Respond to the interaction with "pong"
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "pong",
				},
			})
			if err != nil {
				fmt.Println("Error responding to interaction:", err)
			}
		}
	}
}
