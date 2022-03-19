package util

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func GenQCloudCosPresigned(name string, length int) *url.URL {
	if len(name) == 0 {
		return nil
	}

	type URLHeader struct {
		ContentLength int    `url:"-" header:"content-length"`
		ContentType   string `url:"-" header:"-"`
	}

	ak := os.Getenv("TencentyunAccessKeyID")
	sk := os.Getenv("TencentyunAccessKeySecret")
	bn := os.Getenv("TencentyunBucketName")
	ri := os.Getenv("TencentyunRegionID")
	opt := &URLHeader{
		ContentLength: length,
		// ContentType:   ctype,
	}

	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bn, ri))
	c := cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{})

	ctx := context.Background()
	now := time.Now()
	path := strings.Join([]string{
		"mp-bbhj", "data", now.Format("200601"), now.Format("02"), name}, "/",
	)

	presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodPut, path, ak, sk, time.Hour*24, opt)
	if err != nil {
		return nil
	}

	return presignedURL
}
