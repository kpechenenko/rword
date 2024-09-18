package rword

import (
	"errors"
	"testing"
)

func TestGeneratorGet(t *testing.T) {
	g, err := NewGenerator()
	if err != nil {
		t.Fatal(err)
	}
	w := g.Get()
	if w == "" {
		t.Fatal("empty word")
	}
}

func TestGeneratorGetList(t *testing.T) {
	g, err := NewGenerator()
	if err != nil {
		t.Fatal(err)
	}
	n := 10
	words := g.GetList(n)
	if len(words) != n {
		t.Fatalf("len of generated word list = %d, want %d", len(words), n)
	}
	for _, w := range words {
		if w == "" {
			t.Fatal("empty word in generated word list")
		}
	}
}

func TestGeneratorGetStr(t *testing.T) {
	g, err := NewGenerator()
	if err != nil {
		t.Fatal(err)
	}
	n := 10
	s := g.GetStr(n)
	if len(s) != n {
		t.Fatalf("len of generated str = %d, want %d", len(s), n)
	}
	n = 20
	s = g.GetStr(n)
	if len(s) != n {
		t.Fatalf("len of generated str = %d, want %d", len(s), n)
	}
	n = 1000
	s = g.GetStr(n)
	if len(s) != n {
		t.Fatalf("len of generated str = %d, want %d", len(s), n)
	}
	n = -1
	s = g.GetStr(n)
	if s != "" {
		t.Fatalf("for negavite n generated string should be empty, but was %s", s)
	}
	n = 0
	s = g.GetStr(n)
	if s != "" {
		t.Fatalf("for zero n generated string should be empty, but was %s", s)
	}
}

func TestGeneratorWithDefaultDict(t *testing.T) {
	g, err := NewGenerator()
	if err != nil {
		t.Fatal(err)
	}
	if len(g.dict) == 0 {
		t.Fatal("empty default dict")
	}
	for i, w := range g.dict {
		if w == "" {
			t.Fatalf("empty word in default dict[%d]", i)
		}
	}
}

func TestGeneratorWithDict(t *testing.T) {
	g, err := NewGeneratorWithDict("testdata/valid_dict")
	if err != nil {
		t.Fatal(err)
	}
	word := g.Get()
	found := false
	for _, w := range g.dict {
		if word == w {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("generate word \"%s\" not found in generator dict", word)
	}
}

func TestGeneratorWithEmptyDict(t *testing.T) {
	_, err := NewGeneratorWithDict("testdata/empty_dict")
	if !errors.Is(err, ErrEmptyDict) {
		t.Fatal(err)
	}
}
