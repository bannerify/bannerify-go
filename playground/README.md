# Go SDK Playground

Test the Bannerify Go SDK locally before publishing.

## Setup

1. Make sure you have the API key in `.env`:
   ```bash
   API_KEY=your_api_key_here
   ```

2. Run the test:
   ```bash
   go run main.go
   ```

   Or build and run:
   ```bash
   go build -o test
   ./test
   ```

## What it tests

- ✅ Create image with modifications
- ✅ Create image with type-safe modifications
- ✅ Create SVG image
- ✅ Generate signed URL
- ✅ Error handling

Results will be saved in `playground/output/`.

