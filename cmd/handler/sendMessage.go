package handler

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	RoleMessageID         string
	RankMassageID_Tank    string
	RankMassageID_DPS     string
	RankMassageID_Support string
	TankIcon              = "tank:1030393605350752276"
	DPSIcon               = "dps:1030393602054049842"
	SupportIcon           = "support:1030393603903721542"
	BronzeIcon            = "bronze:1030400694118785024"
	SilverIcon            = "silver:1030400687382732820"
	GoldIcon              = "gold:1030400684526419989"
	PlatinumIcon          = "platinum:1030400692818550794"
	DiamondIcon           = "diamond:1030400688762662993"
	MasterIcon            = "master:1030400695616147477"
	Top500Icon            = "top500:1030400679434522634"
)

func SendMessage(s *discordgo.Session, channelID string, msg string, emoji []string) {
	message, err := s.ChannelMessageSend(channelID, msg)
	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
	SendReaction(s, message.ChannelID, message.ID, emoji)
	//SelectMessage(s, message.ChannelID, message.ID)
	if msg == "チームの作成を行います。ロールを選択してください" {
		RoleMessageID = message.ID
	}

	if msg == "レートを入力してください [タンク]" {
		RankMassageID_Tank = message.ID
	}
	if msg == "レートを入力してください [DPS]" {
		RankMassageID_DPS = message.ID
	}
	if msg == "レートを入力してください [サポート]" {
		RankMassageID_Support = message.ID
	}
}

func SendReaction(s *discordgo.Session, channelID string, messageID string, emojiID []string) {
	for _, id := range emojiID {
		err := s.MessageReactionAdd(channelID, messageID, id)
		if err != nil {
			fmt.Println("EMOJI:", err)
		}
	}
}

func SelectMessage(s *discordgo.Session, channelID string, messageID string) {
	message, err := s.ChannelMessage(channelID, messageID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Message >>> [", message.Content, "]")
	/*
		message, err := s.Message
		if err != nil {
			fmt.Println("message Error:", err)
		}
		fmt.Println("MESSAGE!!!:", message)
	*/

}

/*
func GlobalTest(s *discordgo.Session, channelID string) {
	fmt.Println("[[[", RoleMessageID, RankMassageID, "]]]")
	//roleMessageID = "1032459037486350366"
	roleMessage, err := s.ChannelMessage(channelID, RoleMessageID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ロールメッセージ: ", roleMessage.Reactions[0].Emoji.Name, roleMessage.Reactions[0].Count, roleMessage.Member)
	//rankMassageID = "1032459044453109760"
	rankMassage, err := s.ChannelMessage(channelID, RankMassageID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ランクメッセージ: ", rankMassage)
}

func ReactionTest(s *discordgo.Session, channelID string) {
	user, err := s.MessageReactions(channelID, RoleMessageID, "tank:1030393605350752276", 10, "", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("[User]", user)

}
*/
