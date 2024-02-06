package utils

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

// ListFolderFiles returns a list of files in a folder.
func ListFiles(path string) ([]string, error) {
	var folderList []string

	outputDirRead, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open folder %s: %v", path, err)
	}

	defer outputDirRead.Close()

	outputDirFiles, err := outputDirRead.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("could not read files from folder %s: %v", path, err)
	}

	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		outputNameHere := outputFileHere.Name()
		folderList = append(folderList, outputNameHere)
	}
	return folderList, nil
}

// IsDirectory checks to see if a path corresponds to a directory.
func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	return fileInfo.IsDir()
}

// MapEmailHeaders structures mail data based on headers.
func MapEmailHeaders(key string, value string, emailStruct models.EnronMail) models.EnronMail {
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

	return emailStruct
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
				emailStructure = MapEmailHeaders(key, value, emailStructure)
			}
		}
	}
	// Add the body to the email structure
	emailStructure.Content = bodyLines.String()

	return emailStructure
}

// SendDataToIndex sends data to the index using HTTP.
func SendDataToIndex(data *[]models.EnronMail) error {
	bulkData := models.BulkDocument{
		Index:   "mail2",
		Records: *data,
	}

	jsonBody, err := json.Marshal(bulkData)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}
	req, err := http.NewRequest("POST", os.Getenv("ZS_BASE_URL")+"/api/_bulkv2", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.SetBasicAuth(os.Getenv("ZS_USER"), os.Getenv("ZS_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.Println("Bulk data succesfully inserted")
	}
	return nil
}

var DataBatch []models.EnronMail

// FillIndexBatch fills the batch of emails to be sent to the index.
func FillIndexBatch(path string) error {
	const batchSize = 1000
	jsonEmail := ConvertEmailFileToJSON(path)
	DataBatch = append(DataBatch, jsonEmail)

	if len(DataBatch) == batchSize {
		err := SendDataToIndex(&DataBatch)
		if err != nil {
			return fmt.Errorf("error while %s: %v", path, err)
		}
		DataBatch = nil // Clean bacth after sending the data
	}
	return nil
}

// IndexEmail recursively processes the directory structure and fills the index batch accordingly with FillIndexBatch(path string).
func IndexEmail(path string) error {
	if IsDirectory(path) {
		folders, err := ListFiles(path)
		if err != nil {
			return fmt.Errorf("error while listing files in path %s: %v", path, err)
		}
		for _, folder := range folders {
			err := IndexEmail(path + "/" + folder)
			if err != nil {
				return fmt.Errorf("error while indexing email in path %s: %v", path, err)
			}
		}
	} else {
		err := FillIndexBatch(path)
		if err != nil {
			return fmt.Errorf("error while  %s: %v", path, err)
		}
	}

	return nil
}
