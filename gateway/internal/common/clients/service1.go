package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Service1Client struct {
	BaseURL string
}

func NewService1Client(baseURL string) *Service1Client {
	return &Service1Client{BaseURL: baseURL}
}

func (c *Service1Client) GetUser(id string) (*User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%s", c.BaseURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
