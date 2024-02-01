package models

// type EnronMail struct {
// 	MessageID               string `json:"Message_id"`
// 	Date                    string `json:"Date"`
// 	From                    string `json:"From"`
// 	To                      string `json:"To"`
// 	Subject                 string `json:"Subject"`
// 	MimeVersion             string `json:"Mime_version"`
// 	ContentType             string `json:"Content_type"`
// 	ContentTransferEncoding string `json:"Content_transfer_encoding"`
// 	X_from                  string `json:"X-from"`
// 	X_to                    string `json:"X-to"`
// 	X_CC                    string `json:"X-cc"`
// 	X_BCC                   string `json:"X-bcc"`
// 	X_folder                string `json:"X-folder"`
// 	X_origin                string `json:"X-origin"`
// 	X_fileName              string `json:"X-file_name"`
// 	Body                    string `json:"body"`
// }

type EnronMail struct {
	MessageID               string `json:"message_id"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	To                      string `json:"to"`
	Subject                 string `json:"subject"`
	Cc                      string `json:"cc"`
	MimeVersion             string `json:"mime_version"`
	ContentType             string `json:"content_type"`
	ContentTransferEncoding string `json:"content_transfer_encoding"`
	Bcc                     string `json:"bcc"`
	XFrom                   string `json:"x_from"`
	XTo                     string `json:"x_to"`
	XCc                     string `json:"x_cc"`
	XBcc                    string `json:"x_bcc"`
	XFolder                 string `json:"x_folder"`
	XOrigin                 string `json:"x_origin"`
	XFileName               string `json:"x_file_name"`
	Content                 string `json:"content"`
}
