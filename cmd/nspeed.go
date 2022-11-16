package main

import (
	"flag"
	"fmt"

	"github.com/apolunin/nspeed/service"
)

func main() {
	providerFlag := flag.String(
		"provider",
		"speedtest",
		"specifies a speed measurement provider (speedtest/fast)",
	)

	flag.Parse()

	provider, err := service.ParseProvider(*providerFlag)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("Starting network speed measurement with %q...\n", *providerFlag)

	s, err := service.New(provider)
	if err != nil {
		fmt.Printf("failed to create speed test provider: %v\n", err)
		return
	}

	res, err := s.Measure()
	if err != nil {
		fmt.Printf("failed to measure speed: %v\n", err)
		return
	}

	fmt.Println(res)
}
