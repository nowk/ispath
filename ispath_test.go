package ispath

import (
	"testing"

	"gopkg.in/nowk/assert.v2"
)

func TestIsPath(t *testing.T) {
	for _, v := range []struct {
		path string
		is   bool
	}{
		{"path", false},
		{"/path", true},
		{"path/", true},
		{"path/more", true},
		{"/path/more", true},
		{"/path/", true},
	} {
		assert.Equal(t, v.is, IsPath(v.path))
	}
}
