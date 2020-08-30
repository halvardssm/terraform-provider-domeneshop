package domeneshop

import (
	"fmt"
	"github.com/zclconf/go-cty/cty/json"
)

type DomainServices struct {
	Registrar bool   `json:"registrar"`
	Dns       bool   `json:"dns"`
	Email     bool   `json:"email"`
	Webhotel  string `json:"webhotel"`
}

type Domain struct {
	Id             int            `json:"id"`
	Domain         string         `json:"domain"`
	ExpiryDate     string         `json:"expiry_date"`
	RegisteredDate string         `json:"registered_date"`
	Renew          bool           `json:"renew"`
	Registrant     string         `json:"registrant"`
	Status         string         `json:"status"`
	Nameservers    []string       `json:"nameservers"`
	Services       DomainServices `json:"services"`
}

func (s *Client) ListDomains(tld string) (*[]Domain, error) {
	bytes, err := s.Get("domains", nil)
	if err != nil {
		return nil, err
	}
	var data []Domain
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) GetDomainById(domainId int) (*Domain, error) {
	url := fmt.Sprintf("domains/%d", domainId)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data Domain
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
