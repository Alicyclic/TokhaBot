package cmd

import (
	"bot/popsicle/v1/mod/games"
	"bot/popsicle/v1/mod/uti"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ListValues(maps map[string]interface{}) string {
	var message string
	for k, v := range maps {
		message += fmt.Sprintf("**%s** %v\n", k, v)
	}
	return message
}

func init() {
	command := AddCommand("wizcalc", "Calculate talent expressions")
	for _, att := range []string{"strength", "willpower", "power", "agility", "intelligence"} {
		command.Command.Options = append(command.Command.Options, &discordgo.ApplicationCommandOption{
			Name:        att,
			Description: "The " + att + " attribute of the pet.",
			Type:        discordgo.ApplicationCommandOptionInteger,
			Required:    true,
		})
	}
	command.SetHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		atts := make(map[string]int64)
		for idx, data := range i.ApplicationCommandData().Options {
			atts[data.Name] = i.ApplicationCommandData().Options[idx].IntValue()
		}
		embed := uti.NewEmbed()
		for k, v := range games.Calculate(atts) {
			embed.AddField(k, ListValues(v)).InlineAllFields()
		}
		embed.RespondEmbed(s, i)
	})
}
