package dotaopenapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mrkgnt/dota2-gridify/internal/heroes"
	"github.com/pkg/errors"
)

var endpoint = "https://api.opendota.com/api/heroes"

type dotaOpenApiRepository struct {
	client *http.Client
	data   *heroes.Heroes
}

func newDotaOpenApiClient() *http.Client {
	return &http.Client{Timeout: 10 * time.Second}
}

func (doa *dotaOpenApiRepository) NewDotaOpenApiRepository() (heroes.HeroesRepository, error) {
	repo := &dotaOpenApiRepository{}
	repo.client = newDotaOpenApiClient()
	return repo, nil
}

func (doa *dotaOpenApiRepository) GetHeroes() ([]byte, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "heroes.repository.dotaOpenApi.GetHeroes")
	}

	resp, err := doa.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "heroes.repository.dotaOpenApi.GetHeroes")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "heroes.repository.dotaOpenApi.GetHeroes")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(errors.New(fmt.Sprintf("Unexpected status code: %d", resp.StatusCode)), "heroes.repository.dotaOpenApi.GetHeroes")
	}

	return body, nil
}
