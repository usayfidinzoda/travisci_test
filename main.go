package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	_ "go.uber.org/zap"
	"strings"
)

func main() {
	type T struct {
		I int
	}
	x := []*T{{1}, {2}, {3}}
	y := []*T{{1}, {2}, {4}}
	fmt.Println(cmp.Equal(x, y)) // false
}

func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}
