package handler

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

/*
tank = 1034655699935952947
dps = 1034655717598171146
sup = 1034655739362422805
*/

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
