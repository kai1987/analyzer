# analyzer
this is a simple log analyzer....

[![GoDoc](https://godoc.org/github.com/kai1987/analyzer?status.svg)](https://godoc.org/github.com/kai1987/analyzer)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE)
[![Build Status](https://api.travis-ci.org/kai1987/analyzer.svg?branch=master)](https://travis-ci.org/kai1987/analyzer)
[![Coverage Status](https://coveralls.io/repos/github/kai1987/analyzer/badge.svg?branch=master)](https://coveralls.io/github/kai1987/analyzer?branch=master)


### Thanks for your time

As this is a tool for other developers, I print all the information about an API rather than just print 90%,95%,99%.

the basic solution about the problem is traverse all the logs, group the used time in a map, and then calculate the percentile.

for example:

logTimes = [1,3,4,4,5,6,11...1000]  
I will store them in a map group by 5 ms. so the map is :  
groupedUsedTime = {0:4,5:2,10:1....}  
then when we generate the report, we need to sort the map by key and calculate the percentile.  

N=len(logTImes)  
M=len(groupedUsedTIme)  
traverse all the logs takes O(N) time complexity. and sort the map and generate the report need O(M+M*log(M) +M) time complexity. and we only need a map to store all the data, so the space complexity is O(M). the worst case is each log time is different and group them by 1 ms, which N==M.


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
