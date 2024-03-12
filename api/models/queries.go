package models

type EmailSearchQuery struct {
	SearchType   string     `json:"search_type"`
	Query        EmailQuery `json:"query"`
	From         int        `json:"from"`
	MaxResults   int        `json:"max_results"`
	SourceFields []string   `json:"_source"`
}

type EmailQuery struct {
	SortFields []string `json:"sort_fields"`
	Term       string   `json:"term"`
	StartTime  string   `json:"start_time"`
	EndTime    string   `json:"end_time"`
}

type AllEmailsQuery struct {
	SearchType   string   `json:"search_type"`
	SortFields   []string `json:"sort_fields"`
	From         int      `json:"from"`
	MaxResults   int      `json:"max_results"`
	SourceFields []string `json:"_source"`
}
