package web

import (
	"net/http"
	"testing"
)

func TestQueryLost(t *testing.T) {
	uri := "/apis/mina/v1/lost/query/1"
	resp := APIRequest(uri, "GET", nil)

	if resp.Code != http.StatusOK {
		t.Error(uri, resp.Code)
	}

	t.Log(uri, resp.Code)
	t.Log("resp:", resp.Body)
}
