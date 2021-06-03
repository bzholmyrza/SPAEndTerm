package main

import (
	"gitlab.com/bzholmyrza/kazdreamtest/testquest/newV"
	"io/ioutil"
	"testing"
)

func BenchmarkOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Old(ioutil.Discard)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newV.New(ioutil.Discard)
	}
}
