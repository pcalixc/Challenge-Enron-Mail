package models

type HitsResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		}
		Hits []Hit `json:"hits"`
	} `json:"hits"`
}
