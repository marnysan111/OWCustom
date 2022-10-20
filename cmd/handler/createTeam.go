package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CreateTeam(s *discordgo.Session, channelID string) {
	role, err := GetMemberRole(s, channelID)
	if err != nil {
		fmt.Println(err)
	}
	tankRank, err := GetMemberRank(s, channelID, RankMassageID_Tank)
	if err != nil {
		fmt.Println(err)
	}
	dpsRank, err := GetMemberRank(s, channelID, RankMassageID_DPS)
	if err != nil {
		fmt.Println(err)
	}
	supportRank, err := GetMemberRank(s, channelID, RankMassageID_Support)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ロール", role)
	fmt.Println("タンクらんく", tankRank)
	fmt.Println("DPSランク", dpsRank)
	fmt.Println("サポランク", supportRank)
}
