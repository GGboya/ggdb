package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dir = "E:\\GoProjects\\github.com\\GGboya\\GG_DB\\tempdata"

func TestOpenDataFile(t *testing.T) {
	dataFile1, err := OpenDataFile(dir, 1)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile1)
	t.Log(dataFile1)
	t.Log(dir)
}
