package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	analyzer := newTestAnalyzer()
	analyzer.Load()
	if len(analyzer.allLogs) == 0 {
		t.Error("TestLoad allLogs is empty")
	}
	analyzer.Do()
	analyzer.ReportToFile("./report_test.txt")
}

func TestDo(t *testing.T) {
	analyzer := newTestAnalyzer()

	testable := []struct {
		in  map[int]int
		out map[float64]int
	}{
		{map[int]int{}, map[float64]int{}},
		{map[int]int{1: 1}, map[float64]int{1.0: 1}},
		{map[int]int{1: 1, 2: 1}, map[float64]int{0.5: 1, 1.0: 2}},
		{map[int]int{1: 1, 2: 1, 10: 2}, map[float64]int{0.25: 1, 0.5: 2, 1: 10}},
	}
	for _, oneCase := range testable {
		analyzer.allLogs = oneCase.in
		analyzer.Do()
		if !reflect.DeepEqual(analyzer.reports, oneCase.out) {
			t.Errorf("input:%v expect %v but got:%v", oneCase.in, oneCase.out, analyzer.reports)
		}
	}
}

func TestReport(t *testing.T) {
	analyzer := newTestAnalyzer()
	testable := []struct {
		in  map[float64]int
		out string
	}{
		{map[float64]int{}, "no data to generate report"},
		{map[float64]int{0.5: 10}, "50.0% of requests return a response within 10 ms\n"},
	}
	for _, oneCase := range testable {
		analyzer.reports = oneCase.in
		report := analyzer.Report()
		if report != oneCase.out {
			fmt.Println(report)
			fmt.Println(oneCase.out)
			t.Errorf("input:%v expect %v but got:%v", oneCase.in, oneCase.out, report)
		}
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	//teardown()
	os.Exit(code)
}

func setup() {
	for i := 1; i < 9; i++ {
		str := ""
		for j := 0; j < 10000; j++ {
			str += fmt.Sprintf("10.2.3.4 [2018/13/10:14:02:39] \"GET /api/playeritems?playerId=3\" 200 %d\n", rand.Intn(300))
		}
		err := ioutil.WriteFile(fmt.Sprintf("2018-04-0%d.log", i), []byte(str), 0644)
		if err != nil {
			log.Printf("variance_test:setup error:%v", err)
		}
	}
}

func teardown() {
	for i := 1; i < 9; i++ {
		err := os.Remove(fmt.Sprintf("2018-04-0%d.log", i))
		if err != nil {
			log.Printf("variance_test:teardown error:%v", err)
		}
	}
}

func newTestAnalyzer() *VarianceAnalyzer {
	return NewVarianceAnalyzer(".", ".*.log", ".* /api/playeritems.* ([0-9]+)$", 1)
}
