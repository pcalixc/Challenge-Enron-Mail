package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"indexer/models"
	"log"
	"net/http"
	"os"
)

var DataBatch []models.EnronMail

const batchSize = 1000

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
		outputFile := outputDirFiles[outputIndex]

		outputNameHere := outputFile.Name()
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

// SendDataToIndex sends data to the index using HTTP.
func SendDataToIndex(data *[]models.EnronMail) error {
	bulkData := models.BulkDocument{
		Index:   "mailfff",
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

	if resp.StatusCode == 200 {
		log.Println("Bulk data succesfully inserted")
	}
	return nil
}

// FillIndexBatch fills the batch of emails to be sent to the index.
func FillIndexBatch(path string) error {
	jsonEmail := ConvertEmailFileToJSON(path)
	DataBatch = append(DataBatch, jsonEmail)
	return nil
}

// IndexEmailDirectory recursively processes the directory structure and fills the index batch accordingly with FillIndexBatch(path string).
func IndexEmailDirectory(path string) error {
	if IsDirectory(path) {
		folders, err := ListFiles(path)
		if err != nil {
			return fmt.Errorf("error while listing files in path %s: %v", path, err)
		}
		for _, folder := range folders {
			err := IndexEmailDirectory(path + "/" + folder)
			if err != nil {
				return fmt.Errorf("error while indexing email in path %s: %v", path, err)
			}
		}
	} else {
		err := FillIndexBatch(path)
		if err != nil {
			return fmt.Errorf("error while  %s: %v", path, err)
		}

		if len(DataBatch) == batchSize {
			err := SendDataToIndex(&DataBatch)
			if err != nil {
				return fmt.Errorf("error while %s: %v", path, err)
			}
			DataBatch = nil // Clean bacth after sending the data
		}
	}

	return nil
}
