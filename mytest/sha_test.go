package mytest

import (
	"crypto/sha1"
	"encoding/hex"
	"testing"
)

func TestSignature(t *testing.T) {
	data := "hello"
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	res := hex.EncodeToString(sha1.Sum([]byte(nil)))
	t.Log(res)
}
