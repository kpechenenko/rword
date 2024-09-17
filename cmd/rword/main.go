package main

import (
	"flag"
	"fmt"

	"github.com/kpechenenko/rword/pkg/rword"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 1, "number of words to generate")
	flag.Parse()
	g, err := rword.NewGenerator()
	if err != nil {
		panic(err)
	}
	words := g.GetList(n)
	for _, w := range words {
		fmt.Println(w)
	}
}
