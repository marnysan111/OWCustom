package handler

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

/*
	TankIcon              string
	DPSIcon               string
	SupportIcon           string
	BronzeIcon            string
	SilverIcon            string
	GoldIcon              string
	PlatinumIcon          string
	DiamondIcon           string
	MasterIcon            string
	Top500Icon            string
)
func InitIconID(s *discordgo.Session, guildID string) ([]string, []string) {
	guildEmojis, err := s.GuildEmojis(guildID)
	if err != nil {
		fmt.Println(err)
	}
	var role []string
	var rank []string

	for i := 0; i < len(guildEmojis); i++ {
		fmt.Println(guildEmojis[i])
		switch {
		case guildEmojis[i].Name == "tank_icon":
			role = append(role, guildEmojis[i].Name+":"+guildEmojis[i].ID)
			TankIcon = guildEmojis[i].Name + ":" + guildEmojis[i].ID
		case guildEmojis[i].Name == "dps_icon":
			role = append(role, guildEmojis[i].Name+":"+guildEmojis[i].ID)
			DPSIcon = guildEmojis[i].Name + ":" + guildEmojis[i].ID
		case guildEmojis[i].Name == "support_icon":
			role = append(role, guildEmojis[i].Name+":"+guildEmojis[i].ID)
			SupportIcon = guildEmojis[i].Name + ":" + guildEmojis[i].ID
		case guildEmojis[i].Name == "bronze":
			rank = append(rank, guildEmojis[i].Name+":"+guildEmojis[i].ID)
			BronzeIcon = guildEmojis[i].Name + ":" + guildEmojis[i].ID
		case guildEmojis[i].Name == "silver":
			rank = append(rank, guildEmojis[i].Name+":"+guildEmojis[i].ID)
			SilverIcon = guildEmojis[i].Name + ":" + guildEmojis[i].ID
		case guildEmojis[i].Name == "gold":
			rank = append(rank, guildEmojis[i].Name+":"+guildEmojis[i].ID)
			GoldIcon = guildEmojis[i].Name + ":" + guildEmojis[i].ID
		case guildEmojis[i].Name == "platinum":
			rank = append(rank, guildEmojis[i].Name+":"+guildEmojis[i].ID)
			PlatinumIcon = guildEmojis[i].Name + ":" + guildEmojis[i].ID
		case guildEmojis[i].Name == "diamond":
			rank = append(rank, guildEmojis[i].Name+":"+guildEmojis[i].ID)
			DiamondIcon = guildEmojis[i].Name + ":" + guildEmojis[i].ID
		case guildEmojis[i].Name == "master":
			rank = append(rank, guildEmojis[i].Name+":"+guildEmojis[i].ID)
			MasterIcon = guildEmojis[i].Name + ":" + guildEmojis[i].ID
		case guildEmojis[i].Name == "top500":
			rank = append(rank, guildEmojis[i].Name+":"+guildEmojis[i].ID)
			Top500Icon = guildEmojis[i].Name + ":" + guildEmojis[i].ID
		}
	}

	fmt.Println(role)
	fmt.Println(rank)
	return role, rank
}
*/
