package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatSize(t *testing.T) {
	testData := []struct {
		name     string
		input    int64
		expected string
	}{
		{"Zero", 0, "0 Bytes"},
		{"Less than 1 KB", 1023, "1023 Bytes"},
		{"More than 1 KB", 1025, "1.0 KB"},
		{"Less than 1 MB", 1024*1024 - 1, "1024.0 KB"},
		{"More than 1 MB", 1024*1024 + 1, "1.0 MB"},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			result := formatSize(td.input)
			assert.Equal(t, td.expected, result)
		})
	}
}
