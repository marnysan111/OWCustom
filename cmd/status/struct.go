package status

import "github.com/bwmarrin/discordgo"

type User struct {
	ID   int
	Name string
	Role []string
	Rank []string
}

type Role struct {
	Tank    []*discordgo.User
	DPS     []*discordgo.User
	Support []*discordgo.User
}

type Rank struct {
	Bronze   []*discordgo.User
	Silver   []*discordgo.User
	Gold     []*discordgo.User
	Platinum []*discordgo.User
	Diamond  []*discordgo.User
	Master   []*discordgo.User
	Top500   []*discordgo.User
}
