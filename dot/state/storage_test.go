package state

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/database"
	"github.com/ChainSafe/gossamer/lib/trie"
)

func newTesStorageState(t *testing.T) *StorageState {
	datadir, err := ioutil.TempDir("", "./test_data")
	if err != nil {
		t.Fatal(err)
	}

	db, err := database.NewBadgerDB(datadir)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewStorageState(db, trie.NewEmptyTrie(nil))
	if err != nil {
		t.Fatal(err)
	}

	return s
}

func TestLoadCodeHash(t *testing.T) {
	storage := newTesStorageState(t)
	testCode := []byte("asdf")

	err := storage.SetStorage(codeKey, testCode)
	if err != nil {
		t.Fatal(err)
	}

	resCode, err := storage.LoadCode()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resCode, testCode) {
		t.Fatalf("Fail: got %s expected %s", resCode, testCode)
	}

	resHash, err := storage.LoadCodeHash()
	if err != nil {
		t.Fatal(err)
	}

	expectedHash, err := common.HexToHash("0xb91349ff7c99c3ae3379dd49c2f3208e202c95c0aac5f97bb24ded899e9a2e83")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resHash[:], expectedHash[:]) {
		t.Fatalf("Fail: got %s expected %s", resHash, expectedHash)
	}
}