package controllers

import (
	"challenge/api/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var ErrHTTPRequestFailed = errors.New("HTTP request failed")

func doRequest(query string) (*http.Response, error) {
	req, err := http.NewRequest("POST", os.Getenv("ZS_BASE_URL")+"/api/emails/_search", strings.NewReader(query))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.SetBasicAuth(os.Getenv("ZS_USER"), os.Getenv("ZS_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	return resp, nil
}

func SearchMails(term string, page, max int) (models.HitsResponse, error) {
	from := (page-1)*max + 1

	//ini= (pag-1) * num + 1
	//fin= pag*num

	var query = fmt.Sprintf(`{
        "search_type": "match",
        "query":
        {
            "term": "%s",
            "start_time": "2021-06-02T14:28:31.894Z",
            "end_time": "2028-01-31T15:28:31.894Z"
        },
        "from": %d,
        "max_results": %d,
        "_source": []
    }`, term, from, max)

	resp, err := doRequest(query)
	if err != nil {
		return models.HitsResponse{}, fmt.Errorf("failed to search mails: %w", err)
	}
	defer resp.Body.Close()

	var hitsResponse models.HitsResponse
	if resp.ContentLength == 0 {
		return hitsResponse, nil
	}

	err = json.NewDecoder(resp.Body).Decode(&hitsResponse)
	if err != nil {
		return models.HitsResponse{}, fmt.Errorf("failed to decode response body: %w", err)
	}
	return hitsResponse, nil
}

func GetAllEmails(page, max int) (models.HitsResponse, error) {

	from := (page-1)*max + 1

	var query = fmt.Sprintf(`{
		"search_type": "matchall",
		"sort_fields": ["-Date"],
		"from": %d,
		"max_results": %d,
		"_source": []
	}`, from, max)

	resp, err := doRequest(query)
	if err != nil {
		return models.HitsResponse{}, fmt.Errorf("failed to get all emails: %w", err)
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	var hitsResponse models.HitsResponse
	err = json.NewDecoder(resp.Body).Decode(&hitsResponse)
	if err != nil {
		return models.HitsResponse{}, fmt.Errorf("failed to decode response body: %w", err)
	}
	return hitsResponse, nil
}
