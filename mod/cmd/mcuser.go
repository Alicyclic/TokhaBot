package cmd

import (
	"bot/popsicle/v1/mod/games"
	"bot/popsicle/v1/mod/uti"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := AddCommand("user", "Find a user!")
	command.AddOption("name", "The name of the user to find", discordgo.ApplicationCommandOptionString)
	command.SetHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		val := i.ApplicationCommandData().Options[0].StringValue()
		var finder games.UserFinder = games.UserFinderDestructJson(val)
		uti.NewEmbed().
			SetAuthor(finder.Name).
			// SetDescription("The abyss, the darkest deity of them all, said to control undesirable dimensions and cast away.").
			// AddField("Mission", "The abyss has awakened! Unwilling to accept your fate, the abyss has sent you to the depths of the nether to find your true destiny.").
			SetThumbnail(fmt.Sprintf("https://minotar.net/avatar/%s/100.png", finder.Name)).
			SetFooter(finder.ID).
			InlineAllFields().
			RespondEmbed(s, i)
	})

}
