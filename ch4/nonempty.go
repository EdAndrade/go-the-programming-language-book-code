package main

import (
	"fmt"
)

func nonempty(strings []string) {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
}