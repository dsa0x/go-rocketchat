package rest

import (
	"fmt"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 15,
}

// Client the client struct
type Client struct {
	Protocol string
	Host     string
	Port     string
	Username string
	Password string
}

func init() {

}

// BuildURL builds the client url
func (client *Client) BuildURL() string {
	return fmt.Sprintf("%s//%s:%s", client.Protocol, client.Host, client.Port)
}
