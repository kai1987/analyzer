package main

import (
	"io/ioutil"
	"strings"
)

func readLines(filename string) ([]string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	str := string(b)
	arr := strings.Split(str, "\n")
	return arr, err
}
