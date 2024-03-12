package controllers

import (
	"challenge/api/models"
	"encoding/json"
	"fmt"
)

func SearchEmails(term string, page, max int) (models.HitsResponse, error) {
	from := (page-1)*max + 1

	query := models.EmailSearchQuery{
		SearchType: "matchphrase",
		Query: models.EmailQuery{
			SortFields: []string{"@date"},
			Term:       term,
			StartTime:  "2021-06-02T14:28:31.894Z",
			EndTime:    "2028-01-31T15:28:31.894Z",
		},
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
		return models.HitsResponse{}, fmt.Errorf("failed to search emails: %w", err)
	}
	defer resp.Body.Close()

	var hitsResponse models.HitsResponse
	if resp.ContentLength == 0 {
		return hitsResponse, nil
	}

	err = json.NewDecoder(resp.Body).Decode(&hitsResponse)
	if err != nil {
		return models.HitsResponse{}, fmt.Errorf("failed to decode response body: %w", err)
	}
	return hitsResponse, nil
}
