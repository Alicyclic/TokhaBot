package cmd

import "github.com/bwmarrin/discordgo"

func init() {
	command := AddCommand("wipe", "Wipe a channel")
	command.SetHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		channel, _ := s.ChannelDelete(i.ChannelID)
		s.GuildChannelCreateComplex(i.GuildID, discordgo.GuildChannelCreateData{
			Name:                 channel.Name,
			Type:                 channel.Type,
			ParentID:             channel.ParentID,
			Topic:                channel.Topic,
			NSFW:                 channel.NSFW,
			RateLimitPerUser:     channel.RateLimitPerUser,
			UserLimit:            channel.UserLimit,
			Position:             channel.Position,
			PermissionOverwrites: channel.PermissionOverwrites,
			Bitrate:              channel.Bitrate})
	})
}
