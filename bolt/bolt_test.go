package bolt

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeBoltDBClient(t *testing.T) *Store {
	kv, err := New("/tmp/not_exist_dir/__boltdbtest", WithBucket("bucket"))

	if err != nil {
		t.Fatalf("cannot create store: %v", err)
	}

	return kv
}
func TestMutipleDB(t *testing.T) {
	kv, err := New("/tmp/not_exist_dir/__boltdbtest")
	assert.NoError(t, err)
	assert.NotNil(t, kv)

	kv, err = New("/tmp/not_exist_dir/__boltdbtest")

	assert.Error(t, err)

	_ = os.Remove("/tmp/not_exist_dir/__boltdbtest")

}

func TestGETPUT(t *testing.T) {

	kv, err := New("/tmp/__boltdbtest")
	assert.NoError(t, err)
	assert.NotNil(t, kv)

	key1 := "TestKV1"
	value1 := []byte("TestVal1")
	err = kv.Put(key1, value1)
	assert.NoError(t, err)

	av1, err := kv.Get(key1)
	assert.NoError(t, err)
	assert.NotNil(t, av1)
	assert.Equal(t, av1, value1)
}
