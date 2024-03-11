package models

type SearchQuery struct {
	SearchType   string   `json:"search_type"`
	Query        Query    `json:"query"`
	From         int      `json:"from"`
	MaxResults   int      `json:"max_results"`
	SourceFields []string `json:"_source"`
}

type Query struct {
	SortFields []string `json:"sort_fields"`
	Term       string   `json:"term"`
	StartTime  string   `json:"start_time"`
	EndTime    string   `json:"end_time"`
}
