package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	util "github.com/dsa0x/go-rocketchat/utils"
)

type Email struct {
	address  string
	verified bool
}

type User struct {
	ID        string `json:"_id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Services  struct {
		Password struct {
			Bcrypt string `json:"bcrypt"`
		} `json:"password"`
	} `json:"services"`
	Username string   `json:"username"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Status   string   `json:"status"`
	Active   bool     `json:"active"`
	Roles    []string `json:"roles"`
	Emails   []Email  `json:"emails"`
}

type UserPayload struct {
	email                 string
	name                  string
	password              string
	username              string
	active                bool
	roles                 []string
	joinDefaultChannels   bool
	requirePasswordChange bool
	sendWelcomeEmail      bool
	verified              bool
	customFields          map[string]string
}

type UserStatus struct {
	message          string
	connectionStatus string
	status           string
}

// CreateUser only registered users can create a user
func (client *Client) CreateUser(user UserPayload) User {
	url := client.BuildURL("users.create")
	payload, err := json.Marshal(user)
	util.CheckErr(err, "normal")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	client.authRequired = true
	r, err := client.Do(req)
	util.CheckErr(err, "normal")
	userRes := User{}
	json.NewDecoder(r.Body).Decode(&userRes)
	return userRes

}
func (client *Client) RegisterUser(user UserPayload) User {
	url := client.BuildURL("users.register")
	// body := {}
	payload, err := json.Marshal(user)
	util.CheckErr(err, "normal")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	client.authRequired = true
	r, err := client.Do(req)
	util.CheckErr(err, "normal")
	userRes := User{}
	json.NewDecoder(r.Body).Decode(&userRes)
	return userRes

}

func (client *Client) DeleteOwnAccount(user User) {
	url := client.BuildURL("users.deleteOwnAccount")
	payload, err := json.Marshal("{password:" + user.Password + "}")
	util.CheckErr(err, "normal")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	client.authRequired = true
	r, err := client.Do(req)
	util.CheckErr(err, "normal")
	userRes := User{}
	json.NewDecoder(r.Body).Decode(&userRes)

}

func (client *Client) Delete(user User) {
	url := client.BuildURL("users.delete")
	payload, err := json.Marshal("{username:" + user.Username + "}")
	util.CheckErr(err, "normal")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	client.authRequired = true
	r, err := client.Do(req)
	util.CheckErr(err, "normal")
	userRes := User{}
	json.NewDecoder(r.Body).Decode(&userRes)

}
func (client *Client) ForgotPassword(user User) {
	url := client.BuildURL("users.forgotPassword")
	payload, err := json.Marshal("{email:" + user.Emails[0].address + "}")
	util.CheckErr(err, "normal")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	client.authRequired = true
	r, err := client.Do(req)
	util.CheckErr(err, "normal")
}
func (client *Client) GetStatus(user User) {
	url := client.BuildURL("users.getStatus")
	payload, err := json.Marshal("{username:" + user.Username + "}")
	util.CheckErr(err, "normal")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	client.authRequired = true
	r, err := client.Do(req)
	util.CheckErr(err, "normal")
	userStatus := UserStatus{}
	json.NewDecoder(r.Body).Decode(&userStatus)
}
