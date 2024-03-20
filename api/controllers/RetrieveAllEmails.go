package controllers

import (
	"challenge/api/models"
	"encoding/json"
	"fmt"
)

func RetrieveAllEmails(page, max int) (models.HitsResponse, error) {

	from := (page-1)*max + 1

	query := models.AllEmailsQuery{
		SearchType:   "matchall",
		SortFields:   []string{"-date"},
		From:         from,
		MaxResults:   max,
		SourceFields: []string{"subject", "from", "to", "date", "content"},
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return models.HitsResponse{}, fmt.Errorf("failed to marshal query to JSON: %w", err)
	}

	resp, err := executeRequest(string(queryJSON))
	if err != nil {
		return models.HitsResponse{}, fmt.Errorf("failed to get all emails: %w", err)
	}
	defer resp.Body.Close()

	var hitsResponse models.HitsResponse

	err = json.NewDecoder(resp.Body).Decode(&hitsResponse)

	if err != nil {
		return models.HitsResponse{}, fmt.Errorf("failed to decode response body: %w", err)
	}

	return hitsResponse, nil
}
