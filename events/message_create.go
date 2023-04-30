package events

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/go-pg/pg/v9"
	"github.com/savi2w/dorothy/faker"
	"github.com/savi2w/dorothy/models"
	"github.com/savi2w/dorothy/utils"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	p := "^#[a-fA-F0-9]{6,6}$"
	matched, err := regexp.MatchString(p, m.Content)
	if err != nil {
		panic(err)
	}

	if !matched {
		return
	}

	role := ""
	if user, err := models.GetStoredUser(m.Author.ID); err != nil {
		if err == pg.ErrNoRows {
			next, err := s.GuildRoleCreate(m.GuildID)
			if err != nil {
				panic(err)
			}

			if err := models.SetStoredUser(m.Author.ID, next.ID); err != nil {
				panic(err)
			}

			role = next.ID
		} else {
			panic(err)
		}
	} else {
		role = user.Role
	}

	if err := s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, role); err != nil {
		panic(err)
	}

	name, err := faker.GetProductName()
	if err != nil {
		panic(err)
	}

	color, err := utils.HexToInt(m.Content[1:])
	if err != nil {
		panic(err)
	}

	hoist := false
	perm := 0
	mention := false

	_, err = s.GuildRoleEdit(m.GuildID, role, name, color, hoist, perm, mention)
	if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("<@%v>, I used my nanocamo module to give you the color you asked me for, I hope you like it. :heart:", m.Author.ID)
	_, err = s.ChannelMessageSend(m.ChannelID, message)

	if err != nil {
		panic(err)
	}
}
