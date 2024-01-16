package commands

import (
	"giga-squad-bot/handlers"

	"github.com/bwmarrin/discordgo"
)

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	CommandsList = []*discordgo.ApplicationCommand{
		{
			Name:        "steam",
			Description: "Everything related to the steam api",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "recently",
					Description: "View recently played games",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"steam": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			content := ""

			// As you can see, names of subcommands (nested, top-level)
			// and subcommand groups are provided through the arguments.
			switch options[0].Name {
			case "recently":
				content = handlers.GetRecentlyPlayedGames()
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})
		}}
)
