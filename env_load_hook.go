package main

import (
	"strings"
)

// This likely won't work in all shells. It likely won't even work with all
// possible bash commands.
const LoadHookDelimiter = ";"

// ExtractLoadHooks parses a load hook environment string and returns the list of
// commands.
func ExtractLoadHooks(envValue string) []string {
	loadHooks := strings.Split(envValue, LoadHookDelimiter)
	var nonemptyLoadHooks []string
	for _, str := range loadHooks {
		if str != "" {
			nonemptyLoadHooks = append(nonemptyLoadHooks, str)
		}
	}
	return nonemptyLoadHooks
}

// MissingLoadHooksOrdered Finds items in old that are not in new. Respects order
// such that array_removed_ordered([3, 4, 5], [3, 5, 4]) returns [3, 4].
func MissingLoadHooksOrdered(old, new []string) []string {
	var removed []string
	searchIndex := 0
A0_SEARCH:
	for i := 0; i < len(old); i++ {
		for j := searchIndex; j < len(new); j++ {
			if new[j] == old[i] {
				searchIndex = j + 1
				continue A0_SEARCH
			}
		}

		// old value does not exist in new.
		removed = append(removed, old[i])
	}

	return removed
}
