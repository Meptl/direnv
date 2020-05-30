package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestEnvDiff(t *testing.T) {
	diff := &EnvDiff{map[string]string{"FOO": "bar"}, map[string]string{"BAR": "baz"}}

	out := diff.Serialize()

	diff2, err := LoadEnvDiff(out)
	if err != nil {
		t.Error("parse error", err)
	}

	if len(diff2.Prev) != 1 {
		t.Error("len(diff2.prev) != 1", len(diff2.Prev))
	}

	if len(diff2.Next) != 1 {
		t.Error("len(diff2.next) != 0", len(diff2.Next))
	}
}

// Issue #114
// Check that empty environment variables correctly appear in the diff
func TestEnvDiffEmptyValue(t *testing.T) {
	before := Env{}
	after := Env{"FOO": ""}

	diff := BuildEnvDiff(before, after)

	if !reflect.DeepEqual(diff.Next, map[string]string(after)) {
		t.Errorf("diff.Next != after (%#+v != %#+v)", diff.Next, after)
	}
}

func TestPreUnloadHooks(t *testing.T) {
	expectPreUnloadHook := func(before, after Env, expectation string) {
		t.Helper()
		diff := BuildEnvDiff(before, after)
		preUnloadHooks := diff.PreUnloadHooks(Zsh)
		if preUnloadHooks != expectation {
			t.Errorf("got \"%s\"; wanted \"%s\"", preUnloadHooks, expectation)
		}
	}

	before := Env{DIRENV_PREUNLOAD: "echo 3;"}
	after := Env{}
	expectPreUnloadHook(before, after, "echo 3;")

	before = Env{}
	after = Env{DIRENV_PREUNLOAD: "echo 3;"}
	expectPreUnloadHook(before, after, "")

	before = Env{DIRENV_PREUNLOAD: loadHooksAsString("echo 3", "echo 4")}
	after = Env{}
	expectPreUnloadHook(before, after, "echo 3;echo 4;")

	before = Env{DIRENV_PREUNLOAD: loadHooksAsString("echo 3", "echo 4")}
	after = Env{DIRENV_PREUNLOAD: loadHooksAsString("echo 3", "echo 4")}
	expectPreUnloadHook(before, after, "")

	before = Env{DIRENV_PREUNLOAD: loadHooksAsString("echo 3", "echo 4")}
	after = Env{DIRENV_PREUNLOAD: loadHooksAsString("echo 3", "echo 5", "echo 4")}
	expectPreUnloadHook(before, after, "")

	before = Env{DIRENV_PREUNLOAD: loadHooksAsString("echo 3", "echo 4", "echo 5")}
	after = Env{DIRENV_PREUNLOAD: loadHooksAsString("echo 3", "echo 5", "echo 4")}
	// Strange, but it is what is it is.
	expectPreUnloadHook(before, after, "echo 5;")

	before = Env{DIRENV_PREUNLOAD: loadHooksAsString("echo 3", "echo 4", "echo 5")}
	after = Env{DIRENV_PREUNLOAD: loadHooksAsString("echo 3", "echo 4")}
	expectPreUnloadHook(before, after, "echo 5;")
}

func TestPostLoadHooks(t *testing.T) {
	expectPostLoadHook := func(before, after Env, expectation string) {
		t.Helper()
		diff := BuildEnvDiff(before, after)
		postLoadHooks := diff.PostLoadHooks(Zsh)
		if postLoadHooks != expectation {
			t.Errorf("got \"%s\"; wanted \"%s\"", postLoadHooks, expectation)
		}
	}

	before := Env{}
	after := Env{DIRENV_POSTLOAD: "echo 3;"}
	expectPostLoadHook(before, after, "echo 3;")

	before = Env{DIRENV_POSTLOAD: "echo 3;"}
	after = Env{}
	expectPostLoadHook(before, after, "")

	before = Env{}
	after = Env{DIRENV_POSTLOAD: loadHooksAsString("echo 3", "echo 4")}
	expectPostLoadHook(before, after, "echo 3;echo 4;")

	before = Env{DIRENV_POSTLOAD: loadHooksAsString("echo 3", "echo 4")}
	after = Env{DIRENV_POSTLOAD: loadHooksAsString("echo 3", "echo 4")}
	expectPostLoadHook(before, after, "")

	before = Env{DIRENV_POSTLOAD: loadHooksAsString("echo 3", "echo 4")}
	after = Env{DIRENV_POSTLOAD: loadHooksAsString("echo 3", "echo 5", "echo 4")}
	expectPostLoadHook(before, after, "echo 5;")

	before = Env{DIRENV_POSTLOAD: loadHooksAsString("echo 3", "echo 4")}
	after = Env{DIRENV_POSTLOAD: "echo 3"}
	expectPostLoadHook(before, after, "")
}

func TestIgnoredEnv(t *testing.T) {
	if !IgnoredEnv(DIRENV_BASH) {
		t.Fail()
	}
	if IgnoredEnv(DIRENV_DIFF) {
		t.Fail()
	}
	if !IgnoredEnv("_") {
		t.Fail()
	}
	if !IgnoredEnv("__fish_foo") {
		t.Fail()
	}
	if !IgnoredEnv("__fishx") {
		t.Fail()
	}
}

func loadHooksAsString(commands ...string) string {
	return strings.Join(commands, LoadHookDelimiter)
}
