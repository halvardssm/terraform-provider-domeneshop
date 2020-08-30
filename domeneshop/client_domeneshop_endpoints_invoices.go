package domeneshop

import (
	"fmt"
	"github.com/zclconf/go-cty/cty/json"
)

type Invoice struct {
	Id         int    `json:"id"`
	Type       string `json:"type"`
	Amount     int    `json:"amount"`
	Currency   string `json:"currency"`
	DueDate    string `json:"due_date"`
	IssuedDate string `json:"issued_date"`
	PaidDate   string `json:"paid_date"`
	Status     string `json:"status"`
	Url        string `json:"url"`
}

func (s *Client) ListInvoices(status string) (*[]Invoice, error) {
	url := "invoice"
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data []Invoice
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) GetInvoiceByInvoiceNumber(invoiceId int) (*Invoice, error) {
	url := fmt.Sprintf("invoice/%d", invoiceId)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data Invoice
	_, err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
