package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type DNSRecord struct {
	Id            string   `json:"id"`
	Domain        string   `json:"domain"`
	IPV4Addresses []string `json:"ipv4Addresses"`
	IPV6Addresses []string `json:"ipv6Addresses"`
}

func (c *Client) CreateDNSRecord(record DNSRecord) (*DNSRecord, error) {
	recordJson, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/beta/dns-records", c.BaseURL), bytes.NewBuffer(recordJson))
	if err != nil {
		return nil, err
	}
	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}
	var d DNSRecord
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (c *Client) GetDNSRecord(recordId string) (*DNSRecord, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/beta/dns-records", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}
	var records []DNSRecord
	err = json.Unmarshal(body, &records)
	if err != nil {
		return nil, err
	}
	for _, r := range records {
		if r.Id == recordId {
			return &r, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Route with id %s was not found", recordId))
}

func (c *Client) UpdateDNSRecord(record DNSRecord) error {
	recordJson, err := json.Marshal(record)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/beta/dns-records/%s", c.BaseURL, record.Id), bytes.NewBuffer(recordJson))
	if err != nil {
		return err
	}
	_, err = c.DoRequest(req)
	return err
}

func (c *Client) DeleteDNSRecord(recordId string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/beta/dns-records/%s", c.BaseURL, recordId), nil)
	if err != nil {
		return err
	}
	_, err = c.DoRequest(req)
	return err
}
