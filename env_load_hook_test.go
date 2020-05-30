package main

import (
	"testing"
)

func TestExtractLoadHooks(t *testing.T) {
	expectLoadHooks := func(env_value string, expectation []string) {
		t.Helper()
		hooks := ExtractLoadHooks(env_value)
		if len(hooks) != len(expectation) {
			t.Errorf("got %v; expected %v", hooks, expectation)
			return
		}
		for i, item := range hooks {
			if item != expectation[i] {
				t.Errorf("got %v; expected %v", hooks, expectation)
				return
			}
		}

	}

	expectLoadHooks("", []string{})
	expectLoadHooks(";;;", []string{})
	expectLoadHooks("echo 3;;;", []string{"echo 3"})
	expectLoadHooks("echo 3;echo 4; echo 5;", []string{"echo 3", "echo 4", " echo 5"})
}
