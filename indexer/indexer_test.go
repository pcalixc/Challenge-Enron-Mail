package main

import (
	"encoding/json"
	"fmt"
	"indexer/models"
	"indexer/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertEmailFileToJSONEmptyFile(t *testing.T) {
	emptyFile := "empty_file.txt"
	_ = os.WriteFile(emptyFile, []byte(""), 0644)

	defer func() {
		_ = os.Remove(emptyFile)
	}()

	// call function
	result := utils.ConvertEmailFileToJSON(emptyFile)

	// verify if the result is empty
	expectedResult := models.EnronMail{}
	assert.Equal(t, expectedResult, result)
}

func TestListFiles(t *testing.T) {
	// Create a new directory and some files
	testDir := "test_dir"
	testFiles := []string{"file1.txt", "file2.txt", "file3.txt"}

	err := os.Mkdir(testDir, 0755)
	assert.NoError(t, err)
	defer os.RemoveAll(testDir)

	for _, file := range testFiles {
		filePath := testDir + "/" + file
		err := os.WriteFile(filePath, []byte("content"), 0644)
		assert.NoError(t, err)
		defer os.Remove(filePath)
	}

	result, err := utils.ListFiles(testDir)

	// Verify
	assert.NoError(t, err)
	assert.ElementsMatch(t, testFiles, result)
}

func TestListFilesNonexistentDirectory(t *testing.T) {
	result, err := utils.ListFiles("nonexistent_directory")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestMapEmailHeaders(t *testing.T) {
	//Create an Email structure
	emailStruct := models.EnronMail{}

	emailStruct = utils.MapEmailHeaders("Subject", "Test Subject", emailStruct)
	emailStruct = utils.MapEmailHeaders("From", "sender@example.com", emailStruct)
	emailStruct = utils.MapEmailHeaders("Date", "2022-01-01", emailStruct)

	// Verify
	assert.Equal(t, "Test Subject", emailStruct.Subject)
	assert.Equal(t, "sender@example.com", emailStruct.From)
	assert.Equal(t, "2022-01-01", emailStruct.Date)

	// Verify if other fields were updated
	assert.Empty(t, emailStruct.MessageID)
	assert.Empty(t, emailStruct.To)
}

func TestConvertEmailFileToJSON(t *testing.T) {
	tempFile, err := os.CreateTemp("", "email_test_*.txt")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	emailContent := []byte(
		`From: alan.aronowitz@enron.com
	To: matthias.lee@enron.com
	Subject: Re: counterparty request to migrate trades

	I have no objections, but I am copying Elizabeth Sager on this note to see if she has any issues. She is handling the Credit legal issues.`)
	err = os.WriteFile(tempFile.Name(), emailContent, 0644)
	assert.NoError(t, err)

	email := utils.ConvertEmailFileToJSON(tempFile.Name())

	expectedEmail := models.EnronMail{
		From:    "alan.aronowitz@enron.com",
		To:      "matthias.lee@enron.com",
		Subject: "Re: counterparty request to migrate trades",
		Content: "I have no objections, but I am copying Elizabeth Sager on this note to see if she has any issues. She is handling the Credit legal issues.",
	}
	assert.Equal(t, expectedEmail, email)
}

func TestSendDataToIndex(t *testing.T) {
	// Mock data
	mockData := []models.EnronMail{
		{From: "sender1@example.com", To: "receiver1@example.com", Subject: "Test Email 1", Content: "Content 1"},
		{From: "sender2@example.com", To: "receiver2@example.com", Subject: "Test Email 2", Content: "Content 2"},
	}

	// Mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, os.Getenv("ZS_BASE_URL")+"/api/_bulkv2", r.URL.String())

		// Read request body
		var bulkData models.BulkDocument
		err := json.NewDecoder(r.Body).Decode(&bulkData)
		assert.NoError(t, err)

		// Verify request body
		assert.Equal(t, "mail", bulkData.Index)
		assert.Equal(t, mockData, bulkData.Records)

		// Respond with success
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Bulk data successfully inserted")
	}))
	defer mockServer.Close()

	// Override ZS_BASE_URL for the test
	oldBaseUrl := os.Getenv("ZS_BASE_URL")
	os.Setenv("ZS_BASE_URL", mockServer.URL)
	defer os.Setenv("ZS_BASE_URL", oldBaseUrl)

	// Call the function to test
	err := utils.SendDataToIndex(&mockData)

	// Assertions
	assert.NoError(t, err)
}

// func SendDataToIndex(data *[]models.EnronMail)

// func FillIndexBatch(path string)

// func IndexEmailDirectory(path string)
