package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Netsocs-Team/netsocs_go_sdk/pkg/sdk_errors"
)

const DEVICE_MANAGEMENT_API_DOMAIN = "Device Management API"

type DeviceManagementApi struct {
	Username string
	Password string
	Host     string
	token    string
}

func NewDeviceManagementApi() (*DeviceManagementApi, error) {
	username := os.Getenv("DEVICE_MANAGEMENT_API_USERNAME")
	password := os.Getenv("DEVICE_MANAGEMENT_API_PASSWORD")
	host := os.Getenv("DEVICE_MANAGEMENT_API_HOST")
	if username == "" || password == "" || host == "" {
		return nil, sdk_errors.NewMissingInitialEnvironmentVariablesError(DEVICE_MANAGEMENT_API_DOMAIN, []string{"DEVICE_MANAGEMENT_API_USERNAME", "DEVICE_MANAGEMENT_API_PASSWORD", "DEVICE_MANAGEMENT_API_HOST"})
	}

	instance := &DeviceManagementApi{
		Username: username,
		Password: password,
		Host:     host,
	}
	return instance, instance.CheckHealth()
}

func (d *DeviceManagementApi) GetToken() string {
	return d.token
}

func (d *DeviceManagementApi) SetToken(token string) {
	d.token = token
}

func (d *DeviceManagementApi) doRequest(method string, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, d.Host+path, body)
	if err != nil {
		return nil, err
	}
	// TODO: Improve login handling
	err = d.Login()

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+d.token)

	client := &http.Client{}
	return client.Do(req)

}

func (d *DeviceManagementApi) Login() error {
	client := &http.Client{}
	response := &DeviceManagementApiLoginResponse{}
	loginRequestBody := &DeviceManagementApiLoginRequest{
		Username: d.Username,
		Password: d.Password,
	}
	loginRequestBodyJson, err := json.Marshal(loginRequestBody)
	if err != nil {
		return err
	}
	loginRequest, err := http.NewRequest("POST", d.Host+"/api/v1/login", bytes.NewBuffer(loginRequestBodyJson))

	if err != nil {
		return err
	}

	loginRequest.Header.Set("Content-Type", "application/json")
	loginResponse, err := client.Do(loginRequest)

	if err != nil {
		return err
	}

	defer loginResponse.Body.Close()

	err = json.NewDecoder(loginResponse.Body).Decode(response)
	if err != nil {
		return err
	}

	if response.Token == "" {
		return fmt.Errorf("no token returned")
	}

	d.token = response.Token
	return nil
}

func (d *DeviceManagementApi) CheckHealth() error {
	err := d.Login()
	if err != nil {
		return sdk_errors.NewServiceHealthCheckFailedError(DEVICE_MANAGEMENT_API_DOMAIN, err.Error())
	}
	return nil
}

func (d *DeviceManagementApi) One(id int) (*DeviceManagementApiDeviceSchema, error) {
	response := struct {
		Status string                          `json:"status"`
		Data   DeviceManagementApiDeviceSchema `json:"data"`
		Error  string                          `json:"error"`
	}{}
	resp, err := d.doRequest("GET", fmt.Sprintf("/api/v1/devices/%d", id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		switch response.Error {
		case "Device not found":
			return nil, sdk_errors.NewNotFoundItemsError(DEVICE_MANAGEMENT_API_DOMAIN, []string{fmt.Sprintf("id: %d", id)})
		default:
			return nil, fmt.Errorf(response.Error)
		}
	}

	return &response.Data, nil
}
