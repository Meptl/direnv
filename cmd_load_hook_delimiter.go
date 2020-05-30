package main

import (
	"fmt"
)

// CmdLoadHookDelim is `direnv load-hook-delimiter`
var CmdLoadHookDelim = &Cmd{
	Name: "load-hook-delimiter",
	Desc: "Prints the delimiter value used internally whilst setting load hook environment variables",
	Action: actionWithConfig(func(currentEnv Env, args []string, config *Config) (err error) {
		fmt.Print(LoadHookDelimiter)
		return nil
	}),
}
