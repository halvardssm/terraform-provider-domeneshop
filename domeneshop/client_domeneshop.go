package domeneshop

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseUrl string = "https://api.domeneshop.no/v0/"

type Client struct {
	Token  string
	Secret string
}

func NewBasicAuthClient(token, secret string) *Client {
	return &Client{
		Token:  token,
		Secret: secret,
	}
}

func (s *Client) Request(method, endpoint string, message []byte) ([]byte, error) {
	reqUrl := baseUrl + endpoint

	req, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(message))

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(s.Token, s.Secret)
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || 300 <= res.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

func (s *Client) Get(endpoint string, params map[string]string) ([]byte, error) {
	u, _ := url.Parse(endpoint)
	q, _ := url.ParseQuery(u.RawQuery)

	for key, value := range q {
		q.Add(key, value)
	}

	u.RawQuery = q.Encode()

	return s.Request("GET", u.String(), nil)
}

func (s *Client) Post(endpoint string, message []byte) ([]byte, error) {
	return s.Request("POST", endpoint, message)
}

func (s *Client) Put(endpoint string, message []byte) ([]byte, error) {
	return s.Request("PUT", endpoint, message)
}

func (s *Client) Delete(endpoint string, message []byte) ([]byte, error) {
	return s.Request("DELETE", endpoint, message)
}
