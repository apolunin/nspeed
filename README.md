# Network Speed (nspeed)

This program is intended to test download and upload speed of the current network connection.
For doing network speed measurement one of the two providers can be used:
* Netflix fast.com.
* Ookla speedtest.

Program has a single command line argument which allows to choose a speed test provider:
```shell
Usage of nspeed:
  -provider string
    	specifies a speed measurement provider (speedtest/fast) (default "speedtest")
```

In order to print the aforementioned help the following command can be used:
```shell
nspeed -h
```
### Installation and Usage
In order to run the program from sources do the following:
1. Clone the repository.
2. Run the program via the following command `go run cmd/nspeed.go`.

In order to build a binary do the following:
1. Clone the repository.
2. Build a binary: `go build cmd/nspeed.go`.
3. Run a binary: `./nspeed`.

### Examples
Network speed measurement via `speedtest`:
```shell
$ go run cmd/nspeed.go -provider=speedtest
Starting network speed measurement with "speedtest"...
Source: speedtest.net, Download: 150.053744 Mbps, Upload: 224.004939 Mbps
```
Network speed measurement via `fast.com`:
```shell
$ go run cmd/nspeed.go -provider=fast
Starting network speed measurement with "fast"...
Source: fast.com, Download: 110.968622 Mbps, Upload: 179.001418 Mbps
```