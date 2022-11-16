package service

import (
	"github.com/showwin/speedtest-go/speedtest"
)

type speedTestProvider struct {
}

func newSpeedTestProvider() *speedTestProvider {
	return &speedTestProvider{}
}

// Measure implements SpeedMeasurer interface.
func (p *speedTestProvider) Measure() (res *MeasureResult, err error) {
	var (
		user    *speedtest.User
		servers speedtest.Servers
		targets speedtest.Servers
		server  *speedtest.Server
	)

	steps := []func(){
		func() {
			user, err = speedtest.FetchUserInfo()
		},
		func() {
			servers, err = speedtest.FetchServers(user)
		},
		func() {
			targets, err = servers.FindServer(nil)
		},
		func() {
			server = targets[0]
		},
		func() {
			err = server.DownloadTest(false)
		},
		func() {
			err = server.UploadTest(false)
		},
		func() {
			res = &MeasureResult{
				Source:   "speedtest.net",
				Download: server.DLSpeed,
				Upload:   server.ULSpeed,
			}
		},
	}

	next := func(op func()) {
		if err != nil {
			return
		}

		op()
	}

	for _, step := range steps {
		next(step)
	}

	return
}
