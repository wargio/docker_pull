package iowrite

import "testing"

func TestFileWrite(t *testing.T) {
	tests := []struct {
		name string
	}{
		{}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Uflie("/tmp/test.txt")
			f.BufWriter.WriteString("1.0")
			f.Close()
		})
	}
}
