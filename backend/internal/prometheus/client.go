package prometheus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const BaseURL = "http://localhost:9090"

type QueryResponse struct {
	Status string `json:"status"`

	Data struct {
		Result []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type NetworkMetric struct {
	Receive  float64 `json:"receive"`
	Transmit float64 `json:"transmit"`
}

type StorageMetric struct {
	AvailableGB float64 `json:"availableGB"`
}

func Query(query string) (*QueryResponse, error) {

	queryURL := fmt.Sprintf(
		"%s/api/v1/query?query=%s",
		BaseURL,
		url.QueryEscape(query),
	)

	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result QueryResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ParseMetricValue extracts the numeric value from the first Prometheus result.
func ParseMetricValue(result *QueryResponse) (float64, error) {

	if len(result.Data.Result) == 0 {
		return 0, fmt.Errorf("no metrics returned from Prometheus")
	}

	if len(result.Data.Result[0].Value) < 2 {
		return 0, fmt.Errorf("invalid Prometheus response")
	}

	valueStr, ok := result.Data.Result[0].Value[1].(string)
	if !ok {
		return 0, fmt.Errorf("metric value is not a string")
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}