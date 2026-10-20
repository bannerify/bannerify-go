# Bannerify Go SDK

Official Go SDK for [Bannerify](https://bannerify.co) - Generate images and PDFs at scale via API.

[![Go Reference](https://pkg.go.dev/badge/github.com/bannerify/bannerify-go.svg)](https://pkg.go.dev/github.com/bannerify/bannerify-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/bannerify/bannerify-go)](https://goreportcard.com/report/github.com/bannerify/bannerify-go)

## Installation

```bash
go get github.com/bannerify/bannerify-go
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bannerify/bannerify-go"
)

func ptr[T any](value T) *T { return &value }

func main() {
	client := bannerify.NewClient("your-api-key")

	result := client.CreateImage(context.Background(), "tpl_xxxxxxxxx", &bannerify.CreateImageOptions{
		Modifications: []bannerify.Modification{
			{Name: "title", Text: ptr("Hello World")},
			{Name: "subtitle", Text: ptr("From Go SDK")},
		},
	})

	if result.Error != nil {
		fmt.Printf("Error: %s\n", result.Error.Message)
		return
	}

	if err := os.WriteFile("output.png", result.Result, 0644); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Image created successfully!")
}
```

Every `Modification` field except `Name` is a pointer, so the examples below build them with the `ptr` helper from this program.

## Features

- 🚀 Simple, idiomatic Go API
- 🔒 Type-safe with full type definitions
- ⚡ Efficient with minimal dependencies
- 🎯 Result/Error pattern for explicit error handling
- 📝 Comprehensive documentation
- ✅ Well-tested

## Usage

### Creating Images

```go
// Generate PNG
png := client.CreateImage(ctx, "tpl_xxxxxxxxx", &bannerify.CreateImageOptions{
	Modifications: []bannerify.Modification{
		{Name: "title", Text: ptr("My Title")},
		{Name: "image", Src: ptr("https://example.com/image.jpg")},
	},
})

// Generate WebP
webp := client.CreateImage(ctx, "tpl_xxxxxxxxx", &bannerify.CreateImageOptions{
	Format: "webp",
	Modifications: []bannerify.Modification{
		{Name: "title", Text: ptr("My Title")},
	},
})

if png.Error == nil && webp.Error == nil {
	os.WriteFile("image.png", png.Result, 0644)
	os.WriteFile("image.webp", webp.Result, 0644)
}
```

### Creating PDFs

```go
result := client.CreatePDF(ctx, "tpl_xxxxxxxxx", []bannerify.Modification{
	{Name: "title", Text: ptr("Invoice #123")},
})

if result.Error == nil {
	os.WriteFile("invoice.pdf", result.Result, 0644)
}
```

### Generating Signed URLs

```go
signedURL, err := client.GenerateImageSignedURL("tpl_xxxxxxxxx", &bannerify.CreateImageOptions{
	Modifications: []bannerify.Modification{
		{Name: "title", Text: ptr("Dynamic Title")},
	},
	Format: "png",
})

if err == nil {
	fmt.Printf("<img src='%s' alt='Generated Image' />", signedURL)
}
```

### Rendering Videos

A video template renders as an MP4. Each second of the video counts as one generation, and as two at 60 fps.

```go
// Follow the render: queued, progress, then done or error
for event := range client.RenderVideo(ctx, "tpl_xxxxxxxxx", []bannerify.Modification{
	{Name: "title", Text: ptr("Hi")},
}) {
	switch event.Type {
	case "progress":
		fmt.Printf("%.0f%%\n", event.Progress*100)
	case "done":
		fmt.Println(event.URL)
	case "error":
		fmt.Println(event.Error)
	}
}

// Or wait for the URL of the MP4
url, err := client.RenderAndWaitVideo(ctx, "tpl_xxxxxxxxx", nil)
```

The client timeout does not apply to a render, because a video can wait in line for minutes. Cancel `ctx` to stop.

### Error Handling

```go
result := client.CreateImage(ctx, "tpl_xxxxxxxxx", nil)

if result.Error != nil {
	fmt.Printf("Error Code: %s\n", result.Error.Code)
	fmt.Printf("Message: %s\n", result.Error.Message)
	fmt.Printf("Docs: %s\n", result.Error.Docs)
	fmt.Printf("Request ID: %s\n", result.Error.RequestID)
} else {
	os.WriteFile("output.png", result.Result, 0644)
}
```

When the API answers with an error, `result.Error` holds the `code`, `message`, `docs`, and `requestId` of its error object. Send the request id when you report a problem. `*ErrorResponse` also implements `error`, so `fmt.Println(result.Error)` prints `BAD_REQUEST: Invalid api key (request req:abc123)`.

The SDK creates a few errors itself: `INVALID_FORMAT`, `JSON_ERROR`, `REQUEST_ERROR`, `READ_ERROR`, and `HTTP_ERROR` for a failed connection or for a response that holds no API error.

### Custom Configuration

```go
client := bannerify.NewClient("your-api-key",
	bannerify.WithTimeout(30*time.Second),
	bannerify.WithBaseURL("https://api.bannerify.co/v1"),
)
```

## API Reference

### NewClient

```go
func NewClient(apiKey string, opts ...Option) *Client
```

### CreateImage

```go
func (c *Client) CreateImage(ctx context.Context, templateID string, opts *CreateImageOptions) *Response
```

### CreatePDF

```go
func (c *Client) CreatePDF(ctx context.Context, templateID string, modifications []Modification) *Response
```

### RenderVideo

```go
func (c *Client) RenderVideo(ctx context.Context, templateID string, modifications []Modification) iter.Seq[RenderVideoEvent]
```

### RenderAndWaitVideo

```go
func (c *Client) RenderAndWaitVideo(ctx context.Context, templateID string, modifications []Modification) (string, error)
```

### GenerateImageSignedURL

```go
func (c *Client) GenerateImageSignedURL(templateID string, opts *CreateImageOptions) (string, error)
```

## Examples

### Generate Multiple Images

```go
products := []struct {
	Name  string
	Price string
}{
	{"Product A", "$29.99"},
	{"Product B", "$39.99"},
}

for i, product := range products {
	result := client.CreateImage(ctx, "tpl_product_banner", &bannerify.CreateImageOptions{
		Modifications: []bannerify.Modification{
			{Name: "product_name", Text: ptr(product.Name)},
			{Name: "price", Text: ptr(product.Price)},
		},
	})
	
	if result.Error == nil {
		os.WriteFile(fmt.Sprintf("banner_%d.png", i), result.Result, 0644)
	}
}
```

### Email Campaigns

```go
recipients := []struct{ Name string }{{"Ada"}, {"Linus"}}

for _, recipient := range recipients {
	signedURL, err := client.GenerateImageSignedURL("tpl_email_header", &bannerify.CreateImageOptions{
		Modifications: []bannerify.Modification{
			{Name: "name", Text: ptr(fmt.Sprintf("Hi, %s!", recipient.Name))},
		},
	})
	if err != nil {
		continue
	}
	// Use signedURL in email HTML
	fmt.Println(signedURL)
}
```

## Development

```bash
# Run tests
go test ./...

# Format code
go fmt ./...

# Lint
golangci-lint run
```

## Documentation

Full documentation available at [https://bannerify.co/docs/sdk/go/overview](https://bannerify.co/docs/sdk/go/overview)

## Support

- Documentation: https://bannerify.co/docs
- Issues: https://github.com/bannerify/bannerify-go/issues
- Email: support@bannerify.co

## License

MIT License - see LICENSE file for details
