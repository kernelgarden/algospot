package main

import (
	"bytes"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestGetTotalScore(t *testing.T) {
	testString := makeRandDataSet(10000)
	//testString := makeSameDataSet(10000)
	if len(testString) != 1000 {
		t.Fail()
	}

	if val := GetTotalScore(testString); val != 2500 {
		t.Log(val)
		t.Fail()
	}
}

func makeRandDataSet(n int) string {
	var buffer bytes.Buffer

	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)

	for i := 0; i < n; i++ {
		buffer.WriteString(strconv.Itoa(r.Intn(10)))
	}

	return buffer.String()
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
