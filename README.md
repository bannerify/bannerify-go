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

func main() {
	client := bannerify.NewClient("your-api-key")

	result := client.CreateImage(context.Background(), "tpl_xxxxxxxxx", &bannerify.CreateImageOptions{
		Modifications: []bannerify.Modification{
			{Name: "title", Text: "Hello World"},
			{Name: "subtitle", Text: "From Go SDK"},
		},
	})

	if result.Error != nil {
		fmt.Printf("Error: %s\n", result.Error.Message)
		return
	}

	os.WriteFile("output.png", result.Result, 0644)
	fmt.Println("Image created successfully!")
}
```

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
result := client.CreateImage(ctx, "tpl_xxxxxxxxx", &bannerify.CreateImageOptions{
	Modifications: []bannerify.Modification{
		{Name: "title", Text: "My Title"},
		{Name: "image", Src: "https://example.com/image.jpg"},
	},
})

// Generate WebP
result := client.CreateImage(ctx, "tpl_xxxxxxxxx", &bannerify.CreateImageOptions{
	Format: "webp",
	Modifications: []bannerify.Modification{
		{Name: "title", Text: "My Title"},
	},
})

// Generate thumbnail
result := client.CreateImage(ctx, "tpl_xxxxxxxxx", &bannerify.CreateImageOptions{
	Thumbnail: true,
})
```

### Creating PDFs

```go
result := client.CreatePDF(ctx, "tpl_xxxxxxxxx", []bannerify.Modification{
	{Name: "title", Text: "Invoice #123"},
})

if result.Error == nil {
	os.WriteFile("invoice.pdf", result.Result, 0644)
}
```

### Generating Signed URLs

```go
signedURL, err := client.GenerateImageSignedURL("tpl_xxxxxxxxx", &bannerify.CreateImageOptions{
	Modifications: []bannerify.Modification{
		{Name: "title", Text: "Dynamic Title"},
	},
	Format: "png",
})

if err == nil {
	fmt.Printf("<img src='%s' alt='Generated Image' />", signedURL)
}
```

### Error Handling

```go
result := client.CreateImage(ctx, "tpl_xxxxxxxxx", nil)

if result.Error != nil {
	fmt.Printf("Error Code: %s\n", result.Error.Code)
	fmt.Printf("Message: %s\n", result.Error.Message)
	fmt.Printf("Docs: %s\n", result.Error.Docs)
} else {
	os.WriteFile("output.png", result.Result, 0644)
}
```

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
			{Name: "product_name", Text: product.Name},
			{Name: "price", Text: product.Price},
		},
	})
	
	if result.Error == nil {
		os.WriteFile(fmt.Sprintf("banner_%d.png", i), result.Result, 0644)
	}
}
```

### Email Campaigns

```go
for _, recipient := range recipients {
	signedURL, _ := client.GenerateImageSignedURL("tpl_email_header", &bannerify.CreateImageOptions{
		Modifications: []bannerify.Modification{
			{Name: "name", Text: fmt.Sprintf("Hi, %s!", recipient.Name)},
		},
	})
	// Use signedURL in email HTML
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
