package models

type BulkDocument struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}
