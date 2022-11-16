package service

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gesquive/fast-cli/fast"
	"golang.org/x/sync/errgroup"
)

type fastComProvider struct {
}

func newFastComProvider() *fastComProvider {
	return &fastComProvider{}
}

// Measure implements SpeedMeasurer interface.
func (p *fastComProvider) Measure() (*MeasureResult, error) {
	urls := fast.GetDlUrls(1)

	if len(urls) == 0 {
		return nil, errors.New("no server urls available")
	}

	downloadSpeed, err := measureNetworkSpeed(download, urls[0])
	if err != nil {
		return nil, fmt.Errorf("failed to perform download test: %w", err)
	}

	uploadSpeed, err := measureNetworkSpeed(upload, urls[0])
	if err != nil {
		return nil, fmt.Errorf("failed to perform upload test: %w", err)
	}

	return &MeasureResult{
		Source:   "fast.com",
		Download: downloadSpeed,
		Upload:   uploadSpeed,
	}, nil
}

var client = http.Client{}

const (
	workload      = 8
	payloadSizeMB = 25.0 // download payload is by default 25MB, make upload 25MB also
)

func measureNetworkSpeed(op func(url string) error, url string) (float64, error) {
	eg := errgroup.Group{}

	startTime := time.Now()
	for i := 0; i < workload; i++ {
		eg.Go(
			func() error {
				return op(url)
			},
		)
	}
	if err := eg.Wait(); err != nil {
		return 0, err
	}
	endTime := time.Now()

	return calculateSpeed(startTime, endTime), nil
}

func calculateSpeed(startTime time.Time, endTime time.Time) float64 {
	return payloadSizeMB * 8 * float64(workload) / endTime.Sub(startTime).Seconds()
}

func download(url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	_, err = io.ReadAll(resp.Body)

	return err
}

func upload(uri string) error {
	v := url.Values{}

	// 10b * x MB / 10 = x MB
	v.Add("content", createUploadPayload())

	resp, err := client.PostForm(uri, v)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	_, err = io.ReadAll(resp.Body)

	return err
}

func createUploadPayload() string {
	return strings.Repeat("0123456789", payloadSizeMB*1024*1024/10)
}
