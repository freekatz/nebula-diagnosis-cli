package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func GetNebulaMetrics(ipAddress string, port int32) ([]string, error) {
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := httpClient.Get(fmt.Sprintf("http://%s:%d/stats", ipAddress, port))
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	metrics := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	return metrics, nil
}

func GetNebulaComponentStatus(ipAddress string, port int32) ([]string, error) {
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := httpClient.Get(fmt.Sprintf("http://%s:%d/status", ipAddress, port))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type nebulaStatus struct {
		GitInfoSha string `json:"git_info_sha"`
		Status     string `json:"status"`
	}

	var status nebulaStatus
	if err := json.Unmarshal(bytes, &status); err != nil {
		return nil, err
	}

	statusMetrics := []string{status.Status}

	return statusMetrics, nil
}

func GetNebulaFlags(ipAddress string, port int32) ([]string, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := httpClient.Get(fmt.Sprintf("http://%s:%d/flags", ipAddress, port))
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	configs := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	return configs, nil
}
