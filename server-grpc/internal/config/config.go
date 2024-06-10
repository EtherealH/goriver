package config

import "net/http"

type APIType string
type Config struct {
	ListenOn     string
	ModelPath    string
	OpenaiClient Client
}
type Client struct {
	Token        string
	Model        string
	BaseURL      string
	Organization string
	ApiType      APIType
	HttpClient   Doer

	// required when APIType is APITypeAzure or APITypeAzureAD
	ApiVersion      string
	EmbeddingsModel string
}
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}
