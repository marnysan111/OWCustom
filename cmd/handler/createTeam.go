package handler

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/marnysan111/OWCustom/item"
	"github.com/marnysan111/OWCustom/status"
)

var (
	RoleMember    []status.Role
	TankMember    []status.Rank
	DPSMember     []status.Rank
	SupportMember []status.Rank
)

func CreateTeam( /*s *discordgo.Session,  channelID string*/ ) {
	//role, tank, dps, sup := GetMember(s, channelID)
	role, _, _, _ := TmpMember()
	fmt.Println(role)
	member := append(role.Tank, role.DPS...)
	member = append(member, role.Support...)
	member = SliceUnique(member)
	fmt.Println(member, ":", len(member))
	minRole := role.Tank
	randRole := role.Tank
	minRoleStr := "tank"
	if len(minRole) >= len(role.DPS) {
		minRole = role.DPS
		randRole = role.DPS
		minRoleStr = "dps"
	}
	if len(minRole) >= len(role.Support) {
		minRole = role.Support
		randRole = role.Support
		minRoleStr = "support"
	}
	rand.Seed(time.Now().UnixNano())
	target := randRole[rand.Intn(len(minRole))]
	fmt.Println(target, minRoleStr)
	item.DeleteSlice(member, target)
}

func GetMember(s *discordgo.Session, channelID string) (status.Role, status.Rank, status.Rank, status.Rank) {

	RoleMember, err := GetMemberRole(s, channelID)
	if err != nil {
		fmt.Println(err)
	}
	TankMember, err := GetMemberRank(s, channelID, RankMassageID_Tank)
	if err != nil {
		fmt.Println(err)
	}
	DPSMember, err := GetMemberRank(s, channelID, RankMassageID_DPS)
	if err != nil {
		fmt.Println(err)
	}
	SupportMember, err := GetMemberRank(s, channelID, RankMassageID_Support)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(RoleMember)
	fmt.Println(TankMember)
	fmt.Println(DPSMember)

	fmt.Println(SupportMember)
	return RoleMember, TankMember, DPSMember, SupportMember
}

func TmpMember() (status.TmpRole, status.TmpRank, status.TmpRank, status.TmpRank) {

	role := status.TmpRole{
		Tank:    []string{"mani", "茶釜", "げんじ", "モイラ"},
		DPS:     []string{"げんじ", "トレ", "ソンブラ", "ハンゾー", "茶釜"},
		Support: []string{"あな", "ルシオ", "ぶり", "モイラ", "mani"},
	}

	tank := status.TmpRank{
		Bronze:   []string{"げんじ"},
		Gold:     []string{"モイラ"},
		Platinum: []string{"mani"},
		Top500:   []string{"茶釜"},
	}
	dps := status.TmpRank{
		Bronze:  []string{"げんじ"},
		Silver:  []string{"ソンブラ"},
		Diamond: []string{"ハンゾー"},
		Master:  []string{"トレ", "茶釜"},
	}
	sup := status.TmpRank{
		Bronze:   []string{"モイラ"},
		Silver:   []string{"mani"},
		Gold:     []string{"あな", "ルシオ"},
		Platinum: []string{"ぶり"},
	}
	return role, tank, dps, sup
}

func SliceUnique(t []string) (unique []string) {
	m := map[string]bool{}
	for _, v := range t {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}
