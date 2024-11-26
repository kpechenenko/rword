package rword

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestLoadDictFromFile(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		err      error
		expected Dictionary
	}{
		{
			name:     "empty dictionary",
			file:     "testdata/empty_dict",
			err:      nil,
			expected: Dictionary{},
		},
		{
			name:     "not existing dictionary",
			file:     "testdata/abcde",
			err:      os.ErrNotExist,
			expected: Dictionary{},
		},
		{
			name:     "valid dictionary",
			file:     "testdata/valid_dict",
			err:      nil,
			expected: Dictionary{"a", "b", "c"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d, err := LoadDictFromFile(tc.file)
			if tc.err != nil {
				if !errors.Is(err, tc.err) {
					t.Errorf("expected err: %v, actual: %v", tc.err, err)
				}
				return
			}
			if !reflect.DeepEqual(d, tc.expected) {
				t.Errorf("expected dictionary content: %v, actual: %v", tc.expected, d)
			}
		})
	}
}

func TestGetPathToDefaultDict(t *testing.T) {
	p := GetPathToDefaultDict()
	if !strings.HasSuffix(p, defaultDictFile) {
		t.Errorf("get %v, want somethind ends with %v", p, defaultDictFile)
	}
	if !filepath.IsAbs(p) {
		t.Errorf("get %v, want absolute path", p)
	}
}
