package utilities

import (
	"bufio"
	"indexer/models"
	"log"
	"os"
	"strings"
	"time"
)

func ConvertDateFormat(date string) (time.Time, error) {
	time, err := time.Parse("Mon, _2 Jan 2006 15:04:05 -0700 (MST)", date)

	return time, err
}

// MapEmailHeaders structures mail data based on headers.
func MapEmailHeaders(key string, value string, emailStruct models.EnronMail) (models.EnronMail, error) {
	switch key {
	case "Message-ID":
		emailStruct.MessageID = value
	case "Date":
		formatedDate, err := ConvertDateFormat(value)
		if err != nil {
			return emailStruct, err
		}
		emailStruct.Date = formatedDate
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

	return emailStruct, nil
}

// ConvertEmailFileToJSON converts an email file to JSON format.
func ConvertEmailFileToJSON(filePath string) models.EnronMail {
	var (
		bodyLines      strings.Builder
		emailStructure models.EnronMail
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
				emailStructure, _ = MapEmailHeaders(key, value, emailStructure)
			}
		}
	}
	// Add the body to the email structure
	emailStructure.Content = bodyLines.String()

	return emailStructure
}
