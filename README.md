# Golang System Information API

[![Go Report Card](https://goreportcard.com/badge/github.com/benchan777/go-system-info-api)](https://goreportcard.com/report/github.com/benchan777/go-system-info-api)

This program retrieves information about your system such as CPU name, total disk capacity, current disk usage, current free disk apce, total ram capacity, and current available ram. The information is then saved into a file that can be opened for viewing. The information can also be accessed via an API endpoint.

## How to use
- Run go-system-info-api.exe
- Perform GET request on endpoint: `GET localhost:3000/api`
- System information will also be saved in a file named `output.json` in the directory the executable is run in.

## Building
1. Install Go
2. Clone repo: `git clone https://github.com/benchan777/go-system-info-api.git`
3. Enter directory: `cd go-system-info-api`
4. Build the program: `go build`