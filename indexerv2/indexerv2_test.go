package main

import (
	"fmt"
	"indexer/models"
	"indexerv2/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestFilePath = "./1."

func TestConvertEmailFileToStruct(t *testing.T) {
	email, err := utils.ConvertEmailFileToStruct(TestFilePath)

	if err != nil {
		t.Fatalf("ConvertEmailFileToStruct failed with error: %v", err)
	}
	assert.Equal(t, "Here is our forecast", email.Subject)
	assert.Equal(t, "phillip.allen@enron.com", email.From)
	assert.Equal(t, "tim.belden@enron.com", email.To)
	assert.Equal(t, "Here is our forecast\n", email.Content)
}
func TestMapEmailHeaders(t *testing.T) {
	email := &models.EnronMail{}

	utils.MapEmailHeaders("Subject", "Test Subject", email)
	assert.Equal(t, "Test Subject", email.Subject)

	utils.MapEmailHeaders("From", "test@example.com", email)
	assert.Equal(t, "test@example.com", email.From)

	utils.MapEmailHeaders("To", "recipient@example.com", email)
	assert.Equal(t, "recipient@example.com", email.To)

	utils.MapEmailHeaders("Date", "Mon, 14 May 2001 16:39:00 -0700 (PDT)", email)
	assert.Equal(t, "Mon, 14 May 2001 16:39:00 -0700 (PDT)", email.Date)

	utils.MapEmailHeaders("Unknown", "Unknown Value", email)
}

func TestSendDataToIndex_Success(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Bulk data successfully inserted")
	}))
	defer server.Close()

	// Set the mock server URL as the ZS_BASE_URL environment variable
	os.Setenv("ZS_BASE_URL", server.URL)

	// Test case: successful indexing
	data := []models.EnronMail{ /* mock data */ }
	err := utils.SendDataToIndex(&data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSendDataToIndex_ServerError(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate error response
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}))
	defer server.Close()

	// Set the mock server URL
	os.Setenv("ZS_BASE_URL", server.URL)

	// Test case: server error
	data := []models.EnronMail{ /* mock data */ }
	err := utils.SendDataToIndex(&data)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
