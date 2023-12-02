package testutil

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func funcName(f any) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	i := strings.LastIndexByte(fullName, '/')

	if i == -1 {
		return fullName
	}

	return fullName[i+1:]
}

func RunTest(t *testing.T, f func(string) string, input, want string) {
	got := f(input)
	if got != want {
		t.Errorf("%s() = %s, want %s", funcName(f), got, want)
	}
}

func RunBench(b *testing.B, f func(string) string, input string) {
	b.Run(funcName(f), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f(input)
		}
	})
}
