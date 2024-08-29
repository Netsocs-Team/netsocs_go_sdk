package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Netsocs-Team/netsocs_go_sdk/pkg/errors"
)

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
		return nil, errors.NewMissingInitialEnvironmentVariablesError("Device Management API", []string{"DEVICE_MANAGEMENT_API_USERNAME", "DEVICE_MANAGEMENT_API_PASSWORD", "DEVICE_MANAGEMENT_API_HOST"})
	}

	instance := &DeviceManagementApi{
		Username: username,
		Password: password,
		Host:     host,
	}
	return instance, instance.CheckHealth()
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
		return errors.NewServiceHealthCheckFailedError("Device Management API", err.Error())
	}
	return nil
}
