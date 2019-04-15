package main

import (
	"flag"
	"log"
)

var dir = flag.String("d", ".", "logs dir")
var fileNameReg = flag.String("f", ".*.log", "file name regexp")
var logReg = flag.String("l", ".* /api/playeritems.* ([0-9]+)$", "one line log regexp")
var timePosition = flag.Int("t", 1, "time's position in one line log regexp")
var output = flag.String("o", "./report.txt", "output file path")

func main() {
	flag.Parse()
	analyzer := NewVarianceAnalyzer(*dir, *fileNameReg, *logReg, *timePosition)
	analyzer.Load()
	analyzer.Do()
	analyzer.ReportToFile(*output)
	log.Printf("report done --> %s", *output)
}
