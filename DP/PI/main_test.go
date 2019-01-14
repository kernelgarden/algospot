package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestGetTotalScore(t *testing.T) {
	testString := makeDataSet(1000)
	//testString := makeSameDataSet(10000)
	if len(testString) != 1000 {
		t.Fail()
	}

	if val := GetTotalScore(testString); val != 2500 {
		t.Log(val)
		t.Fail()
	}
}

func makeDataSet(n int) string {
	var buffer bytes.Buffer

	for i := 0; i < n; i++ {
		buffer.WriteString(strconv.Itoa((i % 4) + 1))
	}

	return buffer.String()
}

func makeSameDataSet(n int) string {
	var buffer bytes.Buffer

	for i := 0; i < n; i++ {
		buffer.WriteString(strconv.Itoa(1))
	}

	return buffer.String()
}
