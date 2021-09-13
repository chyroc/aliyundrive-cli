package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_formatSize(t *testing.T) {
	as := assert.New(t)
	as.Equal("2.16M", formatFileSize(2266742))
}
