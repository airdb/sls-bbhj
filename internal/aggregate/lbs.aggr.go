package aggregate

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"
	"sort"
	"strings"

	"github.com/airdb/sls-bbhj/pkg/schema/tencent"
	"github.com/go-resty/resty/v2"
)

const (
	TencentMapEntry = "https://apis.map.qq.com"
	GeoCoderURI     = "/ws/geocoder/v1"
)

// LbsAggr defines functions used to handle lbs request.
type LbsAggr interface {
	GeoCoder(ctx context.Context, addr string) (tencent.RespLoc, error)
}

type lbsAggr struct {
	client *resty.Client
	key    string
	sk     string
}

var _ LbsAggr = (*lbsAggr)(nil)

func newLbs() *lbsAggr {
	var (
		tencentMapKey = os.Getenv("TENCENT_MAP_API_KEY")
		tencentMapSk  = os.Getenv("TENCENT_MAP_API_SK")
	)
	if len(tencentMapKey) == 0 || len(tencentMapSk) == 0 {
		panic("can not new lbs without map key config")
	}

	return &lbsAggr{
		resty.New().SetBaseURL(TencentMapEntry),
		tencentMapKey,
		tencentMapSk,
	}
}

// List returns lbs list in the storage. This function has a good performance.
func (aggr lbsAggr) GeoCoder(ctx context.Context, addr string) (tencent.RespLoc, error) {
	params := map[string]string{
		"address": addr,
		// "smart_address": addr,
	}

	params = aggr.sign(GeoCoderURI, params)
	resp, err := aggr.client.R().
		SetQueryParams(params).
		Get(GeoCoderURI)

	msg := tencent.RespLoc{}
	if err := json.Unmarshal(resp.Body(), &msg); err != nil {
		return tencent.RespLoc{}, err
	}

	return msg, err
}

func (aggr lbsAggr) sign(uri string, params map[string]string) map[string]string {
	params["key"] = aggr.key

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var segs []string
	for _, k := range keys {
		segs = append(segs, k+"="+params[k])
	}

	str := uri + "?" + strings.Join(segs, "&") + aggr.sk
	sign := md5.Sum([]byte(str))
	params["sig"] = hex.EncodeToString(sign[:])

	return params
}
