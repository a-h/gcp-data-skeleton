package models

type SamplePostRequest struct {
	Sample
}

type SamplePostResponse struct {
	OK       bool   `json:"ok"`
	ServerID string `json:"serverId"`
}

type Sample struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
