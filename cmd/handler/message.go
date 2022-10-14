package handler

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func SendMessage(s *discordgo.Session, channelID string, msg string, emoji []string) {
	message, err := s.ChannelMessageSend(channelID, msg)
	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
	SendReaction(s, message.ChannelID, message.ID, emoji)
	SelectMessage(s, message.ChannelID, message.ID)
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
	/*
		message, err := s.Message
		if err != nil {
			fmt.Println("message Error:", err)
		}
		fmt.Println("MESSAGE!!!:", message)
	*/
}
