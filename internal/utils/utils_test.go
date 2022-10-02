package utils_test

import (
	"os"
	"testing"

	"github.com/deestarks/routnd/internal/utils"
)

func TestToInt(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1", 1},
		{"2232", 2232},
		{"323", 323},
	}

	for _, c := range cases {
		got := utils.ToInt(c.in)
		if got != c.want {
			t.Errorf("ToInt(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestSplitCommand(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"echo test", []string{"echo", "test"}},
		{"ls -l ./test/dir", []string{"ls", "-l", "./test/dir"}},
	}

	for _, c := range cases {
		program, args := utils.SplitCommand(c.in)
		if program != c.want[0] {
			t.Errorf("SplitCommand(%q) == %q, want %q", c.in, program, c.want[0])
		}

		if len(args) != len(c.want)-1 {
			t.Errorf("SplitCommand(%q) == %q, want %q", c.in, args, c.want[1:])
		}

		for i, arg := range args {
			if arg != c.want[i+1] {
				t.Errorf("SplitCommand(%q) == %q, want %q", c.in, arg, c.want[i+1])
			}
		}
	}
}

func TestViewFile(t *testing.T) {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		t.Errorf("Error creating test file: %s", err)
	}

	utils.ViewFile(file.Name(), 10, false)
	utils.ViewFile(file.Name(), 2, false)
	utils.ViewFile(file.Name(), 0, false)

	os.Remove(file.Name())
}