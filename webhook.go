package discordwebhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Client représente un client webhook Discord
type Client struct {
	WebhookURL string
	Options    WebhookOptions
	httpClient *http.Client
}

// NewClient crée un nouveau client webhook
func NewClient(webhookURL string, options ...WebhookOptions) *Client {
	client := &Client{
		WebhookURL: webhookURL,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}

	if len(options) > 0 {
		client.Options = options[0]
		if client.Options.Proxy != nil {
			client.httpClient.Transport = &http.Transport{
				Proxy: http.ProxyURL(client.Options.Proxy),
			}
		}
	}

	return client
}

// SendMessage envoie un message simple
func (c *Client) SendMessage(content string) error {
	payload := DiscordPayload{
		Content:  content,
		Username: c.Options.Username,
		Avatar:   c.Options.Avatar,
	}

	return c.sendPayload(payload, "")
}

// SendEmbed envoie un embed
func (c *Client) SendEmbed(embed DiscordEmbed) error {
	payload := DiscordPayload{
		Embeds:   []DiscordEmbed{embed},
		Username: c.Options.Username,
		Avatar:   c.Options.Avatar,
	}

	return c.sendPayload(payload, "")
}

// SendEmbedWithFile envoie un embed avec un fichier
func (c *Client) SendEmbedWithFile(embed DiscordEmbed, filename string) error {
	payload := DiscordPayload{
		Embeds:   []DiscordEmbed{embed},
		Username: c.Options.Username,
		Avatar:   c.Options.Avatar,
	}

	return c.sendPayload(payload, filename)
}

// SendEmbedWithFile envoie un embed avec un fichier
func (c *Client) SendFile(filename string) error {
	payload := DiscordPayload{
		Username: c.Options.Username,
		Avatar:   c.Options.Avatar,
	}

	return c.sendPayload(payload, filename)
}

// SendCustomPayload envoie un payload personnalisé
func (c *Client) SendCustomPayload(payload DiscordPayload) error {
	return c.sendPayload(payload, "")
}

// SendCustomPayloadWithFile envoie un payload personnalisé avec un fichier
func (c *Client) SendCustomPayloadWithFile(payload DiscordPayload, filename string) error {
	return c.sendPayload(payload, filename)
}

func (c *Client) sendPayload(payload DiscordPayload, filename string) error {
	body, contentType, err := c.prepareRequest(payload, filename)
	if err != nil {
		return fmt.Errorf("failed to prepare request: %w", err)
	}

	return c.sendWebhookSafe(body, contentType)
}

func (c *Client) prepareRequest(payload DiscordPayload, filename string) (*bytes.Buffer, string, error) {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Ajouter le fichier si spécifié
	if filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			return nil, "", fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		part, err := writer.CreateFormFile("file", filepath.Base(filename))
		if err != nil {
			return nil, "", fmt.Errorf("failed to create form file: %w", err)
		}

		if _, err := io.Copy(part, file); err != nil {
			return nil, "", fmt.Errorf("failed to copy file: %w", err)
		}
	}

	// Ajouter le payload JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	if err := writer.WriteField("payload_json", string(payloadJSON)); err != nil {
		return nil, "", fmt.Errorf("failed to write payload field: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, "", fmt.Errorf("failed to close writer: %w", err)
	}

	return &requestBody, writer.FormDataContentType(), nil
}

func (c *Client) sendWebhookSafe(body *bytes.Buffer, contentType string) error {
	originalBody := bytes.NewReader(body.Bytes())

	for {
		// Reset reader for retry
		originalBody.Seek(0, 0)

		req, err := http.NewRequest("POST", c.WebhookURL, originalBody)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Content-Type", contentType)
		req.Header.Set("User-Agent", "DiscordWebhook-Go/1.0")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send request: %w", err)
		}

		if resp.StatusCode == 429 {
			var rateLimit struct {
				RetryAfter float64 `json:"retry_after"`
			}
			json.NewDecoder(resp.Body).Decode(&rateLimit)
			resp.Body.Close()

			wait := time.Duration(rateLimit.RetryAfter*1000) * time.Millisecond
			if wait == 0 {
				wait = 1 * time.Second
			}
			time.Sleep(wait)
			continue
		}

		resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		}

		return fmt.Errorf("webhook failed with status: %s", resp.Status)
	}
}
