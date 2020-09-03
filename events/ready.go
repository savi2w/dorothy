package events

import "github.com/bwmarrin/discordgo"

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	if err := s.UpdateStatus(1, "Model Warrior Jullianne"); err != nil {
		panic(err)
	}
}
