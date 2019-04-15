# analyzer
this is a simple log analyzer.

[![GoDoc](https://godoc.org/github.com/kai1987/analyzer?status.svg)](https://godoc.org/github.com/kai1987/analyzer)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE)
[![Build Status](https://api.travis-ci.org/kai1987/analyzer.svg?branch=master)](https://travis-ci.org/kai1987/analyzer)
[![Coverage Status](https://coveralls.io/repos/github/kai1987/analyzer/badge.svg?branch=master)](https://coveralls.io/github/kai1987/analyzer?branch=master)

### Usage

all the executable file are in bin, or you can build it by yourself with go build .

```
linux/
mac/
windows/
```

```
./analyzer --help

Usage of ./analyzer:
  -d string
    	logs dir (default ".")
  -f string
    	file name regexp (default ".*.log")
  -l string
    	one line log regexp (default ".* /api/playeritems.* ([0-9]+)$")
  -o string
    	output file path (default "./report.txt")
  -t int
    	time's position in one line log regexp (default 1)
```

### Example

```
./bin/mac/analyzer -d logs
```


### Running the tests

```
go test -coverprofile analyzer
```

### Thanks for your time

As this is a tool for other developers, I print all the information about an API rather than just print 90%,95%,99%.

the time-complexity is O(N) + O(M * log(M)),the space-complexity is O(M), which N is the logs count, and M is the size of logs grouped by milliseconds;

 
