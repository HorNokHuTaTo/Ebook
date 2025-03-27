package main

import (
	"fmt"
	"log"
	"os"

	"github.com/otiai10/gosseract"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor" // Import the extractor package from unipdf
	"github.com/unidoc/unipdf/v3/model"     // Use unipdf to access NewReader, pdf functionality
)

func main() {
	// Load the license for unipdf (apply your license key and license ID)
	err := license.SetLicenseKey("your-license-key", "your-license-id")
	if err != nil {
		log.Fatal(err)
	}

	// Open the PDF file
	file, err := os.Open("input.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Load the PDF document
	pdfReader, err := model.NewPdfReader(file)
	if err != nil {
		log.Fatal(err)
	}

	// Create an OCR client (Tesseract)
	client := gosseract.Newdsdsdsdsdsd
	defer client.Close()

	// Extract text and images from each page
	for i := 0; i < pdfReader.NumPages(); i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			log.Fatal(err)
		}

		// Extract text using extractor
		extractor, err := extractor.New(page)
		if err != nil {
			log.Fatal(err)
		}

		text, err := extractor.ExtractText()
		if err != nil {
			log.Fatal(err)
		}

		// Print extracted text from the page
		fmt.Println("Page Text:", text)

		// Extract images from the PDF page and run OCR if needed
		// Note: You will need a method to extract and process images
		// The unipdf package itself doesn't provide image extraction directly.
		// You can convert pages to images or use other external tools (like ImageMagick or Ghostscript)
		// Example of using OCR on an image:
		// For now, this is a placeholder for OCR on images (assuming you have images from the PDF)
		// ocrText, err := client.TextFromImage("image.png")
		// if err == nil {
		//     fmt.Println("OCR Text:", ocrText)
		// }
	}
}
