package service

import (
	"fmt"
	"strings"
)

type (
	// Provider is used to enumerate all possible speed test providers.
	Provider int

	// MeasureResult provides a result for network speed measurement.
	MeasureResult struct {
		Source   string
		Download float64
		Upload   float64
	}

	// SpeedMeasurer is used to measure speed via concrete speed test provider.
	SpeedMeasurer interface {
		Measure() (*MeasureResult, error)
	}
)

// Available speed test providers.
const (
	SpeedTest Provider = 1
	FastCom   Provider = 2
)

// String implements Stringer interface.
func (m *MeasureResult) String() string {
	return fmt.Sprintf(
		"Source: %s, Download: %f Mbps, Upload: %f Mbps",
		m.Source, m.Download, m.Upload,
	)
}

// New is a constructor for a speed test provider.
func New(provider Provider) (SpeedMeasurer, error) {
	switch provider {
	case SpeedTest:
		return newSpeedTestProvider(), nil
	case FastCom:
		return newFastComProvider(), nil
	}

	return nil, fmt.Errorf("unknown speed measurement provider: %d", provider)
}

// ParseProvider parses a provider string and returns a corresponding Provider value.
func ParseProvider(provider string) (Provider, error) {
	switch strings.ToLower(provider) {
	case "fast":
		return FastCom, nil
	case "speedtest":
		return SpeedTest, nil
	}

	return 0, fmt.Errorf("unknown speed measurement provider: %v", provider)
}
