package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	util "github.com/dsa0x/go-rocketchat/utils"
)

func (client *Client) AddUserToRole(role string, user User, roomId ...string) {
	url := client.BuildURL("roles.addUserToRole")
	var body string
	if roomId[0] != "" {
		body = fmt.Sprintf("{roleName: %s,username:%s,roomId:%s}", role, user.Username, roomId[0])
	} else {
		body = fmt.Sprintf("{roleName: %s,username:%s}", role, user.Username)
	}
	payload, err := json.Marshal(body)
	util.CheckErr(err, "normal")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	client.authRequired = true
	_, err = client.Do(req)
	util.CheckErr(err, "normal")
}
func (client *Client) CreateRole(role string, description string, scope ...string) {
	url := client.BuildURL("roles.create")
	var body string
	if scope[0] != "" {
		body = fmt.Sprintf("{name: %s,scope:%s,description:%s}", role, scope[0], description)
	} else {
		body = fmt.Sprintf("{name: %s,description:%s}", role, description)
	}
	payload, err := json.Marshal(body)
	util.CheckErr(err, "normal")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	client.authRequired = true
	_, err = client.Do(req)
	util.CheckErr(err, "normal")
}
