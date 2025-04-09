package wal

import (
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/wal"
	"testing"
)

func TestWalWrite(t *testing.T) {
	log, err := wal.Open("mylog", nil)
	assert.Nil(t, err)
	defer log.Close()

	err = log.Write(1, []byte("first entry"))
	assert.Nil(t, err)

	err = log.Write(2, []byte("second entry"))
	assert.Nil(t, err)

	err = log.Write(3, []byte("third entry"))
	assert.Nil(t, err)

	data, err := log.Read(1)
	assert.Nil(t, err)
	t.Log(string(data))
}

func TestWalRead(t *testing.T) {
	log, err := wal.Open("mylog", nil)
	assert.Nil(t, err)
	defer log.Close()

	data, err := log.Read(3)
	assert.Nil(t, err)
	t.Log(string(data))
}

func TestWalTruncate(t *testing.T) {
	log, err := wal.Open("mylog", nil)
	assert.Nil(t, err)
	defer log.Close()

	err = log.TruncateFront(2)
	assert.Nil(t, err)

	err = log.TruncateBack(3)
	assert.Nil(t, err)
}
