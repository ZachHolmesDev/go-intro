package api_test

import (
	"cryptomasters/api"
	"testing"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("ddddd")
	if err != nil {
		t.Errorf("err found: %v", err)
	}
}
