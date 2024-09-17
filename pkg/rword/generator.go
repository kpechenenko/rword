package rword

import (
	"errors"
	"math/rand"
)

// RandomWord a random word generator.
type RandomWord interface {
	// Get random word.
	Get() string
	// GetList of random words.
	GetList(n int) []string
}

// Dictionary _
type Dictionary []string

// Generator generates random words using a saved dictionary.
type Generator struct {
	dict Dictionary // final readonly list of word.
}

// Get a random word.
func (g *Generator) Get() string {
	i := rand.Intn(len(g.dict))
	return g.dict[i]
}

// GetList of random words.
func (g *Generator) GetList(n int) []string {
	var words []string
	if n <= 0 {
		return words
	}
	words = make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = g.Get()
	}
	return words
}

// NewGenerator create a word generator with saved dictionary.
func NewGenerator() (*Generator, error) {
	return NewGeneratorWithDict(GetPathToDefaultDict())
}

var ErrEmptyDict = errors.New("empty dictionary")

// NewGeneratorWithDict create a word generator using custom dictionary.
// dictionary - text file, 1 line = 1 word in dictionary.
func NewGeneratorWithDict(file string) (*Generator, error) {
	dict, err := LoadDictFromFile(file)
	if err != nil {
		return nil, err
	}
	if len(dict) == 0 {
		return nil, ErrEmptyDict
	}
	return &Generator{dict: dict}, nil
}
