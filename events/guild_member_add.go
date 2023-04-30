package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/go-pg/pg/v9"
	"github.com/savi2w/dorothy/models"
)

func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	user, err := models.GetStoredUser(m.User.ID)
	if err != nil {
		if err == pg.ErrNoRows {
			return
		}

		panic(err)
	}

	if err := s.GuildMemberRoleAdd(m.GuildID, m.User.ID, user.Role); err != nil {
		panic(err)
	}
}
