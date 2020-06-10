package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/awpby/discordgo"
)

// Variables used for command line parameters
var (
	Token      string
	UserID string
)

func init() {
	flag.StringVar(&Token, "t", "", "User Token")
	flag.StringVar(&UserID, "id", "", "User id")
	flag.Parse()

	if Token == "" || UserID == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	dg, err := discordgo.New(Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, event *discordgo.Ready) {
		profile, err := s.UserProfile(UserID)
		if err != nil {
			fmt.Println("UserProfile error:", err.Error())
			return
		}
		fmt.Println(profile)
	})

	if err := dg.Open(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
