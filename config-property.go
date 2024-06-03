package dtrack

import (
	"context"
	"net/http"
)

type ConfigProperty struct {
	GroupName     string `json:"groupName"`
	PropertyName  string `json:"propertyName"`
	PropertyValue string `json:"propertyValue,omitempty"`
	PropertyType  string `json:"propertyType"`
	Description   string `json:"description"`
}

type ConfigPropertyService struct {
	client *Client
}

func (s ConfigPropertyService) GetAllConfigProperty(ctx context.Context, po PageOptions) (p Page[ConfigProperty], err error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/api/v1/configProperty", withPageOptions(po))
	if err != nil {
		return
	}

	res, err := s.client.doRequest(req, &p.Items)
	if err != nil {
		return
	}

	p.TotalCount = res.TotalCount
	return
}

func (s ConfigPropertyService) UpdateConfigProperty(ctx context.Context, cp ConfigProperty) (out ConfigProperty, err error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/api/v1/configProperty", withBody(cp))
	if err != nil {
		return
	}

	_, err = s.client.doRequest(req, &out)
	return
}
