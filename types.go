package discordwebhook

import "net/url"

// DiscordEmbed représente un embed Discord
type DiscordEmbed struct {
	Title       string            `json:"title,omitempty"`
	Url         string            `json:"url,omitempty"`
	Description string            `json:"description,omitempty"`
	Color       int               `json:"color,omitempty"`
	Image       map[string]string `json:"image,omitempty"`
	Fields      []EmbedField      `json:"fields,omitempty"`
	Footer      *EmbedFooter      `json:"footer,omitempty"`
	Timestamp   string            `json:"timestamp,omitempty"`
	Thumbnail   map[string]string `json:"thumbnail,omitempty"`
	Author      EmbedAuthor       `json:"author,omitempty"`
}

// EmbedField représente un champ d'embed
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type EmbedAuthor struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url,omitempty"`
}

// EmbedFooter représente le footer d'un embed
type EmbedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

// DiscordPayload représente le payload complet à envoyer
type DiscordPayload struct {
	Embeds   []DiscordEmbed `json:"embeds,omitempty"`
	Username string         `json:"username,omitempty"`
	Avatar   string         `json:"avatar_url,omitempty"`
	Content  string         `json:"content,omitempty"`
	TTS      bool           `json:"tts,omitempty"`
}

// WebhookOptions configure les options du webhook
type WebhookOptions struct {
	Username string
	Avatar   string
	Proxy    *url.URL
}
