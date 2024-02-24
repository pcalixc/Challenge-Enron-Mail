package utilsv2

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"indexer/models"
	"log"
	"net/http"
	"os"
	"strings"
)

// Convert an email file to a Models data structure
func ConvertEmailFileToStruct(filePath string) (*models.EnronMail, error) {
	var (
		bodyLines      strings.Builder
		emailStructure = &models.EnronMail{}
		bodyStarted    bool
	)

	// We read the email file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// We read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// If the line is empty, it indicates the start of the email body
		if line == "" {
			bodyStarted = true
			continue
		}

		if bodyStarted {
			// Store body lines in a slice
			bodyLines.WriteString(line)
			bodyLines.WriteString("\n")
		} else {
			// Parse email headers
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				key := parts[0]
				value := parts[1]
				// Create an Email object
				MapEmailHeaders(key, value, emailStructure)
				if err != nil {
					return emailStructure, fmt.Errorf("error al analizar encabezado: %v", err)
				}
			}
		}
	}
	// Add the body to the email structure
	emailStructure.Content = bodyLines.String()

	return emailStructure, nil
}

// MapEmailHeaders structures mail data based on headers.
func MapEmailHeaders(key string, value string, emailStruct *models.EnronMail) {
	switch key {
	case "Message-ID":
		emailStruct.MessageID = value
	case "Date":
		emailStruct.Date = value
	case "From":
		emailStruct.From = value
	case "To":
		emailStruct.To = value
	case "Subject":
		emailStruct.Subject = value
	case "Mime-Version":
		emailStruct.MimeVersion = value
	case "Content-Type":
		emailStruct.ContentType = value
	case "Content-Transfer-Encoding":
		emailStruct.ContentTransferEncoding = value
	case "X-From":
		emailStruct.XFrom = value
	case "X-To":
		emailStruct.XTo = value
	case "X-cc":
		emailStruct.XCc = value
	case "X-bcc":
		emailStruct.XBcc = value
	case "X-Folder":
		emailStruct.XFolder = value
	case "X-Origin":
		emailStruct.XOrigin = value
	case "X-FileName":
		emailStruct.XFileName = value
	}
}

// SendDataToIndex sends data to the index using HTTP.
func SendDataToIndex(data *[]models.EnronMail) error {
	bulkData := models.BulkDocument{
		Index:   "mail333",
		Records: *data,
	}

	jsonBody, err := json.Marshal(bulkData)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", os.Getenv("ZS_BASE_URL")+"/api/_bulkv2", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.SetBasicAuth(os.Getenv("ZS_USER"), os.Getenv("ZS_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		log.Println("Bulk data succesfully inserted")
	default:
		return fmt.Errorf("error indexing file. Error code:  %d", resp.StatusCode)
	}

	return nil
}
