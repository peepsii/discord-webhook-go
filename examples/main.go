package main

import (
	"log"
	"time"

	discordwebhook "github.com/peepsii/discord-webhook-go"
)

func main() {
	// Créer un client avec des options
	client := discordwebhook.NewClient("https://discord.com/api/webhooks/...",
		discordwebhook.WebhookOptions{
			Username: "Mon Bot",
			Avatar:   "https://example.com/avatar.png",
		})

	// Envoyer un message simple
	err := client.SendMessage("Hello, world!")
	if err != nil {
		log.Fatal(err)
	}

	// Envoyer un embed
	embed := discordwebhook.DiscordEmbed{
		Title:       "Mon Embed",
		Description: "Description de l'embed",
		Color:       0x00ff00,
		Timestamp:   time.Now().Format(time.RFC3339),
		Fields: []discordwebhook.EmbedField{
			{
				Name:   "Champ 1",
				Value:  "Valeur 1",
				Inline: true,
			},
		},
	}

	err = client.SendEmbed(embed)
	if err != nil {
		log.Fatal(err)
	}
}
