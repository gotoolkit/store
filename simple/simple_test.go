package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, s)
	key1 := "TestKV1"
	value1 := []byte("TestVal1")
	err = s.Put(key1, value1)
	assert.NoError(t, err)

	key2 := "TestKV2"
	value2 := []byte("TestVal2")
	err = s.Put(key2, value2)
	assert.NoError(t, err)

	av1, err1 := s.Get(key1)
	assert.NoError(t, err1)
	assert.NotNil(t, av1)
	assert.Equal(t, av1, value1)
}
