package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/bannerify/bannerify-go"
)

func main() {
	apiKey := os.Getenv("BANNERIFY_API_KEY")
	templateID := os.Getenv("BANNERIFY_TEMPLATE_ID")

	if apiKey == "" || templateID == "" {
		fmt.Println("⚠️  Skipping integration tests - BANNERIFY_API_KEY or BANNERIFY_TEMPLATE_ID not set")
		return
	}

	fmt.Println("Running Bannerify Go SDK integration tests...\n")

	ctx := context.Background()
	failed := false

	// Test 1: Create client
	fmt.Print("Test 1: Creating client... ")
	client := bannerify.NewClient(apiKey)
	if client == nil {
		fmt.Println("✗ FAILED: Client is nil")
		os.Exit(1)
	}
	fmt.Println("✓ PASSED")

	// Test 2: Generate signed URL
	fmt.Print("Test 2: Generating signed URL... ")
	signedURL, err := client.GenerateImageSignedURL(templateID, &bannerify.CreateImageOptions{
		Modifications: []bannerify.Modification{
			{Name: "title", Text: "Integration Test"},
		},
	})
	if err != nil {
		fmt.Printf("✗ FAILED: %v\n", err)
		os.Exit(1)
	}
	if !strings.HasPrefix(signedURL, "https://api.bannerify.co/v1/templates/signedurl") {
		fmt.Println("✗ FAILED: Invalid URL format")
		os.Exit(1)
	}
	fmt.Println("✓ PASSED")
	fmt.Printf("   URL: %s...\n", signedURL[:80])

	// Test 3: Create image (JPEG)
	fmt.Print("Test 3: Creating image (JPEG format)... ")
	result := client.CreateImage(ctx, templateID, &bannerify.CreateImageOptions{
		Format: "jpeg",
		Modifications: []bannerify.Modification{
			{Name: "title", Text: "Go SDK JPEG Test"},
		},
	})
	if result.Error != nil {
		fmt.Printf("✗ FAILED: %s\n", result.Error.Message)
		failed = true
	} else if result.Result != nil {
		if len(result.Result) == 0 {
			fmt.Println("✗ FAILED: Empty response")
			failed = true
		} else {
			fmt.Println("✓ PASSED")
			fmt.Printf("   JPEG size: %d bytes\n", len(result.Result))
		}
	}

	// Test 4: Create image (PNG)
	fmt.Print("Test 4: Creating image (PNG format)... ")
	result = client.CreateImage(ctx, templateID, &bannerify.CreateImageOptions{
		Modifications: []bannerify.Modification{
			{Name: "title", Text: "Go SDK Test PNG"},
		},
	})
	if result.Error != nil {
		fmt.Printf("✗ FAILED: %s\n", result.Error.Message)
		failed = true
	} else if result.Result != nil {
		if len(result.Result) == 0 {
			fmt.Println("✗ FAILED: Empty response")
			failed = true
		} else {
			fmt.Println("✓ PASSED")
			fmt.Printf("   PNG size: %d bytes\n", len(result.Result))
		}
	}

	// Test 5: Error handling
	fmt.Print("Test 5: Testing error handling... ")
	result = client.CreateImage(ctx, "invalid_template_id", nil)
	if result.Error == nil {
		fmt.Println("✗ FAILED: Should have returned error")
		failed = true
	} else {
		fmt.Println("✓ PASSED")
		fmt.Printf("   Error code: %s\n", result.Error.Code)
	}

	if failed {
		fmt.Println("\n❌ Some tests failed")
		os.Exit(1)
	}

	fmt.Println("\n✅ All integration tests passed!")
}
