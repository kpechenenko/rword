# rword

**rword** is a golang random word generator.

Can generate words using a saved dict containing over 370_000 words.  
Or you can use a custom dictionary.  
For more information, see the section "How to use it" below.

## Installation

You must use golang version 1.21 or higher. 0 dependencies, except for goland stdlib

## How to use it

```go
package main

import (
	"fmt"

	"github.com/kpechenenko/rword/pkg/rword"
)

func main() {
	var g rword.RandomWord
	var err error
	// Create a random word generator using saved dict with 370_000+ words.
	g, err = rword.NewGenerator()
	if err != nil {
		panic(err)
	}
	// Get a random word.
	word := g.Get()
	fmt.Println(word)
	// Get a list of 10 random words.
	words := g.GetList(10)
	for _, w := range words {
		fmt.Println(w)
	}
	// Create a random word generator using your dictionary. Dictionary is a text file, 1 line - 1 word.
	// Returns an error, if dictionary is empty.
	g2, err := rword.NewGeneratorWithDict("path/to/your/dict")
	if err != nil {
		panic(err)
	}
	// Get a random word.
	word = g2.Get()
	fmt.Println(word)
}
```