package search

import (
	"regexp"
)

var separator = regexp.MustCompile(`[^\w]+`)
