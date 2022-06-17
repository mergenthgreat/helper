package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/mergenemre/helperbot/packages"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var BotID string

func main() {
	packages.Connect()
	dg, err := discordgo.New("Bot " + packages.GetToken())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Başarıyla Açıldı ")
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotID = u.ID

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	<-sc
}

var col *mongo.Collection = packages.GetCollection(packages.DB, "commands")

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "" {
		return
	}
	if m.Content == "!unrealengine" {
		ok := contains(m.Member.Roles, "983731605363376150")
		if ok {
			s.ChannelMessageSend(m.ChannelID, "Unreal Engine Rolu Uzerinizde Bulunyor.")
			return
		}
		err := s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, "983731605363376150")

		if err != nil {
			_, err := s.ChannelMessageSend(m.ChannelID, "Rol Verilirken bir Hata oluştu")
			if err != nil {
				fmt.Println("error")
			}
		}
		s.ChannelMessageSend(m.ChannelID, "Başarıyla Unreal Engine Rolu Verildi.")
	}
	var results packages.Commands
	filter := bson.D{primitive.E{Key: "name", Value: m.Content}}
	err := col.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		return
	}
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+", "+results.Content)
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
