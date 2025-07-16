package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	discordwebhook "github.com/peepsii/discord-webhook-go"
)

func main() {
	// Configuration from environment variables
	webhookURL := "https://discordapp.com/api/webhooks/"

	fmt.Println("Discord Webhook Go - Complete Example")
	fmt.Println("=====================================")

	// Example 1: Basic message
	basicMessageExample(webhookURL)

	// Example 2: Client with options
	clientWithOptionsExample(webhookURL)

	// Example 3: Simple embed
	simpleEmbedExample(webhookURL)

	// Example 4: Complete embed with all fields
	completeEmbedExample(webhookURL)

	// Example 5: Multiple embeds
	multipleEmbedsExample(webhookURL)

	// Example 6: File attachment
	fileAttachmentExample(webhookURL)

	// Example 7: File with embed
	fileWithEmbedExample(webhookURL)

	// Example 8: Custom payload
	customPayloadExample(webhookURL)

	// Example 9: Error notification
	errorNotificationExample(webhookURL)

	// Example 10: System monitoring
	systemMonitoringExample(webhookURL)

	// Example 11: Deployment notification
	deploymentNotificationExample(webhookURL)

	// Example 12: Proxy configuration
	proxyConfigurationExample(webhookURL)

	// Example 13: Rate limiting demonstration
	rateLimitingExample(webhookURL)

	// Example 14: Different color usage
	colorUsageExample(webhookURL)

	fmt.Println("\nAll examples completed!")
}

// Example 1: Basic message
func basicMessageExample(webhookURL string) {
	fmt.Println("\n1. Basic Message Example")
	client := discordwebhook.NewClient(webhookURL)

	err := client.SendMessage("Hello, Discord! This is a basic message.")
	if err != nil {
		log.Printf("Error sending basic message: %v", err)
	} else {
		fmt.Println("✓ Basic message sent successfully")
	}
}

// Example 2: Client with options
func clientWithOptionsExample(webhookURL string) {
	fmt.Println("\n2. Client with Options Example")

	client := discordwebhook.NewClient(
		webhookURL,
		discordwebhook.WebhookOptions{
			Username: "Example Bot",
			Avatar:   "https://cdn.discordapp.com/embed/avatars/0.png",
		},
	)

	err := client.SendMessage("This message is sent with custom username and avatar!")
	if err != nil {
		log.Printf("Error sending message with options: %v", err)
	} else {
		fmt.Println("✓ Message with options sent successfully")
	}
}

// Example 3: Simple embed
func simpleEmbedExample(webhookURL string) {
	fmt.Println("\n3. Simple Embed Example")
	client := discordwebhook.NewClient(webhookURL)

	embed := discordwebhook.DiscordEmbed{
		Title:       "Simple Embed",
		Description: "This is a simple embed with title and description.",
		Color:       0x3498db, // Blue
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	err := client.SendEmbed(embed)
	if err != nil {
		log.Printf("Error sending simple embed: %v", err)
	} else {
		fmt.Println("✓ Simple embed sent successfully")
	}
}

// Example 4: Complete embed with all fields
func completeEmbedExample(webhookURL string) {
	fmt.Println("\n4. Complete Embed Example")
	client := discordwebhook.NewClient(webhookURL)

	embed := discordwebhook.DiscordEmbed{
		Title:       "Complete Embed Example",
		Url:         "https://github.com",
		Description: "This embed demonstrates all available fields including author, fields, image, thumbnail, and footer.",
		Color:       0xe74c3c, // Red

		Author: discordwebhook.EmbedAuthor{
			Name:    "GitHub",
			URL:     "https://github.com",
			IconURL: "https://github.com/favicon.ico",
		},

		Fields: []discordwebhook.EmbedField{
			{
				Name:   "Inline Field 1",
				Value:  "This is an inline field",
				Inline: true,
			},
			{
				Name:   "Inline Field 2",
				Value:  "This is also inline",
				Inline: true,
			},
			{
				Name:   "Inline Field 3",
				Value:  "Third inline field",
				Inline: true,
			},
			{
				Name:   "Full Width Field",
				Value:  "This field spans the full width",
				Inline: false,
			},
		},

		Image: map[string]string{
			"url": "https://via.placeholder.com/400x200/3498db/ffffff?text=Main+Image",
		},

		Thumbnail: map[string]string{
			"url": "https://via.placeholder.com/80x80/e74c3c/ffffff?text=Thumb",
		},

		Footer: &discordwebhook.EmbedFooter{
			Text:    "Footer Text • Example App v1.0",
			IconURL: "https://via.placeholder.com/20x20/95a5a6/ffffff?text=F",
		},

		Timestamp: time.Now().Format(time.RFC3339),
	}

	err := client.SendEmbed(embed)
	if err != nil {
		log.Printf("Error sending complete embed: %v", err)
	} else {
		fmt.Println("✓ Complete embed sent successfully")
	}
}

// Example 5: Multiple embeds
func multipleEmbedsExample(webhookURL string) {
	fmt.Println("\n5. Multiple Embeds Example")
	client := discordwebhook.NewClient(webhookURL)

	payload := discordwebhook.DiscordPayload{
		Content: "Here are multiple embeds in a single message:",
		Embeds: []discordwebhook.DiscordEmbed{
			{
				Title:       "First Embed",
				Description: "This is the first embed",
				Color:       0x2ecc71, // Green
			},
			{
				Title:       "Second Embed",
				Description: "This is the second embed",
				Color:       0xf39c12, // Orange
			},
			{
				Title:       "Third Embed",
				Description: "This is the third embed",
				Color:       0x9b59b6, // Purple
			},
		},
	}

	err := client.SendCustomPayload(payload)
	if err != nil {
		log.Printf("Error sending multiple embeds: %v", err)
	} else {
		fmt.Println("✓ Multiple embeds sent successfully")
	}
}

// Example 6: File attachment
func fileAttachmentExample(webhookURL string) {
	fmt.Println("\n6. File Attachment Example")
	client := discordwebhook.NewClient(webhookURL)

	// Create a sample file
	sampleContent := `# Sample File
This is a sample text file created for the Discord webhook example.

## Features:
- Line 1
- Line 2
- Line 3

End of file.`

	err := os.WriteFile("sample.txt", []byte(sampleContent), 0644)
	if err != nil {
		log.Printf("Error creating sample file: %v", err)
		return
	}
	defer os.Remove("sample.txt")

	err = client.SendFile("sample.txt")
	if err != nil {
		log.Printf("Error sending file: %v", err)
	} else {
		fmt.Println("✓ File sent successfully")
	}
}

// Example 7: File with embed
func fileWithEmbedExample(webhookURL string) {
	fmt.Println("\n7. File with Embed Example")
	client := discordwebhook.NewClient(webhookURL)

	// Create a sample file
	reportContent := `Report Generated: ` + time.Now().Format("2006-01-02 15:04:05") + `
Status: OK
Users: 1,234
Uptime: 99.9%`

	err := os.WriteFile("report.txt", []byte(reportContent), 0644)
	if err != nil {
		log.Printf("Error creating report file: %v", err)
		return
	}
	defer os.Remove("report.txt") // Clean up

	embed := discordwebhook.DiscordEmbed{
		Title:       "System Report",
		Description: "Please find the detailed report in the attached file.",
		Color:       0x3498db,
		Fields: []discordwebhook.EmbedField{
			{
				Name:   "Report Type",
				Value:  "System Status",
				Inline: true,
			},
			{
				Name:   "Generated",
				Value:  time.Now().Format("15:04:05"),
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	err = client.SendEmbedWithFile(embed, "report.txt")
	if err != nil {
		log.Printf("Error sending file with embed: %v", err)
	} else {
		fmt.Println("✓ File with embed sent successfully")
	}
}

// Example 8: Custom payload
func customPayloadExample(webhookURL string) {
	fmt.Println("\n8. Custom Payload Example")
	client := discordwebhook.NewClient(webhookURL)

	payload := discordwebhook.DiscordPayload{
		Content:  "This is a custom payload with everything!",
		Username: "Custom Bot",
		Avatar:   "https://cdn.discordapp.com/embed/avatars/1.png",
		Embeds: []discordwebhook.DiscordEmbed{
			{
				Title:       "Custom Payload Embed",
				Description: "This embed is part of a custom payload",
				Color:       0x1abc9c,
				Fields: []discordwebhook.EmbedField{
					{
						Name:   "Payload Type",
						Value:  "Custom",
						Inline: true,
					},
					{
						Name:   "Components",
						Value:  "Message + Embed",
						Inline: true,
					},
				},
			},
		},
	}

	err := client.SendCustomPayload(payload)
	if err != nil {
		log.Printf("Error sending custom payload: %v", err)
	} else {
		fmt.Println("✓ Custom payload sent successfully")
	}
}

// Example 9: Error notification
func errorNotificationExample(webhookURL string) {
	fmt.Println("\n9. Error Notification Example")
	client := discordwebhook.NewClient(webhookURL)

	// Simulate an error
	simulatedError := fmt.Errorf("database connection failed: timeout after 30s")

	embed := discordwebhook.DiscordEmbed{
		Title:       "🚨 Critical Error",
		Description: fmt.Sprintf("```\n%v\n```", simulatedError),
		Color:       0xff0000, // Red
		Fields: []discordwebhook.EmbedField{
			{
				Name:   "Server",
				Value:  "web-server-01",
				Inline: true,
			},
			{
				Name:   "Environment",
				Value:  "Production",
				Inline: true,
			},
			{
				Name:   "Time",
				Value:  time.Now().Format("15:04:05"),
				Inline: true,
			},
			{
				Name:   "Severity",
				Value:  "Critical",
				Inline: true,
			},
		},
		Footer: &discordwebhook.EmbedFooter{
			Text: "Error Monitoring System",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	err := client.SendEmbed(embed)
	if err != nil {
		log.Printf("Error sending error notification: %v", err)
	} else {
		fmt.Println("✓ Error notification sent successfully")
	}
}

// Example 10: System monitoring
func systemMonitoringExample(webhookURL string) {
	fmt.Println("\n10. System Monitoring Example")
	client := discordwebhook.NewClient(webhookURL)

	// Simulate system metrics
	cpuUsage := 45.6
	memUsage := 78.2
	diskUsage := 34.1

	color := 0x00ff00 // Green
	status := "Healthy"

	if cpuUsage > 80 || memUsage > 80 || diskUsage > 80 {
		color = 0xff0000 // Red
		status = "Critical"
	} else if cpuUsage > 60 || memUsage > 60 || diskUsage > 60 {
		color = 0xffa500 // Orange
		status = "Warning"
	}

	embed := discordwebhook.DiscordEmbed{
		Title:       "System Status: " + status,
		Description: "Current system metrics",
		Color:       color,
		Fields: []discordwebhook.EmbedField{
			{
				Name:   "CPU Usage",
				Value:  fmt.Sprintf("%.1f%%", cpuUsage),
				Inline: true,
			},
			{
				Name:   "Memory Usage",
				Value:  fmt.Sprintf("%.1f%%", memUsage),
				Inline: true,
			},
			{
				Name:   "Disk Usage",
				Value:  fmt.Sprintf("%.1f%%", diskUsage),
				Inline: true,
			},
			{
				Name:   "Uptime",
				Value:  "5d 12h 34m",
				Inline: true,
			},
			{
				Name:   "Active Users",
				Value:  "1,234",
				Inline: true,
			},
			{
				Name:   "Response Time",
				Value:  "45ms",
				Inline: true,
			},
		},
		Footer: &discordwebhook.EmbedFooter{
			Text: "System Monitor v2.1",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	err := client.SendEmbed(embed)
	if err != nil {
		log.Printf("Error sending system monitoring: %v", err)
	} else {
		fmt.Println("✓ System monitoring sent successfully")
	}
}

// Example 11: Deployment notification
func deploymentNotificationExample(webhookURL string) {
	fmt.Println("\n11. Deployment Notification Example")
	client := discordwebhook.NewClient(webhookURL)

	version := "v1.2.3"
	success := true
	duration := "2m 34s"

	title := "🚀 Deployment Successful"
	color := 0x00ff00 // Green

	if !success {
		title = "❌ Deployment Failed"
		color = 0xff0000 // Red
	}

	embed := discordwebhook.DiscordEmbed{
		Title:       title,
		Description: fmt.Sprintf("Application version %s", version),
		Color:       color,
		Fields: []discordwebhook.EmbedField{
			{
				Name:   "Version",
				Value:  version,
				Inline: true,
			},
			{
				Name:   "Environment",
				Value:  "Production",
				Inline: true,
			},
			{
				Name:   "Duration",
				Value:  duration,
				Inline: true,
			},
			{
				Name:   "Branch",
				Value:  "main",
				Inline: true,
			},
			{
				Name:   "Commit",
				Value:  "abc123f",
				Inline: true,
			},
			{
				Name:   "Author",
				Value:  "john.doe",
				Inline: true,
			},
		},
		Footer: &discordwebhook.EmbedFooter{
			Text: "CI/CD Pipeline",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	err := client.SendEmbed(embed)
	if err != nil {
		log.Printf("Error sending deployment notification: %v", err)
	} else {
		fmt.Println("✓ Deployment notification sent successfully")
	}
}

// Example 12: Proxy configuration
func proxyConfigurationExample(webhookURL string) {
	fmt.Println("\n12. Proxy Configuration Example")

	// Note: This is just a demonstration - replace with actual proxy if needed
	proxyURL, err := url.Parse("http://proxy.example.com:8080")
	if err != nil {
		log.Printf("Error parsing proxy URL: %v", err)
		return
	}

	client := discordwebhook.NewClient(
		webhookURL,
		discordwebhook.WebhookOptions{
			Username: "Proxy Bot",
			Proxy:    proxyURL,
		},
	)

	// This would fail with the example proxy, but demonstrates the setup
	err = client.SendMessage("This message would be sent through a proxy (if configured)")
	if err != nil {
		fmt.Printf("⚠ Proxy example failed (expected with demo proxy): %v\n", err)
	} else {
		fmt.Println("✓ Proxy configuration working")
	}
}

// Example 13: Rate limiting demonstration
func rateLimitingExample(webhookURL string) {
	fmt.Println("\n13. Rate Limiting Example")
	client := discordwebhook.NewClient(webhookURL)

	fmt.Println("Sending multiple messages to demonstrate rate limiting...")

	for i := 1; i <= 5; i++ {
		err := client.SendMessage(fmt.Sprintf("Rate limit test message %d/5", i))
		if err != nil {
			log.Printf("Error sending message %d: %v", i, err)
		} else {
			fmt.Printf("✓ Message %d sent successfully\n", i)
		}

		// Small delay to show the rate limiting in action
		time.Sleep(500 * time.Millisecond)
	}
}

// Example 14: Different color usage
func colorUsageExample(webhookURL string) {
	fmt.Println("\n14. Color Usage Example")
	client := discordwebhook.NewClient(webhookURL)

	colors := []struct {
		name  string
		value int
	}{
		{"Red", 0xff0000},
		{"Green", 0x00ff00},
		{"Blue", 0x0000ff},
		{"Yellow", 0xffff00},
		{"Purple", 0x9b59b6},
		{"Orange", 0xe67e22},
		{"Pink", 0xe91e63},
		{"Gray", 0x95a5a6},
	}

	embed := discordwebhook.DiscordEmbed{
		Title:       "Color Palette",
		Description: "Different colors available for embeds",
		Color:       0x3498db,
		Fields:      []discordwebhook.EmbedField{},
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	for _, color := range colors {
		embed.Fields = append(embed.Fields, discordwebhook.EmbedField{
			Name:   color.name,
			Value:  fmt.Sprintf("0x%06x", color.value),
			Inline: true,
		})
	}

	err := client.SendEmbed(embed)
	if err != nil {
		log.Printf("Error sending color example: %v", err)
	} else {
		fmt.Println("✓ Color usage example sent successfully")
	}
}
