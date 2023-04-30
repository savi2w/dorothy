package main

import (
	"os"
	"sync"

	"github.com/savi2w/dorothy/events"
	"github.com/savi2w/dorothy/models"

	"github.com/bwmarrin/discordgo"
)

func main() {
	db := models.InitDB()
	defer db.Close()

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		panic(err)
	}

	dg.AddHandler(events.GuildMemberAdd)
	dg.AddHandler(events.MessageCreate)
	dg.AddHandler(events.Ready)

	if err := dg.Open(); err != nil {
		panic(err)
	}

	defer dg.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
