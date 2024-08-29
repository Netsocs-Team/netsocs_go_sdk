package apis_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Netsocs-Team/netsocs_go_sdk/pkg/apis"
)

func TestRequestConfigToDevice(t *testing.T) {
	// Create a new ConfigModuleApi
	os.Setenv("CONFIGURATION_API_HOST", "http://netsocslabs.com:3691")
	instance, err := apis.NewConfigModuleApi()
	if err != nil {
		t.Error(err)
	}
	resp, err := instance.RequestConfigToDevice("actionAlarmArmPartition", 465, []byte(`{"partitionId": "1", "value": 1}`))

	if err != nil {
		t.Error(err)
	}

	if resp == nil {
		t.Error("Response is nil")
	}

	fmt.Println(string(resp))
}
