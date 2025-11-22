// Playground test for Bannerify Go SDK
//
// Run with: go run playground/main.go
package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bannerify/bannerify-go"
)

func main() {
	// Load environment variables
	if err := loadEnv("../.env"); err != nil {
		fmt.Printf("❌ Error loading .env: %v\n", err)
		os.Exit(1)
	}

	apiKey := os.Getenv("API_KEY")
	templateID := os.Getenv("TEMPLATE_ID")
	if templateID == "" {
		templateID = "tpl_xxxxxxxxx"
	}

	if apiKey == "" {
		fmt.Println("❌ API_KEY not found in .env file")
		os.Exit(1)
	}

	fmt.Println("🎨 Bannerify Go SDK Playground\n")

	// Create output directory
	outputDir := "playground/output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("❌ Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Initialize client
	client := bannerify.NewClient(apiKey)
	fmt.Println("✅ Client initialized")

	ctx := context.Background()

	// Test 1: Create image with modifications
	fmt.Println("\n1️⃣ Test: Create image with modifications")
	result := client.CreateImage(ctx, templateID, &bannerify.CreateImageOptions{
		Modifications: []bannerify.Modification{
			{Name: "title", Text: strPtr("Go SDK Test")},
			{Name: "subtitle", Text: strPtr("Using Go types")},
		},
		Format: "png",
	})

	if result.Error != nil {
		fmt.Printf("❌ Error: %s\n", result.Error.Message)
	} else {
		outputPath := filepath.Join(outputDir, "test-basic.png")
		if err := os.WriteFile(outputPath, result.Result, 0644); err != nil {
			fmt.Printf("❌ Error writing file: %v\n", err)
		} else {
			fmt.Printf("✅ Image created: %s\n", outputPath)
			fmt.Printf("   Size: %d bytes\n", len(result.Result))
		}
	}

	// Test 2: Create image with type-safe modifications
	fmt.Println("\n2️⃣ Test: Create image with type-safe modifications")
	result = client.CreateImage(ctx, templateID, &bannerify.CreateImageOptions{
		Modifications: []bannerify.Modification{
			{Name: "title", Text: strPtr("Type-Safe Go"), Color: strPtr("#00ADD8")},
			{Name: "subtitle", Text: strPtr("Using typed structs"), Visible: boolPtr(true)},
		},
		Format: "png",
	})

	if result.Error != nil {
		fmt.Printf("❌ Error: %s\n", result.Error.Message)
	} else {
		outputPath := filepath.Join(outputDir, "test-typed.png")
		if err := os.WriteFile(outputPath, result.Result, 0644); err != nil {
			fmt.Printf("❌ Error writing file: %v\n", err)
		} else {
			fmt.Printf("✅ Image created: %s\n", outputPath)
			fmt.Printf("   Size: %d bytes\n", len(result.Result))
		}
	}

	// Test 3: Create JPEG
	fmt.Println("\n3️⃣ Test: Create JPEG image")
	result = client.CreateImage(ctx, templateID, &bannerify.CreateImageOptions{
		Format: "jpeg",
		Modifications: []bannerify.Modification{
			{Name: "title", Text: strPtr("JPEG Output")},
		},
	})

	if result.Error != nil {
		fmt.Printf("❌ Error: %s\n", result.Error.Message)
	} else {
		outputPath := filepath.Join(outputDir, "test-jpeg.jpg")
		if err := os.WriteFile(outputPath, result.Result, 0644); err != nil {
			fmt.Printf("❌ Error writing file: %v\n", err)
		} else {
			fmt.Printf("✅ JPEG created: %s\n", outputPath)
			fmt.Printf("   Size: %d bytes\n", len(result.Result))
		}
	}

	// Test 4: Generate signed URL
	fmt.Println("\n4️⃣ Test: Generate signed URL")
	signedURL, err := client.GenerateImageSignedURL(templateID, &bannerify.CreateImageOptions{
		Modifications: []bannerify.Modification{
			{Name: "title", Text: strPtr("Signed URL Test")},
		},
	})

	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Signed URL generated")
		if len(signedURL) > 80 {
			fmt.Printf("   URL: %s...\n", signedURL[:80])
		} else {
			fmt.Printf("   URL: %s\n", signedURL)
		}
	}

	// Test 5: Error handling
	fmt.Println("\n5️⃣ Test: Error handling")
	result = client.CreateImage(ctx, "invalid_template_id", nil)
	if result.Error != nil {
		fmt.Println("✅ Error handling works")
		fmt.Printf("   Code: %s\n", result.Error.Code)
		fmt.Printf("   Message: %s\n", result.Error.Message)
	} else {
		fmt.Println("❌ Should have returned error")
	}

	fmt.Println("\n✨ All tests completed! Check the output/ directory for generated images.")
}

// Helper functions
func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func loadEnv(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(parts[0], parts[1])
		}
	}

	return scanner.Err()
}

