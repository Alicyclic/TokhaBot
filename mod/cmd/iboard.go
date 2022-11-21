package cmd

import (
	"bot/popsicle/v1/mod/uti"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := AddCommand("imageboard", "Search for images!")
	command.AddOption("tags", "The tags to search for", discordgo.ApplicationCommandOptionString)
	command.AddOptionChoices("source", "The imageboard to search", discordgo.ApplicationCommandOptionInteger, []CommandChoiceOption{{
		CommandOption: &CommandOption{
			Name:        "Danbooru",
			Description: "Search Danbooru",
		},
		Value: 1,
	}, {
		CommandOption: &CommandOption{
			Name:        "Rule34",
			Description: "Search Rule34",
		},
		Value: 4,
	}, {
		CommandOption: &CommandOption{
			Name:        "yandere",
			Description: "Search yandere",
		},
		Value: 3,
	}, {
		CommandOption: &CommandOption{
			Name:        "konachan",
			Description: "Search for konachan images",
		},
		Value: 2,
	}})

	command.SetHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var image map[string]interface{}
		tags := i.ApplicationCommandData().Options[0].StringValue()
		channel, _ := s.State.GuildChannel(i.GuildID, i.ChannelID)
		switch {
		case !channel.NSFW:
			uti.NewEmbed().AddField("Error", "This channel is not marked as NSFW!").
				SetColor(0xFF0000).
				RespondEmbed(s, i)
		case image == nil || image["file_url"] == nil:
			uti.NewEmbed().AddField("Error", "No images found!").
				SetColor(0xFF0000).
				RespondEmbed(s, i)
		default:
			image_url := image["file_url"]
			uti.NewEmbed().
				AddField("Search", tags).
				SetImage(image_url.(string)).
				SetColor(0xFF0000).
				InlineAllFields().
				RespondEmbed(s, i)
		}
	})
}
