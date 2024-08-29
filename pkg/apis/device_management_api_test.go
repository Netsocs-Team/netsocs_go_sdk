package apis_test

import (
	"os"
	"testing"

	"github.com/Netsocs-Team/netsocs_go_sdk/pkg/apis"
)

func TestCheckHealth(t *testing.T) {
	os.Setenv("DEVICE_MANAGEMENT_API_HOST", "http://netsocslabs.com:3085")
	os.Setenv("DEVICE_MANAGEMENT_API_USERNAME", "admin")
	os.Setenv("DEVICE_MANAGEMENT_API_PASSWORD", "admin")

	instance, err := apis.NewDeviceManagementApi()
	if err != nil {
		t.Fatal(err)
	}
	if instance == nil {
		t.Fatal("instance is nil")
	}
}
