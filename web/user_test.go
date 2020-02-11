package web

import (
	"net/http"
	"testing"
)

func TestListUser(t *testing.T) {
	uri := "/apis/mina/v1/user/login"
	resp := APIRequest(uri, http.MethodPost, nil)

	if resp.Code != http.StatusOK {
		t.Error(uri, resp.Code)
	}

	t.Log(uri, resp.Code)
	t.Log("resp:", resp.Body)
}
