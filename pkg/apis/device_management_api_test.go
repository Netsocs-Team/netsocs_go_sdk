package apis_test

import (
	"errors"
	"os"
	"testing"

	"github.com/Netsocs-Team/netsocs_go_sdk/pkg/apis"
	"github.com/Netsocs-Team/netsocs_go_sdk/pkg/sdk_errors"
)

func startEnvs() {
	os.Setenv("DEVICE_MANAGEMENT_API_HOST", "http://netsocslabs.com:3085")
	os.Setenv("DEVICE_MANAGEMENT_API_USERNAME", "admin")
	os.Setenv("DEVICE_MANAGEMENT_API_PASSWORD", "admin")
}

func TestCheckHealth(t *testing.T) {
	startEnvs()
	instance, err := apis.NewDeviceManagementApi()
	if err != nil {
		t.Fatal(err)
	}
	if instance == nil {
		t.Fatal("instance is nil")
	}
}

func TestOne(t *testing.T) {
	startEnvs()
	instance, err := apis.NewDeviceManagementApi()
	if err != nil {
		t.Fatal(err)
	}
	device, err := instance.One(1)
	if err != nil {
		t.Fatal(err)
	}
	if device == nil {
		t.Fatal("device is nil")
	}
	if device.ID != 1 {
		t.Fatal("device.ID != 1")
	}
}

func TestOneWhenNotExists(t *testing.T) {
	startEnvs()
	instance, err := apis.NewDeviceManagementApi()
	if err != nil {
		t.Fatal(err)
	}
	device, err := instance.One(999999)
	notFoundErr := &sdk_errors.NOT_FOUND_ITEMS{}
	if !errors.As(err, &notFoundErr) {
		t.Fatal(err)
	}
	if device != nil {
		t.Fatal("device is not nil")
	}
}
