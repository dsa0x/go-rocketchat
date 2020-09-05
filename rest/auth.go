package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	util "github.com/dsa0x/go-rocketchat/utils"
)

type Credentials struct {
	Status string `json:"status"`
	Data   struct {
		UserID string `json:"userId"`
		Token  string `json:"authToken"`
	} `json:"data"`
}

type LoginPayload struct {
	username string `json:"username"`
	password string `json:"password"`
}

// type User struct {
// 	ID       string   `json:"_id"`
// 	Name     string   `json:"name"`
// 	Username string   `json:"username"`
// 	Roles    []string `json:"roles"`
// 	Active   bool     `json:"active"`
// }

func (client *Client) login() {
	url := client.BuildURL("login")
	payload, err := json.Marshal(LoginPayload{username: client.Username, password: client.Password})
	util.CheckErr(err, "normal")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	util.CheckErr(err, "normal")
	req.Header.Set("Content-Type", "application/json")
	client.authRequired = false
	r, err := client.Do(req)
	util.CheckErr(err, "normal")
	json.NewDecoder(r.Body).Decode(client.Credentials)

}

func (client *Client) logout() {
	url := client.BuildURL("logout")
	req, err := http.NewRequest("POST", url, nil)
	client.authRequired = true
	r, err := client.Do(req)
	util.CheckErr(err, "normal")
	json.NewDecoder(r.Body).Decode(client.Credentials)
}

// Authenticated user
func (client *Client) me() User {
	url := client.BuildURL("me")
	req, err := http.NewRequest("GET", url, nil)
	client.authRequired = true
	r, err := client.Do(req)
	util.CheckErr(err, "normal")
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	return user
}
