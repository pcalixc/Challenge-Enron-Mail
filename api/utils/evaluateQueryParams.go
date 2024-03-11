package utils

import (
	"challenge/api/models"
	"strconv"
)

// EvaluateQueryParams evaluates the provided query parameters such as max results,
// page number, and search term, and returns a QueryParameter structure.
// If max and page parameters are not empty, it converts them to integers and assigns them to MaxResults and PageNumber respectively.
// The search term is assigned directly.
func EvaluateQueryParams(max, page, term string) models.QueryParameters {
	// Create an instance of QueryParameter with default values.
	var params = models.QueryParameters{MaxResults: 8, PageNumber: 1, SearchTerm: term}

	if max != "" {
		params.MaxResults, _ = strconv.Atoi(max)
	}
	if page != "" {
		params.PageNumber, _ = strconv.Atoi(page)
	}
	params.SearchTerm = term

	// Return the updated QueryParameter structure.
	return params
}
