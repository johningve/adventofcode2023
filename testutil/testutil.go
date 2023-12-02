package testutil

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func RunTest(t *testing.T, f func(string) string, input, want string) {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	i := strings.LastIndexByte(fullName, '/')

	var name string
	if i == -1 {
		name = fullName
	} else {
		name = fullName[i+1:]
	}

	got := f(input)
	if got != want {
		t.Errorf("%s() = %s, want %s", name, got, want)
	}
}

func RunBench(b *testing.B, f func(string) string, input string) {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	i := strings.LastIndexByte(fullName, '/')

	var name string
	if i == -1 {
		name = fullName
	} else {
		name = fullName[i+1:]
	}

	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f(input)
		}
	})
}
