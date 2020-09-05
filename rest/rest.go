package rest

import (
	"fmt"
	"net/http"
	"time"

	util "github.com/dsa0x/go-rocketchat/utils"
)

var httpClient = &http.Client{
	Timeout: time.Second * 15,
}

// Client the client struct
type Client struct {
	Protocol     string
	Host         string
	Port         string
	Username     string
	Password     string
	Credentials  *Credentials
	Url          string
	authRequired bool
}

func init() {

}

// New creates a new client
func New(protocol, host, port, username, password string) *Client {
	return &Client{Protocol: protocol, Host: host, Port: port}
}

// BuildURL builds the client url
func (client *Client) BuildURL(path string) string {
	client.Url = fmt.Sprintf("%s//%s:%s", client.Protocol, client.Host, client.Port)
	return fmt.Sprintf("%s//%s:%s", client.Protocol, client.Host+"/"+path, client.Port)
}

// Do Make http Request
func (client *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	r, err := httpClient.Do(req)
	if r != nil {
		defer r.Body.Close()
		util.CheckErr(err, "normal")
	}

	if client.authRequired {
		req.Header.Set("X-Auth-Token", client.Credentials.Data.Token)
		req.Header.Set("X-User-Id", client.Credentials.Data.UserID)
	}

	return r, err
}
