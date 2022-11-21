package cmd

import (
	"bot/popsicle/v1/mod/games"
	"bot/popsicle/v1/mod/uti"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := AddCommand("genshinartifacts", "Randomize Genshin Impact Artifacts")
	command.SetHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		emb := uti.NewEmbed().
			SetTitle("Genshin Impact Artifacts").
			SetColor(0x00FF00)
		for k, v := range games.GenerateRandomArtifacts() {
			emb.AddField(k, fmt.Sprint(v)).InlineAllFields()
		}
		emb.RespondEmbed(s, i)
	})
}
