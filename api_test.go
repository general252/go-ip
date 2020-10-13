package go_ip

import (
	"encoding/json"
	"testing"
)

func TestGetIpAddress(t *testing.T) {
	result, err := GetIpAddress("114.115.116.112")
	if err != nil {
		t.Error(err)
	} else {
		data, _ := json.MarshalIndent(result, "", "  ")
		t.Logf("%v", string(data))
	}
}
