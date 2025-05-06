package bot

import (
	"fmt"
	"sync"

	"github.com/Luca-11/BotDiscord.git/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session
var userMessageCount = make(map[string]int)
var mutex = &sync.Mutex{}

func Start(){
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "cc" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "sv ?")
	}

	if m.Content == "657619320364335105" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Rooooooh tu vas la fermer Sixte !!!!")
	}

	if m.Author.ID == "707730579642253462" {
		mutex.Lock()
		count := userMessageCount[m.Author.ID]
		userMessageCount[m.Author.ID] = (count + 1) % 5
		mutex.Unlock()

		var response string
		switch count {
		case 0:
			response = "Est-ce que t'es arabe ?"
		case 1:
			response = "Je t'ai demandé si tu es arabe"
		case 2:
			response = "T'ES ARABE ?!"
		case 3:
			response = "EST-CE QUE TU ES ARABE ?!!"
		case 4:
			response = "AHHHH JE LE SAVAIS BOUGNOULE"
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, response)
	}

	if m.Author.ID == "657619320364335105" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "@everyone Boooooooooo regardez le c'est un POINTEUR ! hummmm il aime les enfants pointez le du doigt 🫵🫵🫵🫵🫵")
	}

}

