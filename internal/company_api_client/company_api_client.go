package companyapiclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/FloRichardAloeCorp/autoprospect/internal/company"
)

const base_url = "https://recherche-entreprises.api.gouv.fr"

type CompanyAPIClient struct {
	resultByPage int
	*http.Client
}

func New() *CompanyAPIClient {
	return &CompanyAPIClient{
		resultByPage: 25,
		Client:       &http.Client{},
	}
}

type Response struct {
	Results      []company.Company `json:"results"`
	TotalResults int               `json:"total_results"`
	Page         int               `json:"page"`
	PerPage      int               `json:"per_page"`
	TotalPages   int               `json:"total_pages"`
}

func (c *CompanyAPIClient) Search(params *SearchRequestQueryParams, page int) (*Response, error) {
	query := params.build()
	query.Add("page", strconv.Itoa(page))
	query.Add("per_page", "25")

	url := base_url + "/search?" + query.Encode()
	res, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read request body: %w", err)
	}

	if res.StatusCode == http.StatusTooManyRequests {
		fmt.Println("Too many requests, retrying in one seconds...")
		time.Sleep(1 * time.Second)
		return c.Search(params, page)
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Invalid status code: " + string(body))
	}

	response := &Response{}

	if err := json.Unmarshal(body, response); err != nil {
		return nil, fmt.Errorf("can't decode response: %w", err)
	}

	return response, nil
}
