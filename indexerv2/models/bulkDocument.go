package models

type BulkDocument struct {
	Index   string      `json:"index"`
	Records []EnronMail `json:"records"`
}
