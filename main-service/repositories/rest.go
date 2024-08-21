package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type RestRepository struct {
	cl *http.Client
}

func NewRestRepository(cl *http.Client) *RestRepository {
	return &RestRepository{
		cl: cl,
	}
}

func (rc *RestRepository) Send(subj string, data string) error {
	reqUrl, err := url.JoinPath("http://localhost:8081/rest", urlMap[subj])
	if err != nil {
		return fmt.Errorf("failed to join path: %w", err)
	}

	var dummy = struct {
		OrderID string `json:"order_id"`
	}{
		OrderID: data,
	}

	byteData, err := json.Marshal(dummy)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader(byteData))
	if err != nil {
		return fmt.Errorf("failed to created new request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := rc.cl.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	if res.StatusCode < 200 && res.StatusCode > 299 {
		return fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	return nil
}

var urlMap = map[string]string{
	"order.created": "orders",
}
