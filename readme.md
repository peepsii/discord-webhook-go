# Discord Webhook Go

[![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/peepsii/discord-webhook-go)](https://goreportcard.com/report/github.com/peepsii/discord-webhook-go)
[![GoDoc](https://godoc.org/github.com/peespii/discord-webhook-go?status.svg)](https://godoc.org/github.com/peepsii/discord-webhook-go)

A simple, efficient and complete Go library for sending Discord webhooks with support for embeds, files, and automatic rate limiting.

## Features

- **Simple messages** - Send text messages
- **Discord embeds** - Full embed support with fields, images, etc.
- **File attachments** - Attach files to messages
- **Automatic rate limiting** - Intelligent handling of Discord limits
- **Proxy support** - Compatible with HTTP/HTTPS proxies
- **Custom payloads** - Full control over sent content
- **Robust error handling** - Automatic retry and error management

## Installation

```bash
go get github.com/peepsii/discord-webhook-go
```

## Examples
- See https://github.com/peepsii/discord-webhook-go/blob/main/examples/main.go for a full demonstration.