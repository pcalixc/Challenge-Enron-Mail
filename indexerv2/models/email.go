package models

import "time"

type Email struct {
	MessageID               string    `json:"message_id"`
	Date                    time.Time `json:"date"`
	From                    string    `json:"from"`
	To                      string    `json:"to"`
	Subject                 string    `json:"subject"`
	Cc                      string    `json:"cc"`
	MimeVersion             string    `json:"mime_version"`
	ContentType             string    `json:"content_type"`
	ContentTransferEncoding string    `json:"content_transfer_encoding"`
	Bcc                     string    `json:"bcc"`
	XFrom                   string    `json:"x_from"`
	XTo                     string    `json:"x_to"`
	XCc                     string    `json:"x_cc"`
	XBcc                    string    `json:"x_bcc"`
	XFolder                 string    `json:"x_folder"`
	XOrigin                 string    `json:"x_origin"`
	XFileName               string    `json:"x_file_name"`
	Content                 string    `json:"content"`
}
