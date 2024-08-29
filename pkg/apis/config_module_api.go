package apis

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Netsocs-Team/netsocs_go_sdk/pkg/sdk_errors"
)

type ConfigModuleApi struct {
	Host string
}

const CONFIG_MODULE_API_DOMAIN = "Configuration Module API"

func NewConfigModuleApi() (*ConfigModuleApi, error) {
	host := os.Getenv("CONFIGURATION_API_HOST")

	if host == "" {
		return nil, sdk_errors.NewMissingInitialEnvironmentVariablesError(CONFIG_MODULE_API_DOMAIN, []string{"CONFIGURATION_API_HOST"})
	}
	return &ConfigModuleApi{
		Host: os.Getenv("CONFIGURATION_API_HOST"),
	}, nil
}

func (c *ConfigModuleApi) doRequest(method string, path string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, c.Host+path, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}

func (c *ConfigModuleApi) RequestConfigToDevice(configKey string, deviceId int, dataValue []byte) ([]byte, error) {
	requestBody := RequestConfigToDeviceBody{
		TopicKey:  configKey,
		DataValue: string(dataValue),
		IdDevice:  deviceId,
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", "/api/v1/configManager", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
