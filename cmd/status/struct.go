package status

import "github.com/bwmarrin/discordgo"

type User struct {
	Name       string
	TankRank   string
	DPSRank    string
	SupprtRank string
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

type Role_Icon struct {
	Tank    string
	DPS     string
	Support string
}

type TmpRole struct {
	Tank    []string
	DPS     []string
	Support []string
}

type TmpRank struct {
	Bronze   []string
	Silver   []string
	Gold     []string
	Platinum []string
	Diamond  []string
	Master   []string
	Top500   []string
}
