package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Config struct {
	BaseURL  string
	Index    string
	Username string
	Password string
}

// We create a struct that contains the structure of the JSON we will send to zincsearch
type Email struct {
	MessageID               string `json:"Message_id"`
	Date                    string `json:"Date"`
	From                    string `json:"From"`
	To                      string `json:"To"`
	Subject                 string `json:"Subject"`
	MimeVersion             string `json:"Mime_version"`
	ContentType             string `json:"Content_type"`
	ContentTransferEncoding string `json:"Content_transfer_encoding"`
	X_from                  string `json:"X-from"`
	X_to                    string `json:"X-to"`
	X_CC                    string `json:"X-cc"`
	X_BCC                   string `json:"X-bcc"`
	X_folder                string `json:"X-folder"`
	X_origin                string `json:"X-origin"`
	X_fileName              string `json:"X-file_name"`
	Body                    string `json:"body"`
}

type BulkDocument struct {
	Index   string  `json:index`
	Records []Email `json:records`
}

// the structure of the API response from the searchDocument() function
type Hit struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source struct {
		MessageID               string `json:"Message_id"`
		Date                    string `json:"Date"`
		From                    string `json:"From"`
		To                      string `json:"To"`
		Subject                 string `json:"Subject"`
		MimeVersion             string `json:"Mime_version"`
		ContentType             string `json:"Content_type"`
		ContentTransferEncoding string `json:"Content_transfer_encoding"`
		X_from                  string `json:"X-from"`
		X_to                    string `json:"X-to"`
		X_CC                    string `json:"X-cc"`
		X_BCC                   string `json:"X-bcc"`
		X_folder                string `json:"X-folder"`
		X_origin                string `json:"X-origin"`
		X_fileName              string `json:"X-file_name"`
		Body                    string `json:"body"`
	} `json:"_source"`
}

type HitsResponse struct {
	Hits struct {
		Hits []Hit `json:"hits"`
	} `json:"hits"`
}

// using the ZincSearch API to create a document and add it to the index
func CreateDocument(bodyQuery []byte, config Config) {
	requestURL := fmt.Sprintf("%s/api/%s/_doc", config.BaseURL, config.Index)
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(bodyQuery))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	log.Println(res.StatusCode)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func BulkCreateDocument(emails []Email, config Config) {
	requestURL := fmt.Sprintf("%s/api/_bulkv2", config.BaseURL)
	bulkData := BulkDocument{
		Index:   config.Index,
		Records: emails,
	}

	jsonBody, err := json.Marshal(bulkData)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	log.Println(res.StatusCode)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func SearchDocument(word string) HitsResponse {
	var query = fmt.Sprintf(`{
		"aggs": {
		  "histogram": "select histogram(_timestamp, '30 second') AS zo_sql_key, count(*) AS zo_sql_num from query GROUP BY zo_sql_key ORDER BY zo_sql_key"
		},
		"query": {
		  "end_time": 1675185660872049,
		  "from": 0,
		  "size": 10,
		  "sql": "select * from k8s ",
		  "start_time": 1675182660872049
		}
	  }`, word)
	req, err := http.NewRequest("POST", "http://localhost:5080/api/default/_search", strings.NewReader(query))
	if err != nil {
		log.Panic(err)
	}
	req.SetBasicAuth("root@example.com", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)

	var hitsResponse HitsResponse
	err = json.NewDecoder(resp.Body).Decode(&hitsResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", hitsResponse)

	return hitsResponse
}
