package config

import "net/http"

type APIType string
type Config struct {
	ListenOn     string `yaml:"listen_on"`
	OpenaiClient Client `yaml:"client"`
}
type Client struct {
	Token        string  `yaml:"token"`
	Model        string  `yaml:"model"`
	BaseURL      string  `yaml:"base_url"`
	Organization string  `yaml:"origanization"`
	ApiType      APIType `yaml:"api_type"`
	HttpClient   Doer    `yaml:"-"`

	// required when APIType is APITypeAzure or APITypeAzureAD
	ApiVersion      string `yaml:"api_ver"`
	EmbeddingsModel string `yaml:"embeddings_model"`
}
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}
