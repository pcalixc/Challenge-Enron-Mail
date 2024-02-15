package controllers

import (
	"challenge/api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func SearchMails(term, page, max string) (models.HitsResponse, error) {
	intMax, _ := strconv.Atoi(max)
	intPage, _ := strconv.Atoi(page)

	from := (intPage-1)*intMax + 1

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
    }`, term, from, intMax)
	req, err := http.NewRequest("POST", os.Getenv("ZS_BASE_URL")+"/api/mail/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(os.Getenv("ZS_USER"), os.Getenv("ZS_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("45", err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)

	var hitsResponse models.HitsResponse
	if resp.ContentLength == 0 {
		log.Println("esta vaciooo")
		return hitsResponse, nil
	}
	err = json.NewDecoder(resp.Body).Decode(&hitsResponse)
	if err != nil {
		log.Fatal(resp, err)
	}
	return hitsResponse, nil
}

func GetAllEmails(page, max string) models.HitsResponse {
	intMax, _ := strconv.Atoi(max)
	intPage, _ := strconv.Atoi(page)

	from := (intPage-1)*intMax + 1

	var query = fmt.Sprintf(`{
		"search_type": "matchall",
		"sort_fields": ["-Date"],
		"from": %d,
		"max_results": %d,
		"_source": []
	}`, from, intMax)
	req, err := http.NewRequest("POST", os.Getenv("ZS_BASE_URL")+"/api/mail/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(os.Getenv("ZS_USER"), os.Getenv("ZS_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)

	var hitsResponse models.HitsResponse
	err = json.NewDecoder(resp.Body).Decode(&hitsResponse)
	if err != nil {
		log.Fatal(err)

	}
	return hitsResponse
}
