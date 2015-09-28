package gtree

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	var r string
	var w int
	var d int
	var s int

	r = "/tmp/a"
	w = 3
	d = 3
	s = 512

	buildDirectoryTree(r, w, d, s)

}
