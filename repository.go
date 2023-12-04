package dtrack

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	RepositoryTypeCargo       = "CARGO"
	RepositoryTypeComposer    = "COMPOSER"
	RepositoryTypeCpan        = "CPAN"
	RepositoryTypeGem         = "GEM"
	RepositoryTypeGoModules   = "GO_MODULES"
	RepositoryTypeHex         = "HEX"
	RepositoryTypeMaven       = "MAVEN"
	RepositoryTypeNpm         = "NPM"
	RepositoryTypeNuget       = "NUGET"
	RepositoryTypePypi        = "PYPI"
	RepositoryTypeUnsupported = "UNSUPPORTED"
)

type RepositoryType string

type Repository struct {
	Type            RepositoryType `json:"type"`
	Identifier      string         `json:"identifier"`
	Url             string         `json:"url"`
	ResolutionOrder int            `json:"resolutionOrder"`
	Enabled         bool           `json:"enabled"`
	Internal        bool           `json:"internal"`
	Username        string         `json:"username,omitempty"`
	Password        string         `json:"password,omitempty"`
	UUID            uuid.UUID      `json:"uuid,omitempty"`
}

type RepositoryMetaComponent struct {
	LatestVersion string `json:"latestVersion"`
}

type RepositoryService struct {
	client *Client
}

func (rs RepositoryService) GetMetaComponent(ctx context.Context, purl string) (r RepositoryMetaComponent, err error) {
	params := map[string]string{
		"purl": purl,
	}

	req, err := rs.client.newRequest(ctx, http.MethodGet, "/api/v1/repository/latest", withParams(params))
	if err != nil {
		return
	}

	_, err = rs.client.doRequest(req, &r)
	return
}
