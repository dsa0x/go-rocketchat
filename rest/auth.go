package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Credentials struct {
	User   string
	Token  string
	Status string
}

type LoginPayload struct {
	username string `json:"username"`
	password string `json:"password"`
}

func login(client *Client) {
	url := client.BuildURL()
	payload, err := json.Marshal(LoginPayload{username: client.Username, password: client.Password})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer())
	req.Header.Set("Content-Type", "application/json")

	r, err := httpClient.Do(req)

}
