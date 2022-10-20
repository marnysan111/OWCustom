package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/marnysan111/OWCustom/status"
)

func GetMemberRole(s *discordgo.Session, channelID string) (status.Role, error) {
	tank, err := s.MessageReactions(channelID, RoleMessageID, TankIcon, 10, "", "")
	if err != nil {
		return status.Role{}, fmt.Errorf("get tank error: %w", err)
	}

	dps, err := s.MessageReactions(channelID, RoleMessageID, DPSIcon, 10, "", "")
	if err != nil {
		return status.Role{}, fmt.Errorf("get dps error: %w", err)
	}

	support, err := s.MessageReactions(channelID, RoleMessageID, SupportIcon, 10, "", "")
	if err != nil {
		return status.Role{}, fmt.Errorf("get support error: %w", err)
	}
	role := status.Role{
		Tank:    tank,
		DPS:     dps,
		Support: support,
	}
	return role, nil
}

func GetMemberRank(s *discordgo.Session, channelID string, messageID string) (status.Rank, error) {
	bronze, err := s.MessageReactions(channelID, messageID, BronzeIcon, 10, "", "")
	if err != nil {
		return status.Rank{}, fmt.Errorf("f1 fail: %w", err)
	}
	silver, err := s.MessageReactions(channelID, messageID, SilverIcon, 10, "", "")
	if err != nil {
		return status.Rank{}, fmt.Errorf("f1 fail: %w", err)
	}
	gold, err := s.MessageReactions(channelID, messageID, GoldIcon, 10, "", "")
	if err != nil {
		return status.Rank{}, fmt.Errorf("f1 fail: %w", err)
	}
	platinum, err := s.MessageReactions(channelID, messageID, PlatinumIcon, 10, "", "")
	if err != nil {
		return status.Rank{}, fmt.Errorf("f1 fail: %w", err)
	}
	diamond, err := s.MessageReactions(channelID, messageID, DiamondIcon, 10, "", "")
	if err != nil {
		return status.Rank{}, fmt.Errorf("f1 fail: %w", err)
	}
	master, err := s.MessageReactions(channelID, messageID, MasterIcon, 10, "", "")
	if err != nil {
		return status.Rank{}, fmt.Errorf("f1 fail: %w", err)
	}
	top500, err := s.MessageReactions(channelID, messageID, Top500Icon, 10, "", "")
	if err != nil {
		return status.Rank{}, fmt.Errorf("f1 fail: %w", err)
	}
	rank := status.Rank{
		Bronze:   bronze,
		Silver:   silver,
		Gold:     gold,
		Platinum: platinum,
		Diamond:  diamond,
		Master:   master,
		Top500:   top500,
	}
	return rank, nil
}
