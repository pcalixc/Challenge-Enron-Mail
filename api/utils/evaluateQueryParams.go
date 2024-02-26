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
	var params = models.QueryParameters{MaxResults: 10, PageNumber: 1, Term: term}

	// If the max parameter is not empty, it is converted to an integer and assigned to MaxResults.
	if max != "" {
		params.MaxResults, _ = strconv.Atoi(max)
	}
	// If the page parameter is not empty, it is converted to an integer and assigned to PageNumber.
	if page != "" {
		params.PageNumber, _ = strconv.Atoi(page)
	}
	// Assign the provided search term.
	params.Term = term

	// Return the updated QueryParameter structure.
	return params
}
