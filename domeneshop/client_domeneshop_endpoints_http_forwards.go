package domeneshop

import (
	"fmt"
	"github.com/zclconf/go-cty/cty/json"
)

type HttpForward struct {
	Host  string `json:"host"`
	Frame bool   `json:"frame"`
	Url   string `json:"url"`
}

func (s *Client) ListHttpForwards(domainId int) (*[]HttpForward, error) {
	url := fmt.Sprintf("domains/%d/forwards", domainId)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data []HttpForward
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) CreateHttpForward(domainId int, newForward HttpForward) (*HttpForward, error) {
	url := fmt.Sprintf("domains/%d/forwards", domainId)
	message, err := json.Marshal(newForward)
	if err != nil {
		return nil, err
	}
	bytes, err := s.Post(url, message)
	if err != nil {
		return nil, err
	}
	var data HttpForward
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) GetDnsRecordByHost(domainId int, host string) (*HttpForward, error) {
	url := fmt.Sprintf("domains/%d/dns/%s", domainId, host)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data HttpForward
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) UpdateHttpForward(domainId int, host string, updatedForward HttpForward) (*HttpForward, error) {
	url := fmt.Sprintf("domains/%d/forwards/%s", domainId, host)
	message, err := json.Marshal(updatedForward)
	if err != nil {
		return nil, err
	}
	bytes, err := s.Post(url, message)
	if err != nil {
		return nil, err
	}
	var data HttpForward
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) DeleteHttpForward(domainId int, host string) (*HttpForward, error) {
	url := fmt.Sprintf("domains/%d/forwards/%s", domainId, host)
	bytes, err := s.Delete(url, nil)
	if err != nil {
		return nil, err
	}
	var data HttpForward
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
