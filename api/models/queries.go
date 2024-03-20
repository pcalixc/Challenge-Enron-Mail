package models

type EmailSearchQuery struct {
	SearchType   string     `json:"search_type"`
	Query        EmailQuery `json:"query"`
	SortFields   []string   `json:"sort_fields"`
	From         int        `json:"from"`
	MaxResults   int        `json:"max_results"`
	SourceFields []string   `json:"_source"`
}

type EmailQuery struct {
	Term      string `json:"term"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type AllEmailsQuery struct {
	SearchType   string   `json:"search_type"`
	SortFields   []string `json:"sort_fields"`
	From         int      `json:"from"`
	MaxResults   int      `json:"max_results"`
	SourceFields []string `json:"_source"`
}

type SearchZincRequest struct {
	SearchType string `json:"search_type"`
	Query      struct {
		Term  string `json:"term"`
		Field string `json:"field"`
	}
	SortFields []string `json:"sort_fields"`
	From       uint     `json:"from"`
	MaxResults uint     `json:"max_results"`
	Source     []string `json:"_source"`
	Highlight  struct {
		Fields map[string]interface{} `json:"fields"`
	}
}
