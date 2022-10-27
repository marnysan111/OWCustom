package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/marnysan111/OWCustom/handler"
)

var (
	GuildID = "929987034343432194"
	Role    = []string{"tank:1030393605350752276", "dps:1030393602054049842", "support:1030393603903721542"}
	Rank    = []string{"bronze:1030400694118785024", "silver:1030400687382732820", "gold:1030400684526419989", "platinum:1030400692818550794", "diamond:1030400688762662993", "master:1030400695616147477", "top500:1030400679434522634"}
	Join    = []string{""}
)

func main() {
	handler.CreateTeam()
	/*
		loadEnv()
		var (
			Token   = "Bot " + os.Getenv("APP_BOT_TOKEN")
			BotName = "<@" + os.Getenv("CLIENT_ID") + ">"
		)
		fmt.Println(BotName)

		discord, err := discordgo.New(Token)
		if err != nil {
			fmt.Println("ログインに失敗しました")
			fmt.Println(err)
		}
		//イベントハンドラを追加
		discord.AddHandler(onMessageCreate)

		err = discord.Open()
		if err != nil {
			fmt.Println("接続エラー: ", err)
			os.Exit(1)
		}

		// 直近の関数（main）の最後に実行される
		defer discord.Close()

		fmt.Println("Listening...")
		stopBot := make(chan os.Signal, 1)
		signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-stopBot
		return
	*/
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf(".env読み込みエラー: %v", err)
	}
	fmt.Println(".envを読み込みました。")
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	/*
		if len(Role) == 0 || len(Rank) == 0 {
			Role, Rank = handler.InitIconID(s, GuildID)
		}
	*/
	//clientId := os.Getenv("CLIENT_ID")
	u := m.Author
	fmt.Printf("%20s %20s(%20s) > %s\n", m.ChannelID, u.Username, u.ID, m.Content)
	switch {
	case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", "<@1027924810560450650>", "!join")):
		handler.SendMessage(s, m.ChannelID, "カスタムに参加する人は次のスタンプを押してください", Join)

	case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", "<@1027924810560450650>", "!team")):
		handler.SendMessage(s, m.ChannelID, "チームの作成を行います。ロールを選択してください", Role)
		handler.SendMessage(s, m.ChannelID, "レートを入力してください [タンク]", Rank)
		handler.SendMessage(s, m.ChannelID, "レートを入力してください [DPS]", Rank)
		handler.SendMessage(s, m.ChannelID, "レートを入力してください [サポート]", Rank)

	case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", "<@1027924810560450650>", "!create")): // or shuffle
		//handler.CreateTeam(s, m.ChannelID)

		/*
			case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", "<@1027924810560450650>", "!test")):
				//handler.GlobalTest(s, m.ChannelID)
				handler.ReactionTest(s, m.ChannelID)
		*/
	}
}

/*
func sendReply(s *discordgo.Session, channelID string, msg string, reference *discordgo.MessageReference) {
	_, err := s.ChannelMessageSendReply(channelID, msg, reference)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
*/
