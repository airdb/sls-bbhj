package util

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

func TestGenQCloudCosPut(t *testing.T) {
	log.Println(GenQCloudCosPresigned("demo.png", 5))

	ak := os.Getenv("TencentyunAccessKeyID")
	sk := os.Getenv("TencentyunAccessKeySecret")
	bn := os.Getenv("TencentyunBucketName")
	ri := os.Getenv("TencentyunRegionID")

	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bn, ri))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 COS_SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: ak,
			// 环境变量 COS_SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: sk,
			// Debug 模式，把对应 请求头部、请求内容、响应头部、响应内容 输出到标准输出
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	name := "test/example"
	f := strings.NewReader("test")

	_, err := c.Object.Put(context.Background(), name, f, nil)

	log.Println(err)
}
