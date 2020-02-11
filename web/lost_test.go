package web

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/airdb/mina-api/model/vo"
	"github.com/airdb/sailor"
)

func TestStatus(t *testing.T) {
	uri := "/?aaaa=a1"
	resp := APIRequest(uri, "GET", nil)

	if resp.Code != http.StatusOK {
		t.Error(uri, resp.Code)
	}

	t.Log(uri, resp.Code)
	t.Log("output:")
	t.Log(resp.Body)
}

func TestQueryLost(t *testing.T) {
	uri := "/apis/mina/v1/lost/query/1"
	resp := APIRequest(uri, "GET", nil)

	if resp.Code != http.StatusOK {
		t.Error(uri, resp.Code)
	}

	var output vo.Lost

	err := json.Unmarshal(UnmarshalHTTPResponse(resp.Body), &output)
	if err != nil {
		t.Error("errors")
	}

	t.Log(output.Nickname)
}

func TestLostList(t *testing.T) {
	uri := "/apis/mina/v1/lost/list?category=1&page=0&pageSize=10"

	resp := APIRequest(uri, http.MethodGet, nil)

	if resp.Code != http.StatusOK {
		t.Error(uri, resp.Code)
	}

	t.Log(uri, resp.Code)

	var output vo.Lost

	err := json.Unmarshal(UnmarshalHTTPResponse(resp.Body), &output)
	if err != nil {
		t.Error("errors")
	}

	t.Log("output:")
	t.Log(output)
}

func UnmarshalHTTPResponse(body io.Reader) []byte {
	var resp sailor.HTTPAirdbResponse

	err := json.NewDecoder(body).Decode(&resp)
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.Marshal(resp.Data)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}
