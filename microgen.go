package microgen

import (
	netUrl "net/url"

	"github.com/mejik-dev/microgen-v3-go/client"
)

type Client struct {
	apiKey   string
	queryUrl string
	headers  map[string]string

	Auth     *client.AuthClient
	Storage  *client.StorageClient
	Realtime *client.RealtimeClient
}

const defaultQueryUrl = "https://database-query.v3.microgen.id/api/v1"
const defaultStreamUrl = "https://database-stream.v3.microgen.id"

func DefaultURL() string {
	return defaultQueryUrl
}

func DefaultStreamURL() string {
	return defaultStreamUrl
}

func NewClient(apiKey string, url string, streamUrl string) *Client {
	if url == "" {
		url = defaultQueryUrl
	}

	if streamUrl == "" {
		streamUrl = defaultStreamUrl
	}

	urlWithApiKey, err := netUrl.JoinPath(url, apiKey)
	if err != nil {
		panic(err)
	}

	c := &Client{
		apiKey:   apiKey,
		headers:  map[string]string{},
		queryUrl: urlWithApiKey,
	}

	c.Auth = client.NewAuthClient(c.queryUrl+"/auth", c.headers)
	c.Storage = client.NewStorageClient(c.queryUrl+"/storage", c.Auth)
	c.Realtime = client.NewRealtimeClient(streamUrl, c.apiKey, c.Auth)

	return c
}

func (c *Client) getHeaders() map[string]string {
	headers := c.headers
	authBearer := c.Auth.Token()

	if authBearer != "" {
		headers["Authorization"] = "Bearer " + authBearer
	}

	return headers
}

func (c *Client) Service(name string) *client.QueryClient {
	return client.NewQueryClient(c.queryUrl+"/"+name, c.getHeaders())
}
