package fb_client

import (
	"net/http"
)

// ...
type ClientInfo struct {
	FBUCUsername string
	GrowthLevel  int
	RespondTo    string
	Uid          string
}

// ...
type Client struct {
	ClientInfo ClientInfo
	HttpClient http.Client
	AuthServer string
}

// ...
func CreateClient(authServer string) *Client {

	authclient := &Client{
		HttpClient: http.Client{Transport: &SecretLoadingTransport{
			secret: "Fluid",
		}},
		AuthServer: authServer,
	}

	return authclient
}
