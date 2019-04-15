package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"sync"
)

//VarianceAnalyzer
//this is a analyzer that analyzer all the logs and get the variance about an API.
type VarianceAnalyzer struct {
	dir             string         //logs file's dir
	fileNameReg     *regexp.Regexp //log file name regexp
	logReg          *regexp.Regexp //one line log's regexp
	logTimePosition int            //time position in the log's regexp
	allLogs         IntMap         //{ms period : logsCount}
	allCount        int
	reports         map[float64]int //{percent : within ms}
	mutex           sync.RWMutex
	msGap           int //the ms gap, default 5ms
}

func NewVarianceAnalyzer(dir, fileNameRegStr, logRegStr string, logTimePosition int) *VarianceAnalyzer {
	analyzer := &VarianceAnalyzer{
		dir:             dir,
		allLogs:         make(map[int]int),
		reports:         make(map[float64]int),
		logTimePosition: logTimePosition,
		fileNameReg:     regexp.MustCompile(fileNameRegStr),
		logReg:          regexp.MustCompile(logRegStr),
		msGap:           5,
	}
	return analyzer
}

//Load
//load all the logs in allLogs, use multi go routine.
func (analyzer *VarianceAnalyzer) Load() {
	var wg sync.WaitGroup
	files, err := ioutil.ReadDir(analyzer.dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !analyzer.fileNameReg.MatchString(f.Name()) {
			continue
		}
		wg.Add(1)
		go func(fileName string) {
			analyzer.loadOneFile(filepath.Join(analyzer.dir, fileName))
			wg.Done()
		}(f.Name())
	}
	wg.Wait()
}

//Do do the analyze
func (analyzer *VarianceAnalyzer) Do() {
	analyzer.allCount = analyzer.allLogs.ValueSum()
	keys := analyzer.allLogs.SortedKeys()
	lastPercent := 0.0
	lastCount := 0
	for _, v := range keys {
		lastCount += analyzer.allLogs[v]
		newPercent := float64(lastCount) / float64(analyzer.allCount)
		if newPercent-lastPercent > 0.01 {
			analyzer.reports[newPercent] = v
			lastPercent = newPercent
		}
	}
}

//Report make a report
func (analyzer *VarianceAnalyzer) Report() string {
	percents := make([]float64, 0, len(analyzer.reports))
	for k, _ := range analyzer.reports {
		percents = append(percents, k)
	}
	sort.Float64s(percents)
	reportStr := ""
	for _, p := range percents {
		reportStr += fmt.Sprintf("%.1f%% of requests return a response within %d ms\n", p*100, analyzer.reports[p])
	}
	if reportStr == "" {
		return "no data to generate report"
	}
	return reportStr
}

//ReportToFile report to a file
func (analyzer *VarianceAnalyzer) ReportToFile(filePath string) {
	err := ioutil.WriteFile(filePath, []byte(analyzer.Report()), 0644)
	if err != nil {
		log.Printf("variance_test:setup error:%v", err)
	}
}

func (analyzer *VarianceAnalyzer) pickMs(line string) (ms int, found bool) {
	result := analyzer.logReg.FindStringSubmatch(line)
	if len(result) > analyzer.logTimePosition {
		ms, err := strconv.Atoi(result[analyzer.logTimePosition])
		if err != nil {
			log.Printf("VarianceAnalyzer:pickMs Atoi err:%v", err)
			return 0, false
		}
		return ms, true
	}
	return 0, false
}

func (analyzer *VarianceAnalyzer) loadOneFile(fileName string) {
	lines, err := readLines(fileName)
	if err != nil {
		log.Fatalf("VarianceAnalyzer:Load:%s:err:%v", fileName, err)
	}
	tempMap := make(map[int]int)
	for _, line := range lines {
		ms, found := analyzer.pickMs(line)
		if !found {
			continue
		}
		msTruncated := (ms - 1) / analyzer.msGap * analyzer.msGap
		if msTruncated < 0 {
			msTruncated = 0
		}
		tempMap[msTruncated]++
	}
	analyzer.mutex.Lock()
	defer analyzer.mutex.Unlock()
	for k, v := range tempMap {
		analyzer.allLogs[k] += v
	}
}
