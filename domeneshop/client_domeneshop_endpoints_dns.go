package domeneshop

import (
	"fmt"
	"github.com/zclconf/go-cty/cty/json"
)

type Dns struct {
	Host     string `json:"host"`
	Ttl      int16  `json:"ttl"`
	Type     string `json:"type"`
	Data     string `json:"data"`
	Priority int16  `json:"priority"`
	Weight   int16  `json:"weight"`
	Port     int16  `json:"port"`
}

type DnsCreateResponse struct {
	Id int `json:"id"`
}

func (s *Client) ListDnsRecords(domainId int, host, recordType string) (*[]Dns, error) {
	url := fmt.Sprintf("domains/%d/dns", domainId)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data []Dns
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) CreateDnsRecord(domainId int, newRecord Dns) (*DnsCreateResponse, error) {
	url := fmt.Sprintf("domains/%d/dns", domainId)
	message, err := json.Marshal(newRecord)
	if err != nil {
		return nil, err
	}
	bytes, err := s.Post(url, message)
	if err != nil {
		return nil, err
	}
	var data DnsCreateResponse
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) GetDnsRecordById(domainId, recordId int) (*Dns, error) {
	url := fmt.Sprintf("domains/%d/dns/%d", domainId, recordId)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data Dns
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) UpdateDnsRecord(domainId, recordId int, updatedRecord Dns) (*DnsCreateResponse, error) {
	url := fmt.Sprintf("domains/%d/dns/%d", domainId, recordId)
	message, err := json.Marshal(updatedRecord)
	if err != nil {
		return nil, err
	}
	bytes, err := s.Put(url, message)
	if err != nil {
		return nil, err
	}
	var data DnsCreateResponse
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) DeleteDnsRecord(domainId, recordId int) (*DnsCreateResponse, error) {
	url := fmt.Sprintf("domains/%d/dns/%d", domainId, recordId)
	bytes, err := s.Delete(url, nil)
	if err != nil {
		return nil, err
	}
	var data DnsCreateResponse
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
